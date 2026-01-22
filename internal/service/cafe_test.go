package service

import (
	"testing"

	"github.com/MaximK0valev/cafe-api-go/internal/model"
	"github.com/stretchr/testify/assert"
)

// TestGetCafes_UnknownCity проверяет,
// что сервис возвращает ошибку, если запрошен город,
// которого нет в источнике данных.
func TestGetCafes_UnknownCity(t *testing.T) {
	_, err := GetCafes(model.CafeQuery{
		City: "omsk",
	})

	assert.ErrorIs(t, err, ErrUnknownCity)
}

// TestGetCafes_Limit проверяет,
// что параметр Limit ограничивает количество возвращаемых кафе.
func TestGetCafes_Limit(t *testing.T) {
	result, err := GetCafes(model.CafeQuery{
		City:  "moscow",
		Limit: 2,
	})

	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

// TestGetCafes_LimitOverflow проверяет,
// что если Limit больше количества доступных кафе,
// сервис возвращает все доступные кафе и не падает.
func TestGetCafes_LimitOverflow(t *testing.T) {
	result, err := GetCafes(model.CafeQuery{
		City:  "tula",
		Limit: 100,
	})

	assert.NoError(t, err)
	assert.Len(t, result, 3)
}

// TestGetCafes_Search проверяет,
// что поиск по названию кафе работает
// и не зависит от регистра символов.
func TestGetCafes_Search(t *testing.T) {
	result, err := GetCafes(model.CafeQuery{
		City:   "moscow",
		Search: "кофе",
	})

	assert.NoError(t, err)
	assert.Len(t, result, 2)
}
