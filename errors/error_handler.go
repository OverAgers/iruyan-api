package errors

import (
	"iruyan-api/utils" // 汎用のJSONユーティリティをインポート
	"net/http"
)

// ErrorHandler構造体
type ErrorHandler struct{}

// BadRequest: 400エラーを返す
func (e *ErrorHandler) BadRequest(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Bad Request"
	}
	data := map[string]interface{}{
		"message": message,
		"status":  400,
	}
	utils.SendJSON(w, data, http.StatusBadRequest)
}

// Unauthorized: 401エラーを返す
func (e *ErrorHandler) Unauthorized(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Unauthorized"
	}
	data := map[string]interface{}{
		"message": message,
		"status":  401,
	}
	utils.SendJSON(w, data, http.StatusUnauthorized)
}

// Forbidden: 403エラーを返す
func (e *ErrorHandler) Forbidden(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Forbidden"
	}
	data := map[string]interface{}{
		"message": message,
		"status":  403,
	}
	utils.SendJSON(w, data, http.StatusForbidden)
}

// NotFoundError: 404エラーを返す
func (e *ErrorHandler) NotFoundError(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Not Found"
	}
	data := map[string]interface{}{
		"message": message,
		"status":  404,
	}
	utils.SendJSON(w, data, http.StatusNotFound)
}

// InternalServerError: 500エラーを返す
func (e *ErrorHandler) InternalServerError(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Internal Server Error"
	}
	data := map[string]interface{}{
		"message": message,
		"status":  500,
	}
	utils.SendJSON(w, data, http.StatusInternalServerError)
}

// MessageError: 汎用的なエラーメッセージを返す
func (e *ErrorHandler) MessageError() map[string]interface{} {
	return map[string]interface{}{
		"code":    "---",
		"message": "Error",
	}
}
