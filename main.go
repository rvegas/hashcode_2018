package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var rows int
var columns int
var fleet int
var totalRides int
var bonus int
var steps int

type Ride struct {
	start  Position
	finish Position
	early  int
	late   int
}

type Position struct {
	row    int
	column int
}

var rides []Ride

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

	log.Printf("Initial setup is: %d rows, %d cols, %d cars, %d totalRides, %d bonus, %d steps\n", rows, columns, fleet, totalRides, bonus, steps)
	log.Println("Individual rides are as follows:")
	log.Println(rides)
}

func parseFirstRow(input string) {
	parts := strings.Split(input, " ")
	rows, _ = strconv.Atoi(parts[0])
	columns, _ = strconv.Atoi(parts[1])
	fleet, _ = strconv.Atoi(parts[2])
	totalRides, _ = strconv.Atoi(parts[3])
	bonus, _ = strconv.Atoi(parts[4])
	steps, _ = strconv.Atoi(parts[5])
}

func parseRides(input string) {
	parts := strings.Split(input, " ")
	startRow, _ := strconv.Atoi(parts[0])
	startColumn, _ := strconv.Atoi(parts[1])
	endRow, _ := strconv.Atoi(parts[2])
	endColumn, _ := strconv.Atoi(parts[3])
	earlyStep, _ := strconv.Atoi(parts[4])
	lateStep, _ := strconv.Atoi(parts[5])

	startPosition := Position{
		row:    startRow,
		column: startColumn,
	}
	finishPosition := Position{
		row:    endRow,
		column: endColumn,
	}

	rides = append(rides, Ride{
		start:  startPosition,
		finish: finishPosition,
		early:  earlyStep,
		late:   lateStep,
	})
}

func distance(pointA, pointB Position) int {
	distCol := pointA.column - pointB.column
	distRow := pointA.row - pointB.row
	return int(math.Abs(float64(distCol) + float64(distRow)))
}
