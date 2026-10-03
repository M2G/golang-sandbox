package main

import (
	"fmt"
	"io"
	"os"
)

// file readFile
func readFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close() // clean ressource

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}

func main() {
	err := readFile("output.txt")
	if err != nil {
		fmt.Println(err)
	}
}
