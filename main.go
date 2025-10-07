package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	go receiveLine(ch, f)
	return ch
}

func receiveLine(ch chan string, f io.ReadCloser) {
	b := make([]byte, 8)
	currentLine := ""
	for {
		n, err := f.Read(b)
		parts := strings.Split(string(b[:n]), "\n")
		for i := 0; i < len(parts)-1; i++ {
			currentLine += parts[i]
			ch <- currentLine
			currentLine = ""
		}
		currentLine += parts[len(parts)-1]
		if err == io.EOF {
			if currentLine != "" {
				ch <- currentLine
			}
			close(ch)
			break
		}
	}
}

func main() {
	fileName := "messages.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file, %s", fileName)
	}
	ch := getLinesChannel(file)
	for line := range ch {
		fmt.Printf("read: %s\n", line)
	}
}
