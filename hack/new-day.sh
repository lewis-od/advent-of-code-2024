#!/usr/bin/env bash

set -euo pipefail

if [ "$#" -ne 1 ]; then
    echo "usage: $0 <day num>"
    echo "    e.g: $0 day02"
    exit 1
fi

if [ -d "cmd/$1/input" ]; then
    echo "$1 already exists"
    exit 2
fi

mkdir -p "cmd/$1/input" "internal/$1"
touch "cmd/$1/input/example.txt" "cmd/$1/input/full.txt"

cat <<EOF > "cmd/$1/main.go"
package main

import (
	"fmt"
	"github.com/lewis-od/aoc24/internal/common"
	"github.com/lewis-od/aoc24/internal/$1"
)

func main() {
	input := common.ReadInputFileLines()
	result := $1.Part1(input)
	fmt.Printf("Part 1: %d\n", result)
}
EOF
echo "Created cmd/$1/main.go"

cat <<EOF > "internal/$1/main.go"
package $1

func Part1(input []string) int {
    return 0
}
EOF
echo "Created internal/$1/main.go"

cat <<EOF > "internal/$1/main_test.go"
package $1_test

import (
	"testing"

	"github.com/lewis-od/aoc24/internal/$1"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
    input := []string{}

    result := $1.Part1(input)

    require.Equal(t, 0, result)
}
EOF
echo "Created internal/$1/main_test.go"

echo "Done!"
