package file

import (
	"bufio"
	"fmt"
	utils "main/utils"
	"os"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"
)

const maxDepth = 10

func Read_directory(path string, config utils.Config) {
	scan_folder("./", 0, config.Excluded_folders)
}

func scan_folder(workingDirectory string, depth int, except []string) {
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
			scan_folder(workingDirectory+"/"+file.Name(), depth+1, except)
		}
		fmt.Println(scan_file(workingDirectory, file.Name()))
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

func Read_file(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
