package main

import (
	"fmt"
	file "main/file"
)

func main() {
	config := file.Parse_config("./config.yaml")
	file.Read_directory("./", config)
	fmt.Println(config)
}
