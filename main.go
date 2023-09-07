package main

import (
	"api-bandeira-tarifaria-go/client"
	"api-bandeira-tarifaria-go/data"
	"api-bandeira-tarifaria-go/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	crawler := client.GetData()
	formattedData := data.FormatData(crawler)

	fmt.Println(formattedData)

	fmt.Println("Listening on port 8080")

	flags := []models.Flag{
		{
			ID:        1,
			Flag:      "Verde",
			Value:     0,
			Mounth:    "Janeiro",
			MountNum:  1,
			Year:      2023,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(flags)
	})

	http.ListenAndServe(":8080", nil)
}
