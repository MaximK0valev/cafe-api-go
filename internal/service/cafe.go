package service

import (
	"errors"
	"strings"

	"github.com/MaximK0valev/cafe-api-go/data"

	"github.com/MaximK0valev/cafe-api-go/internal/model"
)

var (
	ErrUnknownCity = errors.New("unknown city")
)

func GetCafes(q model.CafeQuery) ([]string, error) {
	cafes, ok := data.Cafes[q.City]
	if !ok {
		return nil, ErrUnknownCity
	}

	if q.Search != "" {
		var filtered []string
		search := strings.ToLower(q.Search)

		for _, c := range cafes {
			if strings.Contains(strings.ToLower(c), search) {
				filtered = append(filtered, c)
			}
		}
		cafes = filtered
	}

	if q.Limit <= 0 || q.Limit > len(cafes) {
		q.Limit = len(cafes)
	}

	return cafes[:q.Limit], nil
}
