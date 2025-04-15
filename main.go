package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
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

func load(filepath string) (files []File) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&files)
	if err != nil {
		panic(err)
	}

	return files
}

func sha256sum(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Printf("Error computing file hash for %s: %v", path, err)
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func getFiles(directory string) (files []File) {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		} else if fi, _ := os.Lstat(path); fi.Mode()&os.ModeSymlink != 0 { // skip symbolic links
			return nil
		}

		// log.Printf("Checking: %s", path)

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
		return nil
	}

	return files
}

func verify(oldFiles []File, newFiles []File) (added []File, modified []File, removed []File) {
	old, new := make(map[string]File), make(map[string]File)

	for _, f := range oldFiles {
		old[f.Path] = f
	}

	for _, f := range newFiles {
		new[f.Path] = f
	}

	for path, oldFile := range old {
		if _, exists := new[path]; !exists {
			removed = append(removed, oldFile)
		}
	}

	for path, newFile := range new {
		oldFile, exists := old[path]

		if !exists {
			added = append(added, newFile)
		} else if oldFile.Hash != newFile.Hash {
			modified = append(modified, oldFile)
		}
	}

	return added, modified, removed
}

func dumpFiles(files []File) {
	for _, file := range files {
		log.Println(file.Name)
		log.Println(file.Path)
		log.Println(file.ModifiedTime)
		log.Println(file.Size)
		log.Println(file.Hash)

		fmt.Println("-----------------------------------")
	}
}

func main() {
	const directory string = "/home/ao/.local/share/lutris/games/zzz"
	const savePath string = "files.json"
	logFile, _ := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // err later
	defer logFile.Close()

	log.SetFlags(log.Ltime | log.Lshortfile)
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	verifiedData := load(savePath)
	if verifiedData == nil {
		log.Println("No verified data found, creating new file.")
		verifiedData = getFiles(directory)
		save(savePath, verifiedData)
	} else {
		log.Println("Verified data found, loading existing file.")
	}

	for {
		log.Println("Checking for changes...")
		files := getFiles(directory)

		// Check files.json exists
		if _, err := os.Stat(savePath); os.IsNotExist(err) {
			log.Println("files.json not found, creating new file.")
			save(savePath, files)
			continue
		}

		added, modified, removed := verify(verifiedData, files)
		if changed := len(added) + len(modified) + len(removed); changed > 0 {
			log.Printf("%d changes detected, saving to %s...", changed, savePath)
			verifiedData = files
			save(savePath, files)
		}

		for _, file := range added {
			log.Printf("New file found at %s\n", file.Path)
		}

		for _, file := range modified {
			log.Printf("File changed at %s\n", file.Path)
		}

		for _, file := range removed {
			log.Printf("File removed at %s\n", file.Path)
		}

		time.Sleep(5 * time.Minute)
	}
}
