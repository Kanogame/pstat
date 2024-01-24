package main

import (
	"fmt"
	file "main/file"
)

func main() {
	config := file.Parse_config("./config.yaml")
	file.File_lenght("./config.yaml")
	file.Read_directory("./")
	fmt.Println(config)
}
