package stats

import (
	"fmt"
	"main/utils"
	"slices"
	"sort"
)

func Display_stats(folderData []utils.File, config utils.Config) {
	statArr, keys, sum := sort_map(folderData)
	fmt.Printf("project lenght: %vk characters\n", sum/1000)
	fmt.Println("language breakdown:")
	for _, key := range keys {
		data := statArr[key]
		fmt.Printf(".%v%v: %v%% (%v files)\n", key, known_lang(config.Known_extensions, key), (100*data.Lenght)/sum, data.FileCount)
	}
}

func sort_map(folderData []utils.File) (map[string]utils.Language, []string, int64) {
	var lang_size_stat = make(map[string]utils.Language)
	var sum int64 = 0
	var keys []string
	for _, file := range folderData {
		sum += file.File_lenght
		data, exist := lang_size_stat[file.Extension]
		if exist {
			lang_size_stat[file.Extension] = utils.Language{Lenght: data.Lenght + file.File_lenght, FileCount: data.FileCount + 1}
			continue
		}
		keys = append(keys, file.Extension)
		lang_size_stat[file.Extension] = utils.Language{Lenght: file.File_lenght, FileCount: 1}
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return lang_size_stat[keys[i]].Lenght > lang_size_stat[keys[j]].Lenght
	})
	return lang_size_stat, keys, sum
}

func known_lang(known_langs []string, lang string) string {
	if slices.Contains(known_langs, lang) {
		return ""
	} else {
		return " (unknown)"
	}
}
