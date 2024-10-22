package handlers

import (
	"fmt"
	"net/http"

	"iruyan-api/errors"
	"iruyan-api/models"

	"github.com/gorilla/mux"
)

// ログイン画面表示
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Page"))
}

// ログイン処理
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// フォームデータからusernameとpasswordを取得
	username := r.FormValue("username")
	password := r.FormValue("password")

	// エラーハンドラのインスタンスを作成
	errorHandler := errors.ErrorHandler{}

	// ここでは仮にデータベースからユーザーを取得する関数 GetUserByUsername があるとする
	// 本来はデータベースからusernameをもとにユーザーを検索する処理が必要です。
	user, err := models.GetUserByUsername(username)
	if err != nil {
		// ユーザーが見つからなかった場合、エラーレスポンスを返す
		errorHandler.BadRequest(w, "User not found")
		return
	}

	// パスワードが一致するかチェック
	if !user.CheckPasswordHash(password) {
		// パスワードが一致しない場合、エラーレスポンスを返す
		errorHandler.Unauthorized(w, "Invalid password")
		return
	}

	// ログイン成功時のレスポンス
	response := fmt.Sprintf("Login successful! Welcome %s", user.Username)
	w.Write([]byte(response))
}

// 新規登録画面表示
func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register Page"))
}

// 新規登録処理
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// フォームデータから値を取得
	username := r.FormValue("username")
	password := r.FormValue("password")
	name := r.FormValue("name")
	icon := r.FormValue("icon")
	task := r.FormValue("task")
	status := r.FormValue("status")
	email := r.FormValue("email")

	// エラーハンドラのインスタンスを作成
	errorHandler := errors.ErrorHandler{}

	// User構造体のインスタンスを作成（この時点でパスワードハッシュ化とバリデーションも行われる）
	user, err := models.NewUser(name, username, password, icon, task, status, email)
	if err != nil {
		// パスワードが4桁でない場合や他のバリデーションに失敗した場合、エラーレスポンスを返す
		errorHandler.BadRequest(w, err.Error())
		return
	}

	// すべてのデータをレスポンスとして表示（パスワードはハッシュ化されているので表示しない）
	response := fmt.Sprintf(
		"Register User: \nUsername: %s\nName: %s\nIcon: %s\nTask: %s\nStatus: %s\nEmail: %s\n",
		user.Username,
		user.Name,
		user.Icon,
		user.Task,
		user.Status,
		user.Email,
	)

	// レスポンスを送信
	w.Write([]byte(response))
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
