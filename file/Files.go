package file

import (
	"bufio"
	"fmt"
	utils "main/utils"
	"os"

	"gopkg.in/yaml.v3"
)

func Read_directory(path string) {
	dirs, err := os.ReadDir(path)
	utils.HandleError(err)
	for _, da := range dirs {
		fmt.Println(da.IsDir())
	}
}

func File_lenght(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(stat.Size())
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
