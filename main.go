package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// Takes tab separated input file (src	dst) and renames found files in the folder.

var re = regexp.MustCompile(`^(.+)\t(.+)$`)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("usage: tabrenamer [filelist]")
		fmt.Println("tabrenamer takes tab separated input file (src\tdst) and renames found files in the folder.")
		os.Exit(0)
	}

	inputPath := os.Args[1]
	links, err := readLines(inputPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(links) < 1 {
		fmt.Println("ERROR: input file is empty")
		return
	}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		return
	}

	filesMap := make(map[string]bool)
	for _, f := range files {
		filesMap[f.Name()] = true
	}

	for _, link := range links {
		if strings.Count(link, "\t") > 1 {
			fmt.Println("ERROR: input file contains more then one tab per line")
			return
		}

		if !re.MatchString(link) && link != "" {
			fmt.Println("ERROR: input file contains wrong string pattern")
			fmt.Printf("%q\n", link)
			return
		}
	}

	for _, link := range links {
		if link == "" {
			continue
		}

		in := re.ReplaceAllString(link, "${1}")
		out := re.ReplaceAllString(link, "${2}")

		if _, ok := filesMap[in]; ok {
			fmt.Println("RENAMING: " + in + " > " + out)
			err = os.Rename(in, out)
			if err != nil {
				fmt.Println(err)
				return
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
