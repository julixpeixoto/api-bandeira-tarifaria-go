package data

import (
	"api-bandeira-tarifaria-go/models"
	"strconv"
	"strings"
)

func FormatData(data [][]string) []models.Flag {
	var flags []models.Flag

	for _, item := range data {
		value := strings.ReplaceAll(item[4], ",", ".")
		valueFloat, _ := strconv.ParseFloat(value, 64)
		yearInt, _ := strconv.Atoi(item[1])

		flag := models.Flag{
			Flag:     item[3],
			Value:    valueFloat,
			Mounth:   item[2],
			MountNum: getMounthNum(item[2]),
			Year:     int64(yearInt),
		}

		flags = append(flags, flag)
	}

	return flags
}

func getMounthNum(mounth string) int64 {
	switch {
	case mounth == "Janeiro":
		return 1
	case mounth == "Fevereiro":
		return 2
	case mounth == "Março":
		return 3
	case mounth == "Abril":
		return 4
	case mounth == "Maio":
		return 5
	case mounth == "Junho":
		return 6
	case mounth == "Julho":
		return 7
	case mounth == "Agosto":
		return 8
	case mounth == "Setembro":
		return 9
	case mounth == "Outubro":
		return 10
	case mounth == "Novembro":
		return 11
	case mounth == "Dezembro":
		return 12
	}

	return 0
}
