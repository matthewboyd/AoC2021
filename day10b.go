package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type resultsArr []string
type points int

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	allowed_mappings := make(map[string]string)
	allowed_mappings["["] = "]"
	allowed_mappings["("] = ")"
	allowed_mappings["<"] = ">"
	allowed_mappings["{"] = "}"
	var score points = 0
	var full_results []points

	for scanner.Scan() {
		skip := false

		score = 0
		stringArr := strings.Split(scanner.Text(), "")
		var results resultsArr
		for char := range stringArr {
			if stringArr[char] == "(" || stringArr[char] == "[" || stringArr[char] == "<" || stringArr[char] == "{" {
				results.Push(stringArr[char])
			} else {
				err := results.Pop(stringArr[char], allowed_mappings)
				if err != nil {
					skip = true
					break
				}
			}
		}
		fmt.Println("results is equal to", results)
		for test := range results {
			fmt.Println("ignore", test)
			if skip == true {
				break
			}
			results.PopNoCompare(&score, allowed_mappings)
		}
		if score != 0 {
			full_results = append(full_results, score)
		}
	}

	sort.Slice(full_results, func(i, j int) bool {
		return full_results[i] < full_results[j]
	})
	for _, v := range full_results {
		fmt.Println(v)
	}
	fmt.Println("score", full_results[len(full_results)/2])

}

func (r *resultsArr) Push(char string) {
	*r = append(*r, char)
}

func (r *resultsArr) Pop(char string, allowed_mappings map[string]string) (err error) {
	last_value := (*r)[len(*r)-1]
	if char != allowed_mappings[last_value] {
		return errors.New("Not correct")
	} else {
		*r = (*r)[:len(*r)-1]
		return nil
	}

}
func (r *resultsArr) PopNoCompare(p *points, allowed_mappings map[string]string) {
	last_value := (*r)[len(*r)-1]
	lenOfResult := len(*r)
	p.AddPoints(allowed_mappings[last_value], lenOfResult)
	*r = (*r)[:len(*r)-1]
}

func (p *points) AddPoints(char string, lenOfResults int) {
	switch char {
	case "}":
		*p += 3
	case ")":
		*p += 1
	case ">":
		*p += 4
	case "]":
		*p += 2
	}
	if lenOfResults != 1 {
		*p = *p * 5
	}
}
