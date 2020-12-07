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
	part1(input)
	part2(input)
}

func part1(input string) {
	var min, max, res int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		fmt.Sscanf(parts[0], "%d-%d", &min, &max)
		c := strings.Trim(parts[1], ":")
		pass := parts[2]

		count := strings.Count(pass, c)

		if count >= min && count <= max {
			res++
		}
	}
	fmt.Println(res)
}

func part2(input string) {
	var min, max, res int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		fmt.Sscanf(parts[0], "%d-%d", &min, &max)
		c := parts[1][0]
		pass := parts[2]

		if (pass[min-1] == c) != (pass[max-1] == c) {
			res++
		}
	}
	fmt.Println(res)
}
