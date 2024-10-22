package routes

import (
	"iruyan-api/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoomRoutes(r *mux.Router) {
	// Room関連のルーティング
	r.HandleFunc("/r/{room_id}", handlers.RoomHandler).Methods("GET")
	r.HandleFunc("/r/{room_id}/{seat_id}", handlers.SeatHandler).Methods("PUT")      // 着席
	r.HandleFunc("/r/{room_id}/{seat_id}", handlers.LeaveSeatHandler).Methods("PUT") // 離席
}
