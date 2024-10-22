package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	// テスト用のリクエストとレスポンスレコーダーを作成
	request := httptest.NewRequest("POST", "/login", strings.NewReader("username=john&password=1234&icon=smile"))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	// ハンドラを実行
	LoginHandler(recorder, request)

	// レスポンスをチェック
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Login User: john\nPassword: 1234\nIcon: smile"
	if recorder.Body.String() != expected {
		t.Errorf("unexpected body: got %v want %v", recorder.Body.String(), expected)
	}
}

func TestLoginHandlerBadRequest(t *testing.T) {
	// テスト用のリクエストとレスポンスレコーダーを作成
	request := httptest.NewRequest("POST", "/login", strings.NewReader("username=john&password=123&icon=smile"))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	// ハンドラを実行
	LoginHandler(recorder, request)

	// レスポンスをチェック
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expectedMessage := "Password must be exactly 4 digits"
	var responseBody map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("could not parse response body: %v", err)
	}
	if responseBody["message"] != expectedMessage {
		t.Errorf("unexpected message: got %v want %v", responseBody["message"], expectedMessage)
	}
}
