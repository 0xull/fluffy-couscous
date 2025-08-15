package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./message.txt")
	if err != nil {
		log.Fatalf("failed top open file: %v", err)
	}
	defer file.Close()

	buf := make([]byte, 8)

	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}

		fmt.Printf("read: %s\n", string(buf[:n]))
	}
}
