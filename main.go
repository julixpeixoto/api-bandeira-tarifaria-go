package main

import (
	"api-bandeira-tarifaria-go/database"
	"api-bandeira-tarifaria-go/schedule"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	database.Migrate()
	schedule.Run()

	flags := database.GetAll()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(flags)
	})

	fmt.Println("Listening on port", 8080)
	http.ListenAndServe(":8080", nil)
}
