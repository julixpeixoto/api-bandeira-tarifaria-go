package schedule

import (
	"api-bandeira-tarifaria-go/client"
	"api-bandeira-tarifaria-go/data"
	"api-bandeira-tarifaria-go/database"
	"fmt"
	"time"
)

func Run() {
	go scheduler(1440, insert)
}

func scheduler(intervalSeconds int, taskFunc func()) {
	for {
		taskFunc()
		time.Sleep(time.Duration(intervalSeconds) * time.Second)
	}
}

func insert() {
	databaseData := database.GetAll()
	crawlerData := client.GetData()
	var dataToInsert [][]string

	for i, itemCrawler := range crawlerData {
		if i == 0 {
			continue
		}

		exists := false

		for _, itemDatabase := range databaseData {
			if itemCrawler[2] == itemDatabase.Mounth && itemCrawler[1] == fmt.Sprintf("%d", itemDatabase.Year) {
				exists = true
			}
		}

		if exists == false {
			dataToInsert = append(dataToInsert, itemCrawler)
		}
	}

	flagsToInsert := data.FormatData(dataToInsert)

	if len(flagsToInsert) > 0 {
		database.Insert(flagsToInsert)
		fmt.Println("Inserted", len(flagsToInsert), "new flag(s)")
	} else {
		fmt.Println("No new flags to insert")
	}
}
