package util

import (
	"bufio"
	"os"
)

func ReadFromFile(filename string) ([]string, error) {
	read := []string{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		read = append(read, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return read, nil
}
