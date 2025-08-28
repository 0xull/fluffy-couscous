package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer conn.Close()
	
	buf := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := buf.ReadString('\n')
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		
		conn.Write([]byte(line))
	}
}