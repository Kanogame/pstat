package main

import (
	"fmt"
	file "main/file"
	stats "main/stats"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You should provide path")
		os.Exit(0)
	}
	path := os.Args[1]
	config := file.Parse_config("./config.yaml")
	folderData := file.Read_directory(path, config)
	stats.Display_stats(folderData, config)
}
