package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strings"
)

func readFile(filename string) string {
	// Read the text file with the given filename;
	// return a list of the lines of text in the file.

	// Note: can just use os.ReadFile(filename)!
	// os.ReadFile(name string) ([]byte, error)

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	buff := bytes.NewBuffer(nil)

	_, err = io.Copy(buff, file)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	return buff.String()

}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func getWordFromLineList(text string) []string {
	clean := nonAlphanumericRegex.ReplaceAllString(text, "")
	return strings.Split(strings.ToLower(clean), " ")
}

func countFrequency(words []string) map[string]int {
	count := make(map[string]int)

	for _, word := range words {
		tally, ok := count[word]
		if !ok {
			count[word] = 1
		} else {
			count[word] = tally + 1
		}
	}
	return count
}

func innerProduct(d1, d2 map[string]int) float64 {
	sum := 0

	for key1 := range d1 {
		value2, ok := d2[key1]
		if ok {
			sum += d1[key1] * value2
		}

	}
	return float64(sum)
}

func normSq(d1 map[string]int) float64 {
	return float64(innerProduct(d1, d1) * innerProduct(d1, d1))
}

func angle(d1, d2 map[string]int) float64 {

	num := innerProduct(d1, d2)
	den := math.Sqrt(normSq(d1) + normSq(d2))

	return math.Acos(num / den)

}
