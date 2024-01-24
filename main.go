package main

import (
	file "main/file"
)

func main() {
	config := file.Parse_config("./config.yaml")
	folderData := file.Read_directory("./", config)
}
