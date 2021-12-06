package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	starting_value := 0
	counter := 0
	increased := 0
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		if counter == 0 {
			starting_value = number
		} else {
			if number > starting_value {
				increased += 1
			}
		}
		counter++
		starting_value = number
	}
	fmt.Println(increased)
}
