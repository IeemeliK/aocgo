// Package day01 of aoc25
package day01

import (
	"bufio"
	"errors"
)

func Part1(inputBuffer bufio.Reader) (string, error) {
	return "", errors.New("not implemented")
}

func Part2(inputBuffer bufio.Reader) (string, error) {
	return "", errors.New("not implemented")
}

func Solve(part int, inputBuffer bufio.Reader) (string, error) {
	switch part {
	case 1:
		return Part1(inputBuffer)
	case 2:
		return Part2(inputBuffer)
	default:
		return "", errors.New("invalid part input")
	}
}
