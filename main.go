package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	Name         string
	Path         string
	ModifiedTime string
	Hash         string
	Size         int64
}

func save(filepath string, files []File) error {
	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(files)
}

func load(filepath string) []File {
	file, err := os.Open(filepath)
	if err != nil {
		return nil
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

func sha256sum(path string) (string, error) {
	file, err := os.Open(path)
	if os.IsNotExist(err) || os.IsPermission(err) {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		panic(err)
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func getFiles(directory string) ([]File, error) {
	var files []File

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		hash, _ := sha256sum(path)
		files = append(files, File{
			Name:         info.Name(),
			Path:         path,
			ModifiedTime: info.ModTime().String(),
			Size:         info.Size(),
			Hash:         hash,
		})

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return files, nil
}

func constaint(files []File, file File) bool {
	for _, item := range files {
		if item.Name == file.Name && item.Path == file.Path && item.ModifiedTime == file.ModifiedTime && item.Size == file.Size && item.Hash == file.Hash {
			return true
		}
	}

	return false
}

func verify(oldFiles []File, newFiles []File) (bool, []File) {
	// not efficient, but works

	var newVerifiedData []File

	for _, file := range newFiles {
		if constaint(oldFiles, file) {
			continue
		}

		newVerifiedData = append(newVerifiedData, file)
		log.Println("New file found:", file.Path)
	}

	if len(newVerifiedData) != 0 {
		return true, newVerifiedData
	}

	return false, nil
}

func main() {
	const directory string = "/home/ao/Downloads"
	const savePath string = "files.json"
	logFile, _ := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // err later
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiWriter)

	verifiedData := load(savePath)
	if verifiedData == nil {
		log.Println("No verified data found, creating new file.")
		verifiedData, _ = getFiles(directory)
		save(savePath, verifiedData)
	} else {
		log.Println("Verified data found, loading existing file.")
	}

	// fmt.Println(verifiedData)

	// for _, file := range verifiedData {
	// 	log.Println(file.Name)
	// 	log.Println(file.Path)
	// 	log.Println(file.ModifiedTime)
	// 	log.Println(file.Size)
	// 	log.Println(file.Hash)

	// 	fmt.Println("-----------------------------------")
	// 	// break
	// }

	for {
		time.Sleep(10 * time.Second)
		log.Println("Checking for changes...")
		files, err := getFiles(directory)
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		// Check files.json exists
		if _, err := os.Stat(savePath); os.IsNotExist(err) {
			log.Println("files.json not found, creating new file.")
			save(savePath, files)
			continue
		}

		status, newFile := verify(verifiedData, files)
		if status {
			verifiedData = append(verifiedData, newFile...)
			save(savePath, verifiedData)
		} else {
			log.Println("No changes found.")
		}
	}
}
