package main

import (
	"api-bandeira-tarifaria-go/database"
	"api-bandeira-tarifaria-go/schedule"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening on port", 8080)
	database.Migrate()
	schedule.Insert()

	flags := database.GetAll()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(flags)
	})

	http.ListenAndServe(":8080", nil)
}
