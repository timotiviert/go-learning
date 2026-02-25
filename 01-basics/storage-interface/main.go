package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Storage interface {
	Save(key, value string) error
	Load(key string) (string, error)
}

type FileStorage struct {
	Path string
}

func NewFileStorage(path string) (*FileStorage, error) {
	header := []string{"key", "value"}

	// Get file stats.
	info, err := os.Stat(path)
	var isNewFile bool
	if os.IsNotExist(err) {
		isNewFile = true
	} else if err != nil {
		return nil, err
	} else {
		isNewFile = info.Size() == 0
	}

	if isNewFile {
		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		writer := csv.NewWriter(f)

		if err := writer.Write(header); err != nil {
			return nil, err
		}
		writer.Flush()
		if err := writer.Error(); err != nil {
			return nil, err
		}
	}

	return &FileStorage{
		Path: path,
	}, nil
}

func (fs *FileStorage) Save(key, value string) error {
	f, err := os.OpenFile(fs.Path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)

	if err := writer.Write([]string{key, value}); err != nil {
		return err
	}
	writer.Flush()
	return writer.Error()
}

func (fs *FileStorage) Load(key string) (string, error) {

}
