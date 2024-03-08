package main

import (
	"fmt"
	"log"
	"net/http"
	"pbp/UTS/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/rooms/{id_game}", controllers.GetAllRoomsByIdGame).Methods("GET")
	router.HandleFunc("/rooms2", controllers.GetDetailRooms).Methods("GET")
	router.HandleFunc("/rooms/enter", controllers.InsertRoom).Methods("POST")
	// router.HandleFunc("/rooms", controllers.LeaveRoom).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
