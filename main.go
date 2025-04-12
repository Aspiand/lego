package main

import (
	"log"
	"os"
)

type File struct {
	Name         string
	Path         string
	ModifiedTime string
	Size         int64
}

func get_files(path string) []File {
	entries, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	var files []File

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		stat, err := os.Stat(path + "/" + entry.Name())

		if err != nil {
			log.Fatal(err)
			return nil
		}

		files = append(files, File{
			Name:         stat.Name(),
			Path:         path + "/" + stat.Name(),
			ModifiedTime: stat.ModTime().String(),
			Size:         stat.Size(),
		})
	}

	return files
}

func main() {
	var files []File = get_files("/home/ao/Downloads")

	for _, file := range files {
		log.Println(file.Name)
		log.Println(file.Path)
		log.Println(file.ModifiedTime)
		log.Println(file.Size)
	}
}
