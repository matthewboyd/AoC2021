package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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

	for scanner.Scan() {
		stringArr := strings.Split(scanner.Text(), "")
		var results resultsArr
		for char := range stringArr {
			if stringArr[char] == "(" || stringArr[char] == "[" || stringArr[char] == "<" || stringArr[char] == "{" {
				results.Push(stringArr[char])
			} else {
				err := results.Pop(stringArr[char], allowed_mappings)
				if err != nil {
					(*&score).AddPoints(stringArr[char])
					break
				}
			}
		}
	}
	fmt.Println(score)

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

func (p *points) AddPoints(char string) {
	switch char {
	case "}":
		*p += 1197
	case ")":
		*p += 3
	case ">":
		*p += 25137
	case "]":
		*p += 57
	}
}
