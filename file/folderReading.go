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
	scan_folder(path, 0, config.Excluded_folders, &folderData)
	return folderData, folderSize
}

func scan_folder(workingDirectory string, depth int, except []string, folderData *[]utils.File) {
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
			scan_folder(workingDirectory+"/"+file.Name(), depth+1, except, folderData)
		} else {
			if strings.Contains(file.Name(), ".") {
				*folderData = append(*folderData, scan_file(workingDirectory, file.Name()))
			}
		}
	}
}

func scan_file(folder_name string, file_name string) utils.File {
	Split := strings.Split(file_name, ".")
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
