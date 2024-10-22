package routes

import (
	"iruyan-api/handlers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {
	// User関連のルーティング
	r.HandleFunc("/u/ranking", handlers.RankingHandler).Methods("GET") // クエリパラメータは任意で処理
	r.HandleFunc("/login", handlers.LoginPageHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/register", handlers.RegisterPageHandler).Methods("GET")
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/u/{user_id}", handlers.UserPageHandler).Methods("GET") // 動的パスは固定パスの後に定義
	r.HandleFunc("/u/{user_id}/work_info", handlers.WorkInfoHandler).Methods("GET")
	r.HandleFunc("/u/{user_id}/together", handlers.TogetherTimeHandler).Methods("GET") // クエリパラメータは任意
}
