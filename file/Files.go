package file

import (
	"bufio"
	utils "main/utils"
	"os"

	"gopkg.in/yaml.v3"
)

func Read_directory(path string) {
	scan_folder(path)
}

func scan_folder(workingDirectory string) {
	// TODO
}

func scan_file(folderName string, fileName string) {

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
