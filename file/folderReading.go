package file

import (
	utils "main/utils"
	"os"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"
)

const maxDepth = 10

func Read_directory(path string, config utils.Config) ([]utils.File, int64) {
	var folderSize = get_file_lenght(path)
	var folderData = make([]utils.File, 0)
	scan_folder(path, 0, config.Excluded_folders, &folderData, config.Known_extensions, config.Known_only)
	return folderData, folderSize
}

func scan_folder(workingDirectory string, depth int, except []string, folderData *[]utils.File, knownArr []string, isKnown bool) {
	if depth >= maxDepth {
		return
	}
	dir, err := os.ReadDir(workingDirectory)
	utils.HandleError(err)
	for _, file := range dir {
		if slices.Contains(except, file.Name()) {
			continue
		}
		if file.IsDir() {
			scan_folder(workingDirectory+"/"+file.Name(), depth+1, except, folderData, knownArr, isKnown)
		} else {
			if strings.Contains(file.Name(), ".") {
				scanned := scan_file(workingDirectory, file.Name(), knownArr, isKnown)
				if (scanned != utils.File{}) {
					*folderData = append(*folderData, scanned)
				}
			}
		}
	}
}

func scan_file(folder_name string, file_name string, knownArr []string, isKnown bool) utils.File {
	Split := strings.Split(file_name, ".")
	if isKnown && !slices.Contains(knownArr, Split[len(Split)-1]) {
		return utils.File{}
	}
	return utils.File{File_lenght: get_file_lenght(folder_name + "/" + file_name), Extension: Split[len(Split)-1]}
}

func get_file_lenght(path string) int64 {
	file, err := os.Open(path)
	utils.HandleError(err)
	defer file.Close()
	stat, err := file.Stat()
	utils.HandleError(err)
	return stat.Size()
}

func Parse_config(path string) utils.Config {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data := utils.Config{}
	yaml.Unmarshal(file, &data)
	return data
}
