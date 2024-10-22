package utils

import (
	"encoding/json"
	"net/http"
)

// sendJSON: JSONレスポンスを返す汎用的な関数
func SendJSON(w http.ResponseWriter, data map[string]interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
