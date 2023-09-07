package data

func FormatData(data [][]string) []Flag {
	var flags []Flag

	for _, item := range data {
		flag := Flag{
			Flag:   item[3],
			Value:  item[4],
			Mounth: item[2],
			Year:   item[1],
		}

		flags = append(flags, flag)
	}

	return flags
}
