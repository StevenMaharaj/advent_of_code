package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func d1() {
	// read lines of file
	f, err := os.Open("./data/d1.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()
	b := bufio.NewReader(f)

	// lines := make([]string,1000,1500)

	for {
		line,err := b.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// append(lines,line)
		fmt.Println(line)
	}

}
