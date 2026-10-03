package main

import (
	"fmt"
	"io"
	"os"
	"time"
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

func processData(data []int) {
	start := time.Now()
	defer func() {
		fmt.Println("process time:", time.Since(start))
	}()

	for _, value := range data {
		fmt.Println(value)
		time.Sleep(time.Millisecond * 100)
	}
}

func safeOperation() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic", err)
		}
	}()
	panic("Something went wrong")
	fmt.Println("Cannot reach here")
}

func main() {
	err := readFile("output.txt")
	if err != nil {
		fmt.Println(err)
	}

	data := []int{1, 2, 3, 4}
	processData(data)

	safeOperation()

}
