#!/usr/bin/env bash
set -eu

usage() {
	echo "Usage: $0 <year> <day>" >&2
	echo "Example: $0 2025 1" >&2
	exit 1
}

[ "$#" -eq 2 ] || usage

YEAR="$1"
DAY_RAW="$2"

# Validate year
case "$YEAR" in
[0-9][0-9][0-9][0-9]) ;;
*)
	echo "year must be a 4-digit number, e.g. 2025" >&2
	exit 1
	;;
esac

# Validate day is 1–25
case "$DAY_RAW" in
'' | *[!0-9]*)
	echo "day must be a number between 1 and 25" >&2
	exit 1
	;;
esac
if [ "$DAY_RAW" -lt 1 ] || [ "$DAY_RAW" -gt 25 ]; then
	echo "day must be between 1 and 25" >&2
	exit 1
fi

# Zero-pad day
DAY=$(printf "%02d" "$DAY_RAW")

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
ROOT="$(realpath "$SCRIPT_DIR/..")"
DIR="$ROOT/internal/year$YEAR/day$DAY"
FILE="$DIR/day$DAY.go"

if [ -e "$FILE" ]; then
	echo "file already exists: $FILE" >&2
	exit 1
fi

mkdir -p "$DIR"

cat >"$FILE" <<EOF
// Package day$DAY of aoc$YEAR
package day$DAY

import (
	"bufio"
	"errors"
)

func Part1(inputBuffer bufio.Reader) (any, error) {
	return "", errors.New("not implemented")
}

func Part2(inputBuffer bufio.Reader) (any, error) {
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
EOF

echo "Created $FILE"
