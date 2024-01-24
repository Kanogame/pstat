package main

import (
	file "main/file"
	stats "main/stats"
)

func main() {
	config := file.Parse_config("./config.yaml")
	folderData := file.Read_directory("./", config)
	stats.Get_stats(folderData)
}
