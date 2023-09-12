package main

import (
	"api-bandeira-tarifaria-go/database"
	"api-bandeira-tarifaria-go/schedule"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	port := viper.GetString("PORT")
	database.Migrate()
	schedule.Run()

	flags := database.GetAll()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(flags)
	})

	fmt.Println("Listening on port", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
