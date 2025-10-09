package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	UDPAddr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal("Could not resolve UDP address")
	}
	conn, err := net.DialUDP("udp", nil, UDPAddr)
	if err != nil {
		log.Fatal("Could not open UDP connection")
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error parsing string")
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("error writing ti UDP connection")
		}
	}
}
