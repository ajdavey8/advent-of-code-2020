package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
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

	var res int
	req := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, passport := range strings.Split(input, "\n\n") {
		for i, field := range req {
			if !strings.Contains(passport, field) {
				break
			} else if i == 6 {
				res++
			}
		}

	}
	fmt.Println(res)
}

func part2(input string) {

	var res int
	req := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, passport := range strings.Split(input, "\n\n") {
		for i, field := range req {
			if !strings.Contains(passport, field) {
				break
			} else if i == 6 && validate(passport) {
				res++
			}
		}

	}
	fmt.Println(res)
}

func validate(passport string) bool {
	parts := map[string]string{}

	for _, field := range strings.Fields(passport) {
		part := strings.Split(field, ":")
		parts[part[0]] = part[1]
	}

	if !validateRange(parts["byr"], 1920, 2002) {
		return false
	}

	if !validateRange(parts["iyr"], 2010, 2020) {
		return false
	}
	if !validateRange(parts["eyr"], 2020, 2030) {
		return false
	}
	if hcl, _ := regexp.MatchString("#[0-9a-f]{6}", parts["hcl"]); !hcl {
		return false
	}

	if ecl, _ := regexp.MatchString("\\b(amb|blu|brn|gry|grn|hzl|oth)\\b", parts["ecl"]); !ecl {
		return false
	}
	if pid, _ := regexp.MatchString("^[0-9]{9}$", parts["pid"]); !pid {
		return false
	}

	if strings.HasSuffix(parts["hgt"], "cm") {
		if !validateRange(strings.TrimSuffix(parts["hgt"], "cm"), 150, 193) {
			return false
		}
	}

	if strings.HasSuffix(parts["hgt"], "in") {
		if !validateRange(strings.TrimSuffix(parts["hgt"], "in"), 59, 76) {
			return false
		}
	}
	return true
}

func validateRange(value string, min, max int) bool {
	num, err := strconv.Atoi(value)
	return err == nil && num >= min && num <= max
}
