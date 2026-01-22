package main

import (
	"log"
	"net/http"

	"github.com/MaximK0valev/cafe-api-go/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cafes", handler.CafeHandler)

	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
