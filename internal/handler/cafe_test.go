package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCafeHandler_OK проверяет,
// что handler корректно обрабатывает валидный запрос
// и возвращает HTTP 200 и JSON-массив кафе.
func TestCafeHandler_OK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafes?city=moscow&count=2", nil)
	rec := httptest.NewRecorder()

	CafeHandler(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var result []string
	err := json.Unmarshal(rec.Body.Bytes(), &result)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

// TestCafeHandler_UnknownCity проверяет,
// что handler возвращает 400,
// если запрошен неизвестный город.
func TestCafeHandler_UnknownCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafes?city=omsk", nil)
	rec := httptest.NewRecorder()

	CafeHandler(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// TestCafeHandler_BadCount проверяет,
// что handler возвращает 400,
// если параметр count нельзя преобразовать в число.
func TestCafeHandler_BadCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafes?city=moscow&count=abc", nil)
	rec := httptest.NewRecorder()

	CafeHandler(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
