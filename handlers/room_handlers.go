package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Roomの詳細表示ハンドラ
func RoomHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID := vars["room_id"]
	w.Write([]byte("Room ID: " + roomID))
}

// 着席ハンドラ
func SeatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seatID := vars["seat_id"]
	w.Write([]byte("Seat ID: " + seatID + " - Seated"))
}

// 離席ハンドラ
func LeaveSeatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seatID := vars["seat_id"]
	w.Write([]byte("Seat ID: " + seatID + " - Left"))
}
