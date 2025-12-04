// Package day01 of aoc25
package day01

import (
	"bufio"
	"errors"
	"log"
	"strconv"
	"strings"
)

func Part1(inputBuffer bufio.Reader) (any, error) {
	dial := 50
	count := 0

	for {
		line, err := inputBuffer.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)

		direction := line[0]
		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case 'R':
			dial = (dial + rotation) % 100
		case 'L':
			dial = (dial - rotation + 100) % 100
		}

		if dial == 0 {
			count++
		}

	}

	return count, nil
}

func Part2(inputBuffer bufio.Reader) (string, error) {
	return "", errors.New("not implemented")
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
