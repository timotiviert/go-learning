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
	err := os.WriteFile(path, []byte(value), 0644)
	if err != nil {
		return err
	}
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

type MemoryStorage struct {
	Map map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Map: make(map[string]string),
	}
}

func (ms *MemoryStorage) Save(key, value string) error {
	ms.Map[key] = value
	return nil
}

func (ms *MemoryStorage) Load(key string) (string, error) {
	value, ok := ms.Map[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return value, nil
}

func main() {
	var fs, ms Storage

	fs, err := NewFileStorage()
	if err != nil {
		log.Fatal(err)
	}

	ms = NewMemoryStorage()

	if err := fs.Save("type", "FileStorage"); err != nil {
		log.Fatal(err)
	}

	if err := ms.Save("type", "MemoryStorage"); err != nil {
		log.Fatal(err)
	}

	value, err := fs.Load("type")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

	value, err = ms.Load("type")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
