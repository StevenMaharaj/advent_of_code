package main

import (
	"bufio"
	"log"
	"os"
)

func last(line string, numLenMap *map[string]int, numMap *map[string]int) int {

	// keep track of i

	n := len(line)
	s := n - 1

	for s >= 0 {
		// CHECK first char

		c := line[s]
		if c >= 48 && c <= 57 {
			return int(c) - 48
		}

		// check letters
		for k, v := range *numLenMap {
			if s-v >= 0 {

				if string(line[s-v:s]) == k {
					return (*numMap)[k]
				}
			}

		}
		s--
	}

	// check if dight

	// check letters
	return 0
}

func first(line string, numLenMap *map[string]int, numMap *map[string]int) int {

	// keep track of i

	n := len(line)
	s := 0

	for s < n {
		// CHECK first char

		c := line[s]
		if c >= 48 && c <= 57 {
			return int(c) - 48
		}

		// check letters
		for k, v := range *numLenMap {
			if s+v <= n {

				if string(line[s:s+v]) == k {
					return (*numMap)[k]
				}
			}

		}
		s++
	}

	// check if dight

	// check letters
	return 0
}

func d2() int {

	numMap := make(map[string]int)
	numMap["one"] = 1
	numMap["two"] = 2
	numMap["three"] = 3
	numMap["four"] = 4
	numMap["five"] = 5
	numMap["six"] = 6
	numMap["seven"] = 7
	numMap["eight"] = 8
	numMap["nine"] = 9

	numLenMap := make(map[string]int)
	numLenMap["one"] = 3
	numLenMap["two"] = 3
	numLenMap["three"] = 5
	numLenMap["four"] = 4
	numLenMap["five"] = 4
	numLenMap["six"] = 3
	numLenMap["seven"] = 5
	numLenMap["eight"] = 5
	numLenMap["nine"] = 4

	f, err := os.Open("./data/d1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	res := 0
	b := bufio.NewReader(f)
	for {
		line, err := b.ReadString('\n')
		if err != nil {
			break
		}

		f := first(line, &numLenMap, &numMap)
		l := last(line, &numLenMap, &numMap)

		// fmt.Printf("%d%d\n",f,l)
		res += f*10 + l

	}

	return res
}
