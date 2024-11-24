package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil) // Создаю запрос с указанными параметрами

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)   // Проверяю код состояния
	require.NotEmpty(t, responseRecorder.Body.String(), "") // Проверяю не пустое ли тело ответа
}

func TestMainHandlerWhenMissingCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=tula", nil) // Создаю запрос с указанными параметрами

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest) // Проверяю код состояния
	assert.Equal(t, responseRecorder.Body.String(), `wrong city value`)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	count := strconv.Itoa(totalCount + 2)                                       // Запрашиваю больше кафе, чем существует
	req := httptest.NewRequest("GET", "/cafe?count="+count+"&city=moscow", nil) // Создаю запрос с указанными параметрами

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusOK) // Проверяю код состояния
	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)
}
