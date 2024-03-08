package models

type RoomById struct {
	Id        int    `json:"id"`
	Room_name string `json:"room_name"`
}

type RoomByIdResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    RoomById `json:"data"`
}

type RoomByIdResponses struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []RoomById `json:"data"`
}

type Game struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Max_player int    `json:"max_player"`
}

type Participant struct {
	ID        int    `json:"id"`
	IDAccount int    `json:"id_account"`
	Username  string `json:"username"`
}

type RoomDetail struct {
	ID           int         `json:"id"`
	RoomName     string      `json:"room_name"`
	Participants Participant `json:"participants"`
}

type RoomDetailResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    RoomDetail `json:"data"`
}

type RoomsDetailResponse struct {
	RoomDetail []RoomDetail `json:"room"`
}

type RoomDetailResponses struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    RoomsDetailResponse `json:"data"`
}

type Room struct {
	ID        int    `json:"id"`
	Room_name string `json:"room_name"`
	ID_game   Game   `json:"id_game"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data"`
}

type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
