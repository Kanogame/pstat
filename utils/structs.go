package utils

type Config struct {
	Excluded_folders []string
	Show_full_stat   bool
}

type File struct {
	File_lenght int64
	Extension   string
}

type Language struct {
	Lenght    int64
	FileCount int32
}
