package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// ルーターの作成
	r := mux.NewRouter()

	// エンドポイントを定義
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/users", UsersHandler).Methods("GET")

	// サーバーを起動
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
