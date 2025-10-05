package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := "messages.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file, %s", fileName)
	}

	b := make([]byte, 8)
	for {
		n, err := file.Read(b)
		if n > 0 {
			fmt.Printf("read: %s\n", b[:n])
		}
		if err == io.EOF {
			os.Exit(0)
		}
		if err != io.EOF && err != nil {
			log.Fatalf("Failed to read file, %s", fileName)
		}
	}
}
