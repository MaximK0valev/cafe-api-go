package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MaximK0valev/cafe-api-go/internal/model"
	"github.com/MaximK0valev/cafe-api-go/internal/service"
)

func CafeHandler(w http.ResponseWriter, r *http.Request) {
	limit := 25
	if v := r.URL.Query().Get("count"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil {
			http.Error(w, "incorrect count", http.StatusBadRequest)
			return
		}
		limit = n
	}

	query := model.CafeQuery{
		City:   r.URL.Query().Get("city"),
		Search: r.URL.Query().Get("search"),
		Limit:  limit,
	}

	result, err := service.GetCafes(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}
