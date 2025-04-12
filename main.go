package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"os"
)

type File struct {
	Name         string
	Path         string
	ModifiedTime string
	Hash         string
	Size         int64
}

func save(files []File) error {
	file, err := os.Create("files.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(files)
}

func load() []File {
	file, err := os.Open("files.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var files []File
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&files)
	if err != nil {
		panic(err)
	}

	return files
}

func sha256sum(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		panic(err)
	}

	return hex.EncodeToString(hash.Sum(nil))
}

func getFiles(directory string) []File {
	entries, err := os.ReadDir(directory)

	if err != nil {
		panic(err)
	}

	var files []File

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		var filePath string = directory + "/" + entry.Name()
		stat, err := os.Stat(filePath)

		if err != nil {
			panic(err)
		}

		files = append(files, File{
			Name:         stat.Name(),
			Path:         filePath,
			ModifiedTime: stat.ModTime().String(),
			Size:         stat.Size(),
			Hash:         sha256sum(filePath),
		})
	}

	return files
}

func main() {
	var files []File = getFiles("/home/ao/Downloads")
	save(files)

	// var files []File = load()

	for _, file := range files {
		log.Println(file.Name)
		log.Println(file.Path)
		log.Println(file.ModifiedTime)
		log.Println(file.Size)
		log.Println(file.Hash)
		break
	}
}
