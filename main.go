package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: 1, Name: "Sample Item 1"},
	{ID: 2, Name: "Sample Item 2"},
}

func main() {
	http.HandleFunc("/items", getItems)
	http.HandleFunc("/health", healthCheck)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getItems(w http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query().Get("limit")

	if limitParam != "" {
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			http.Error(w, "invalid limit value", http.StatusBadRequest)
			return
		}

		if limit > len(items) {
			// This could be improved → mock issue opportunity
			http.Error(w, "limit exceeds available items", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(items[:limit])
		return
	}

	json.NewEncoder(w).Encode(items)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
