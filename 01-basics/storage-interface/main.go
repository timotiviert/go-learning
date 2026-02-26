package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Storage interface {
	Save(key, value string) error
	Load(key string) (string, error)
}

type FileStorage struct {
	Dir string
}

func NewFileStorage() (*FileStorage, error) {
	// Return absolute path to executed Go file
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return nil, errors.New("could not get filename")
	}

	dir := filepath.Join(filepath.Dir(filename), "files")

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err

	}

	return &FileStorage{
		Dir: dir,
	}, nil
}

func (fs *FileStorage) Save(key, value string) error {
	path := filepath.Join(fs.Dir, key)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintln(f, value)
	return nil
}

func (fs *FileStorage) Load(key string) (string, error) {
	path := filepath.Join(fs.Dir, key)
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	var fs Storage

	fs, err := NewFileStorage()
	if err != nil {
		log.Fatal()
	}

	if err := fs.Save("key", "hi"); err != nil {
		log.Fatal(err)
	}

	value, err := fs.Load("key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
