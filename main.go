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
	defer file.Close()

	buf := make([]byte, 8)
	var current strings.Builder
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed to read file: %v", err)
		}

		current.Write(buf[:n])
		s := strings.Split(current.String(), "\n")
		if len(s) == 2 {
			fmt.Printf("read: %s\n", s[0])
			current.Reset()
			current.WriteString(s[1])
		}
	}
	fmt.Printf("read: %s\n", current.String())
}
