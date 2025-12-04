package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/IeemeliK/aocgo/internal/year2025/day01"
)

var days = map[int]func(int, bufio.Reader) (string, error){
	1: day01.Solve,
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
	flag.Parse()

	solver, ok := days[*dayFlag]
	if !ok {
		log.Fatalf("day %d not implemented", *dayFlag)
	}

	inputPath := *inputFlag
	if inputPath == "" {
		inputPath = filepath.Join("internal",
			fmt.Sprintf("day%02d", *dayFlag),
			"input.txt",
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
