package main

import (
	"bufio"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Excluded_folders []string
	Show_full_stat   bool
}

func main() {
	config := parse_config("./config.yaml")
	fmt.Println(config)
}

func parse_config(path string) Config {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data := Config{}
	yaml.Unmarshal(file, &data)
	return data
}

func read_file(path string) []string {
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
