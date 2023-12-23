package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func first_id(line string) int {
	for i, c := range line {
		if c >= 48 && c <= 57 {
			return i
		}

	}
	return 0
}

func last_id(line string) int {
	for fi := range line {
		i := len(line) - fi - 1
		c := line[i]
		if c > 47 && c < 58 {
			return i
		}

	}
	return 0

}

func d1() int {
	// read lines of file
	f, err := os.Open("./data/d1.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()
	b := bufio.NewReader(f)

	// lines := make([]string,1000,1500)

	res := 0
	for {
		line, err := b.ReadString('\n')
		if err != nil  {
			break

			// append(lines,line)
		}
		tens := int(line[first_id(line)]) - 48
		ones := int(line[last_id(line)]) - 48
		fmt.Printf("tens: %d, ones: %d\n", tens, ones)

		res += int(tens)*10 + int(ones)
		// break

	}
	return res

}
