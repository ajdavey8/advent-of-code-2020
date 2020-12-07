package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	file, error := ioutil.ReadFile("input.txt")
	if error != nil {
		log.Fatal(error)
	}
	input := (string(file))
	// part1(input)
	part2(input)
}

func part1(input string) {
	res := checkNumOfTrees(input, 3, 1)
	fmt.Println(res)
}

func part2(input string) {
	res := checkNumOfTrees(input, 1, 1)
	res *= checkNumOfTrees(input, 3, 1)
	res *= checkNumOfTrees(input, 5, 1)
	res *= checkNumOfTrees(input, 7, 1)
	res *= checkNumOfTrees(input, 1, 2)
	fmt.Println(res)
}

func checkNumOfTrees(input string, right int, down int) int {
	numberOfTrees := 0
	rows := strings.Split(input, "\n")
	for i := 0; i < len(rows); i += down {
		j := (i / down * right) % len(rows[i])
		if rows[i][j] == '#' {
			numberOfTrees++
		}
	}
	return numberOfTrees
}
