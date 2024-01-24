package utils

type Config struct {
	Excluded_folders []string
	Show_full_stat   bool
	Known_only       bool
	Known_extensions []string
}

type File struct {
	File_lenght int64
	Extension   string
}

type Language struct {
	Lenght    int64
	FileCount int32
}
