package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 4a")
	contents, _ := os.ReadFile("input.txt")
	// var boards [][]int
	data := strings.Split(string(contents), "\n")
	for i := 2; i < len(data); i += 6 {
		full_board := data[i : i+5]
		counter := 0
		counter_row := 0
		var row_array [5][5]int
		for _, row := range full_board {
			row_data := strings.Split(row, " ")
			for _, value := range row_data {
				if len(value) != 0 {
					single_number, _ := strconv.Atoi(value)
					row_array[counter][counter_row] = single_number
					counter_row += 1
				}
			}
			if counter_row == 5 {
				counter_row = 0
			}
			counter += 1
			fmt.Println("row_array", row_array)
			// append(boards, row_array)
		}
	}

}
