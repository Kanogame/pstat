package utils

type Config struct {
	Excluded_folders []string
	Show_full_stat   bool
}

type File struct {
	File_lenght int64
	extension   string
}
