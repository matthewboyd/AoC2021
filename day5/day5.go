package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("testing for day 5")
	contents, _ := os.Open("input.txt")
	defer contents.Close()

	var lines []string
	scanner := bufio.NewScanner(contents)
	for scanner.Scan() {
		line := strings.Replace(scanner.Text(), " -> ", ",", -1)
		lines = append(lines, line)
	}

	new_contents := lines
	results := []string{}

	for i := 0; i < len(new_contents); i++ {
		row_data := strings.Split(new_contents[i], ",")
		x1, _ := strconv.Atoi(row_data[0])
		x2, _ := strconv.Atoi(row_data[2])
		y1, _ := strconv.Atoi(row_data[1])
		y2, _ := strconv.Atoi((row_data[3]))
		if x1 == x2 || y1 == y2 {
			if x1 == x2 {
				if y1 > y2 {
					temp := y1
					y1 = y2
					y2 = temp
				}

				for i := y1; i <= y2; i++ {
					x_value := strconv.Itoa(x1)
					y_value := strconv.Itoa(i)
					results = append(results, x_value+"+"+y_value)
				}
			}
			if y1 == y2 {
				if x1 > x2 {
					temp := x1
					x1 = x2
					x2 = temp
				}
				for i := x1; i <= x2; i++ {
					x_value := strconv.Itoa(i)
					y_value := strconv.Itoa(y1)
					results = append(results, x_value+"+"+y_value)
				}
			}
		}
		if diff(x1, x2) == diff(y1, y2) {
			fmt.Println("in the diff", x1, y1, x2, y2)
			if x1 > x2 && y1 < y2 {
				y2_new := y2
				for i := x2; i <= x1; i++ {
					x_value := strconv.Itoa(i)
					y_value := strconv.Itoa(y2_new)
					// fmt.Println(x_value + "+" + y_value)
					results = append(results, x_value+"+"+y_value)
					y2_new -= 1
				}
			} else if x1 > x2 && y1 > y2 {
				y2_new := y2
				for i := x2; i <= x1; i++ {
					x_value := strconv.Itoa(i)
					y_value := strconv.Itoa(y2_new)
					// fmt.Println(x_value + "+" + y_value)
					results = append(results, x_value+"+"+y_value)
					y2_new += 1

				}
			} else if x1 < x2 && y1 < y2 {
				y1_new := y1
				for i := x1; i <= x2; i++ {
					x_value := strconv.Itoa(i)
					y_value := strconv.Itoa(y1_new)
					fmt.Println(x_value + "+" + y_value)
					results = append(results, x_value+"+"+y_value)
					y1_new += 1

				}
			} else if x1 < x2 && y1 > y2 {
				y1_new := y1
				for i := x1; i <= x2; i++ {
					x_value := strconv.Itoa(i)
					y_value := strconv.Itoa(y1_new)
					fmt.Println(x_value + "+" + y_value)
					results = append(results, x_value+"+"+y_value)
					y1_new -= 1

				}
			}
		}
	}
	dictionary := make(map[string]int)
	fmt.Println("results_arr", results)
	for _, item := range results {

		_, exists := dictionary[item]

		if exists {
			dictionary[item] += 1
		} else {
			dictionary[item] = 1
		}

	}
	results_counter := 0
	for _, value := range dictionary {
		if value > 1 {
			results_counter += 1
		}
	}
	fmt.Println("results", results_counter)

}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
