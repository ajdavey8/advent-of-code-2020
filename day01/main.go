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

	var nums []int
	for _, line := range strings.Split(input, "\n") {
		var n int
		fmt.Sscanf(line, "%d", &n)
		nums = append(nums, n)
	}

	part1(nums)
	part2(nums)
}

func part1(nums []int) {

	var res int
	for i, n := range nums {
		for _, n2 := range nums[i+1:] {
			if n+n2 == 2020 {
				res = n * n2
			}
		}
	}
	fmt.Println(res)
}

func part2(nums []int) {
	var res int
	var sumMap map[int]int
	sumMap = make(map[int]int)

	for i, n := range nums {
		for _, n2 := range nums[i+1:] {
			if n+n2 < 2020 {
				sumMap[n+n2] = n * n2
			}
		}
	}
	for _, n := range nums {
		if v, ok := sumMap[2020-n]; ok {
			res = v * n
		}
	}
	fmt.Println(res)
}
