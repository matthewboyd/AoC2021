package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var input_array [10][10]int
	input_arr := populateArray(scanner, input_array)
	for i := 0; i < 1; i++ {
		passOverArray(&input_arr)
		duplicate_array := make(map[string]int)
		findAll9s(&input_arr, duplicate_array)
		findAll9s(&input_arr, duplicate_array)
		resetAll9s(&input_arr)
		fmt.Println(duplicate_array)
		fmt.Println(input_arr)

	}

}

func resetAll9s(input_arr *[10][10]int) {
	fmt.Println("in the reset all 9s")
	for index, h := range input_arr {
		for i, _ := range h {
			if input_arr[index][i] > 9 {
				input_arr[index][i] = 0
			}
		}
	}
}

func findAll9s(input_arr *[10][10]int, array_of_checked map[string]int) {
	fmt.Println("find all 9s")
	for index, h := range input_arr {
		for i, cell := range h {
			if cell > 9 {
				if _, ok := array_of_checked[strconv.Itoa(index)+":"+strconv.Itoa(i)]; !ok {
					fmt.Println("====== in the if statements ======")
					array_of_checked[strconv.Itoa(index)+":"+strconv.Itoa(i)] = 1

					//need to check if these locations exist
					if index-1 >= 0 && i-1 >= 0 {
						input_arr[index-1][i-1] = input_arr[index-1][i-1] + 1
						// array_of_checked[strconv.Itoa(index)+":"+strconv.Itoa(i)] = 1
					}
					if index-1 >= 0 && i+1 <= len(input_arr)-1 {
						input_arr[index-1][i+1] = input_arr[index-1][i+1] + 1
						// array_of_checked[strconv.Itoa(index-1)+":"+strconv.Itoa(i+1)] = 1

					}
					if index+1 <= len(input_arr)-1 && i-1 >= 0 {
						input_arr[index+1][i-1] = input_arr[index+1][i-1] + 1
						// array_of_checked[strconv.Itoa(index+1)+":"+strconv.Itoa(i-1)] = 1

					}
					if index+1 <= len(input_arr)-1 && i+1 <= len(input_arr)-1 {
						input_arr[index+1][i+1] = input_arr[index+1][i+1] + 1
						// array_of_checked[strconv.Itoa(index+1)+":"+strconv.Itoa(i+1)] = 1

					}
					if i+1 <= len(input_arr)-1 {
						input_arr[index][i+1] = input_arr[index][i+1] + 1
						// array_of_checked[strconv.Itoa(index)+":"+strconv.Itoa(i+1)] = 1

					}
					if i-1 >= 0 {
						input_arr[index][i-1] = input_arr[index][i-1] + 1
						// array_of_checked[strconv.Itoa(index)+":"+strconv.Itoa(i-1)] = 1

					}
					if index+1 <= len(input_arr)-1 {
						input_arr[index+1][i] = input_arr[index+1][i] + 1
						// array_of_checked[strconv.Itoa(index+1)+":"+strconv.Itoa(i)] = 1

					}
					if index-1 >= 0 {
						input_arr[index-1][i] = input_arr[index-1][i] + 1
						// array_of_checked[strconv.Itoa(index-1)+":"+strconv.Itoa(i)] = 1

					}
				}
			}
		}
	}
}

func passOverArray(input_arr *[10][10]int) {
	fmt.Println("In the passoverarray")
	for index, h := range input_arr {
		for i, cell := range h {
			input_arr[index][i] = cell + 1
		}
	}
}

func populateArray(scanner *bufio.Scanner, input_arr [10][10]int) [10][10]int {
	x_index := 0
	for scanner.Scan() {
		line := strings.Split(string(scanner.Text()), "")
		for char := range line {
			number, _ := strconv.Atoi(line[char])
			input_arr[x_index][char] = number
		}
		x_index++
	}
	return input_arr
}
