package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
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
	if err != nil && os.IsNotExist(err) {
		log.Fatal("File not found:", path)
		return ""
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		panic(err)
	}

	return hex.EncodeToString(hash.Sum(nil))
}

func getFiles(directory string) ([]File, error) {
	var files []File

	e := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		files = append(files, File{
			Name:         info.Name(),
			Path:         path,
			ModifiedTime: info.ModTime().String(),
			Size:         info.Size(),
			Hash:         sha256sum(path),
		})

		return nil
	})

	if e != nil {
		panic(e)
	}

	return files, nil
}

func main() {
	path := "/home/ao/Downloads"
	// path = "/nix/store"
	files, err := getFiles(path)

	if err != nil {
		log.Fatal(err)
	}

	// return

	// save(files)
	// // var files []File = load()

	for _, file := range files {
		log.Println(file.Name)
		log.Println(file.Path)
		log.Println(file.ModifiedTime)
		log.Println(file.Size)
		log.Println(file.Hash)

		// break
	}

	// for {
	// 	time.Sleep(1 * time.Second)
	// 	log.Println("Checking for changes...")
	// }
}
