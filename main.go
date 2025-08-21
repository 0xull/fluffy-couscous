package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./message.txt")
	if err != nil {
		log.Fatalf("failed top open file: %v", err)
	}
	
	for line := range getLineConnection(file) {
		fmt.Printf("read: %s\n", line)
	}
}

func getLineConnection(f io.ReadCloser) <-chan string {
	c := make(chan string)

	go func(f io.ReadCloser, c chan<- string) {
		defer f.Close()

		buf := make([]byte, 8)
		var current strings.Builder
		for {
			n, err := f.Read(buf)
			if err != nil {
				if current.String() != "" {
					c <- current.String()
				}
				close(c)
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
