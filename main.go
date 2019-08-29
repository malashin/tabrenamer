package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

// Takes tab separated input file (src	dst) and renames found files in the folder.

var inputPath = "input.txt"
var re = regexp.MustCompile(`^(.+)\t(.+)$`)

func main() {
	links, err := readLines(inputPath)
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err)
	}

	filesMap := make(map[string]bool)

	for _, f := range files {
		filesMap[f.Name()] = true
	}

	for _, link := range links {
		if !re.MatchString(link) {
			fmt.Println("ERROR: input file contains wrong string pattern")
			return
		}

		in := re.ReplaceAllString(link, "${1}")
		out := re.ReplaceAllString(link, "${2}")

		if _, ok := filesMap[in]; ok {
			fmt.Println("RENAMING: " + in + " > " + out)
			err = os.Rename(in, out)
			if err != nil {
				panic(err)
			}
		}
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
