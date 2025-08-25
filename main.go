package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalf("failed to listen on port 42069: %v", err)
	}
	defer listener.Close()
	var wg sync.WaitGroup
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			wg.Wait()
			log.Fatalf("error occured while waiting for network connection: %v", err)
		}
		
		fmt.Println("connection accepted")
		wg.Add(1)
		
		go func(conn io.ReadCloser) {
			defer conn.Close()
			defer wg.Done()
			
			for line := range getLinesChannel(conn) {
				fmt.Printf("%s\n", line)
			}
			fmt.Println("connection closed")
		}(conn)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	c := make(chan string)

	go func(f io.ReadCloser, c chan<- string) {
		defer f.Close()
		defer close(c)
		
		buf := make([]byte, 8)
		var current strings.Builder
		for {
			n, err := f.Read(buf)
			if err != nil {
				if current.String() != "" {
					c <- current.String()
				}
				break
			}

			current.Write(buf[:n])
			s := strings.Split(current.String(), "\n")
			if len(s) == 2 {
				c <- s[0]
				current.Reset()
				current.WriteString(s[1])
			}
		}
	}(f, c)

	return c
}
