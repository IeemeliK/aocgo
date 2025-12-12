// Package day02 of aoc2025
package day02

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Part1(inputBuffer bufio.Reader) (any, error) {
	count := 0
	for {
		line, err := inputBuffer.ReadString(',')
		line = strings.TrimSuffix(line, ",")
		line = strings.TrimSpace(line)
		if err != nil && len(line) == 0 {
			break
		}
		first, last, ok := strings.Cut(line, "-")

		if !ok {
			log.Fatal("separator \"-\" not found in line")
		}

		if len(first)%2 != 0 && len(last)%2 != 0 {
			fmt.Printf("skipped line: %v\n", line)
			continue
		}

		start, err := strconv.Atoi(first)
		if err != nil {
			log.Fatal(err)
		}

		end, err := strconv.Atoi(last)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("line: %v, firstNum: %v, lastNum: %v\n", line, start, end)
		for num := start; num <= end; num++ {
			stringNum := strconv.Itoa(num)
			stringLen := len(stringNum) / 2
			if stringNum[:stringLen] == stringNum[stringLen:] {
				count += num
			}

		}
	}
	return count, nil
}

func Part2(inputBuffer bufio.Reader) (any, error) {
	count := 0
	for {
		line, err := inputBuffer.ReadString(',')
		line = strings.TrimSuffix(line, ",")
		line = strings.TrimSpace(line)
		if err != nil && len(line) == 0 {
			break
		}
		first, last, ok := strings.Cut(line, "-")

		if !ok {
			log.Fatal("separator \"-\" not found in line")
		}

		start, err := strconv.Atoi(first)
		if err != nil {
			log.Fatal(err)
		}

		end, err := strconv.Atoi(last)
		if err != nil {
			log.Fatal(err)
		}

		for num := start; num <= end; num++ {
			s := strconv.Itoa(num)
			if isPeriodic(s) {
				count += num
			}
		}
	}
	return count, nil
}

// This solution was made by AI, I initially had an alternative solution which worked but was MUCH slower (1s).
// My original solution broke the string into slices so that slice element had strings of length i given len(string) % i == 0
// and then compared the values in each slice
func isPeriodic(s string) bool {
	n := len(s)
	for k := 1; k <= n/2; k++ {
		if n%k != 0 {
			continue
		}

		ok := true
		for i := k; i < n && ok; i++ {
			if s[i] != s[i%k] {
				ok = false
			}
		}
		if ok {
			return true
		}
	}
	return false
}

func Solve(part int, inputBuffer bufio.Reader) (any, error) {
	switch part {
	case 1:
		return Part1(inputBuffer)
	case 2:
		return Part2(inputBuffer)
	default:
		return "", errors.New("invalid part input")
	}
}
