package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . filename1 filename2")
		os.Exit(-1)
	}

	filname1, filename2 := os.Args[1], os.Args[2]

	s := readFile(filname1)
	t := getWordFromLineList(s)
	d2 := countFrequency(t)

	s1 := readFile(filename2)
	t1 := getWordFromLineList(s1)
	d1 := countFrequency(t1)

	val := angle(d1, d2)

	fmt.Printf("The commonality between the documents is %.4f.\n", val)

}
