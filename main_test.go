package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerSuccess(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=4", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидался статус 200 OK")
	assert.NotEmpty(t, responseRecorder.Body.String(), "Тело ответа не должно быть пустым")
}

func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=St.Petersburg&count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Ожидался статус 400 Bad Request")
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Ожидался ответ 'wrong city value'")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=8&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	responseBody := responseRecorder.Body.String()
	cafes := strings.Split(responseBody, ",")
	assert.Equal(t, totalCount, len(cafes), "expected cafe count: %d, got %d", totalCount, len(cafes))
	// здесь нужно добавить необходимые проверки
}
