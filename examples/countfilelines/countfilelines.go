package main

import (
	"bufio"
	"os"
	"path/filepath"
)

func main() {
	files, _ := filepath.Glob(os.Getenv("HOME") + "/crystal/src/*.cr")

	channel := make(chan int, 1)
	for _, f := range files {
		go func(f string) {
			file, _ := os.Open(f)
			fileScanner := bufio.NewScanner(file)
			lineCount := 0
			for fileScanner.Scan() {
				lineCount++
			}
			channel <- lineCount
		}(f)
	}

	totalLines := 0
	for i := 0; i < len(files); i++ {
		totalLines += <-channel
	}

	close(channel)
}
