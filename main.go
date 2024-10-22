package main

import (
	"fmt"
	"log"
	"net/http"

	"iruyan-api/routes" // routesパッケージをインポート

	"github.com/gorilla/mux"
)

func main() {
	// 新しいルーターを作成
	r := mux.NewRouter()

	// ルート定義（"/"へのリクエストを処理）
	r.HandleFunc("/", HomeHandler).Methods("GET")

	// Room関連のルーティングを登録
	routes.RegisterRoomRoutes(r)

	// User関連のルーティングを登録
	routes.RegisterUserRoutes(r)

	// サーバーをポート8080で開始
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// HomeHandler は "/" へのGETリクエストを処理
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Home Page!"))
}
