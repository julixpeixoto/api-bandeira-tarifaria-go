package client

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetData() [][]string {
	res, err := http.Get("https://sic.cercos.com.br/sic/bandeiras_tarifarias/")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result [][]string

	doc.Find("table.scGridTabela tbody tr").Each(func(_ int, rowSelection *goquery.Selection) {
		var row []string

		rowSelection.Find("td").Each(func(_ int, cellSelection *goquery.Selection) {
			cellText := cellSelection.Text()
			row = append(row, strings.TrimSpace(cellText))
		})

		if len(row) > 0 {
			result = append(result, row)
		}
	})

	return result
}
