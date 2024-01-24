package stats

import (
	"fmt"
	"main/utils"
)

func Get_stats(folderData []utils.File, folderSize int64) {
	var lang_size_stat = make(map[string]utils.Language)
	var sum int64 = 0
	for _, file := range folderData {
		sum += file.File_lenght
		data, exist := lang_size_stat[file.Extension]
		if exist {
			lang_size_stat[file.Extension] = utils.Language{Lenght: data.Lenght + file.File_lenght, FileCount: data.FileCount + 1}
			continue
		}
		lang_size_stat[file.Extension] = utils.Language{Lenght: file.File_lenght, FileCount: 1}
	}

	fmt.Printf("%vk characters\n", sum/1000)
	fmt.Println(folderSize)
	for lang, data := range lang_size_stat {
		fmt.Printf(".%v: %v%%\n", lang, (100*data.Lenght)/sum)
	}
}
