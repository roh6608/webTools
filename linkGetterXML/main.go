package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

// A function to check the equality of string slices
func equalSlice(compare1 []string, compare2 []string) bool {

	if len(compare1) != len(compare2) {
		return false
	} else {
		i := 0

		for i < len(compare1) {
			if compare1[i] != compare2[i] {
				return false
			}

			i += 1
		}
	}

	return true

}

// A function that extracts the string between two XML tags
func getBetween(input []string, tag string) (list []string) {
	left := "<" + tag + ">"
	right := "</" + tag + ">"
	rx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))

	i := 0
	for i < len(input) {
		matches := rx.FindAllStringSubmatch(input[i], -1)

		if matches != nil {
			list = append(list, matches[0][1])
		}

		i += 1

	}
	return list
}

// A function that will add the stem of the URL
func addStem(list []string, stem string) []string {
	links := make([]string, len(list))

	i := 0

	for i < len(list) {
		links[i] = stem + list[i]
		i += 1
	}
	return links
}

// A function to read XML
func readXml(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// A function to write the output to a file
func writeFile(filepath string, list []string) {
	f, err := os.Create(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	i := 0
	for i < len(list) {
		_, err2 := f.WriteString(list[i] + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}
		i += 1
	}
	fmt.Println("File written successfully to " + filepath)

}

func main() {
	outputPointer := flag.Bool("o", false, "If the output goes to a file, else will print")
	inputPointer := flag.String("i", "./testFiles/plant_catalog.xml", "The input filepath for the XML file")
	tagPointer := flag.String("t", "COMMON", "The tag to search between")
	stemPointer := flag.String("s", "", "The stem of the url")

	flag.Parse()

	links := addStem(getBetween(readXml(*inputPointer), *tagPointer), *stemPointer)

	i := 0

	if *outputPointer == true {
		writeFile("output.txt", links)
	} else {
		for i < len(links) {
			fmt.Println(links[i])
			i += 1
		}

	}

}
