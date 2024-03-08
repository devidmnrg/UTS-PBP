package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	m "pbp/UTS/models"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllRoomsByIdGame(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	idGame := vars["id_game"]

	query := "SELECT id, room_name FROM rooms WHERE id_game = ?"

	rows, err := db.Query(query, idGame)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 400, "Can't connect to Database")
		return
	}
	defer rows.Close()

	var room m.RoomById
	var rooms []m.RoomById

	for rows.Next() {

		if err := rows.Scan(&room.Id, &room.Room_name); err != nil {
			log.Println(err)
			SendErrorResponse(w, 500, "Internal Server Error")
			return
		} else {
			rooms = append(rooms, room)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	var response m.RoomByIdResponses

	response.Status = 200
	response.Message = "Success"
	response.Data = rooms
	json.NewEncoder(w).Encode(response)
}

func GetDetailRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := `SELECT r.ID, r.room_name, p.ID, p.id_account, a.username FROM rooms r LEFT JOIN participants p ON r.ID = p.id_room LEFT JOIN accounts a ON p.id_account = a.ID`

	rows, err := db.Query(query)
	if err != nil {
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	var roomDetail m.RoomDetail
	var roomDetails []m.RoomDetail

	for rows.Next() {
		if err := rows.Scan(
			&roomDetail.ID, &roomDetail.RoomName, &roomDetail.Participants.ID, &roomDetail.Participants.IDAccount, &roomDetail.Participants.Username); err != nil {
			print(err.Error())
			return
		} else {
			roomDetails = append(roomDetails, roomDetail)
		}
	}
	var response m.RoomDetailResponses
	w.Header().Set("Content-Type", "application/json")
	response.Status = 200
	response.Message = "Success"
	response.Data.RoomDetail = roomDetails
	json.NewEncoder(w).Encode(response)
}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	if err := r.ParseForm(); err != nil {
		SendErrorResponse(w, 500, "error parsing form data")
		return
	}

	roomIDStr := r.Form.Get("id_room")
	accountIDStr := r.Form.Get("id_account")

	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 400, "Invalid room ID")
		return
	}

	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 400, "Invalid account ID")
		return
	}

	var room m.Room
	err = db.QueryRow("SELECT r.id, r.room_name, g.id, g.name, g.max_player FROM rooms r JOIN games g ON r.id_game = g.id WHERE r.id = ?", roomID).
		Scan(&room.ID, &room.Room_name, &room.ID_game.Id, &room.ID_game.Name, &room.ID_game.Max_player)
	if err == sql.ErrNoRows {
		SendErrorResponse(w, 404, "Room not found")
		return
	} else if err != nil {
		log.Println(err)
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	var participantCount int
	err = db.QueryRow("SELECT COUNT(id) FROM participants WHERE id_room = ?", roomID).Scan(&participantCount)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	if participantCount >= room.ID_game.Max_player {
		SendErrorResponse(w, 400, "Room is at its maximum limit")
		return
	}

	_, err = db.Exec("INSERT INTO participants (id_room, id_account) VALUES (?, ?)", roomID, accountID)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	SendSuccessResponse(w, 200, "Entered the room successfully")
}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	if err := r.ParseForm(); err != nil {
		SendErrorResponse(w, 500, "error parsing form data")
		return
	}

	roomIDStr := r.Form.Get("id_room")
	accountIDStr := r.Form.Get("id_account")

	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 400, "Invalid room ID")
		return
	}

	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 400, "Invalid account ID")
		return
	}

	var participantID int
	err = db.QueryRow("SELECT id FROM participants WHERE id_room = ? AND id_account = ?", roomID, accountID).
		Scan(&participantID)
	if err == sql.ErrNoRows {
		SendErrorResponse(w, 404, "Participant not found in the room")
		return
	} else if err != nil {
		log.Println(err)
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	_, err = db.Exec("DELETE FROM participants WHERE id = ?", participantID)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	SendSuccessResponse(w, 200, "Left the room successfully")
}
