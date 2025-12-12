package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/IeemeliK/aocgo/internal/year2025/day01"
	"github.com/IeemeliK/aocgo/internal/year2025/day02"
)

var days = map[int]func(int, bufio.Reader) (any, error){
	1: day01.Solve,
	2: day02.Solve,
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	dayFlag := flag.Int("day", 0, "day number 1-12")
	partFlag := flag.Int("part", 1, "part 1 or 2")
	inputFlag := flag.String("input", "", "optional input file (default: internal/dayXX/input.txt)")
	testFlag := flag.Bool("test", false, "determines wether to use test input or not")
	flag.Parse()

	solver, ok := days[*dayFlag]
	if !ok {
		log.Fatalf("day %d not implemented", *dayFlag)
	}

	inputPath := *inputFlag
	if inputPath == "" {
		inputFile := "input.txt"
		if *testFlag {
			inputFile = "input_test.txt"
		}

		inputPath = filepath.Join("internal",
			"year2025",
			fmt.Sprintf("day%02d", *dayFlag),
			inputFile,
		)
	}

	f, err := os.Open(inputPath)
	check(err)

	defer func() {
		err := f.Close()
		check(err)
	}()

	reader := bufio.NewReader(f)

	ans, err := solver(*partFlag, *reader)
	if err != nil {
		log.Fatalf("solver day %d part %d: %v", *dayFlag, *partFlag, err)
	}

	fmt.Println(ans)
}
