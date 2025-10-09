package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(conn net.Conn) <-chan string {
	ch := make(chan string)
	go receiveLine(ch, conn)
	return ch
}

func receiveLine(ch chan string, conn io.ReadCloser) {
	b := make([]byte, 8)
	currentLine := ""
	for {
		n, err := conn.Read(b)
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
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalf("Failed to open port, %s", listener.Addr().String())
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Failed to open connection")
		}
		fmt.Printf("Connection established: %s\n", conn.LocalAddr())
		ch := getLinesChannel(conn)
		for line := range ch {
			fmt.Printf("%s\n", line)
		}
		fmt.Println("Connection closed")
		conn.Close()
	}
}
