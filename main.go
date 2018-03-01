package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
)

var rows int
var columns int
var fleet int
var rides int
var bonus int
var steps int

func main() {
	var file = flag.String("file", "", "File to read")
	flag.Parse()
	fileObject, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer fileObject.Close()
	parseInput(fileObject)
}

func parseInput(file io.Reader) {
	scanner := bufio.NewScanner(file)

	lineNumber := 0
	for scanner.Scan() {
		content := scanner.Text()
		if lineNumber == 0 {
			parseFirstRow(content)
		} else {
			parseRides(content)
		}

		lineNumber++
	}
}

func parseFirstRow(input string) {

}
