package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ログイン画面表示
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Page"))
}

// ログイン処理
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login User"))
}

// 新規登録画面表示
func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register Page"))
}

// 新規登録処理
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register User"))
}

// ユーザー画面表示
func UserPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	w.Write([]byte("User ID: " + userID))
}

// 作業情報取得
func WorkInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	w.Write([]byte("Work info for user ID: " + userID))
}

// 一緒にいた時間の取得
func TogetherTimeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	w.Write([]byte("Together time for user ID: " + userID))
}

// 集中ランキング取得
func RankingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ranking list"))
}
