package main

import (
	"fmt"
	"log"
	"os"
)

func d1() {
	// read lines of file
	body, err := os.Open("./data/d1.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	buf := make([]byte, 1024)
	for {
		n, err := body.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		fmt.Printf("%s", buf[:n])
	}

}
