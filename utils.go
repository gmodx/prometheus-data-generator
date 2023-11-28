package main

import (
	"encoding/json"
	"io"
	"os"
)

func GetTFromFile[T any](path string) (T, error) {
	var item T

	file, err := os.Open(path)
	if err != nil {
		return item, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return item, err
	}

	err = json.Unmarshal(data, &item)
	return item, err
}
