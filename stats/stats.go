package stats

import (
	"fmt"
	"main/utils"
)

func Get_stats(folderData []utils.File) {
	var lang_size_stat = make(map[string]utils.Language)

	for _, file := range folderData {
		data, exist := lang_size_stat[file.Extension]
		if exist {
			lang_size_stat[file.Extension] = utils.Language{Lenght: data.Lenght + file.File_lenght, FileCount: data.FileCount + 1}
			continue
		}
		lang_size_stat[file.Extension] = utils.Language{Lenght: file.File_lenght, FileCount: 1}
	}

	fmt.Println(lang_size_stat)
}
