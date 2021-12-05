package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("testing for day 5")
	contents, _ := os.ReadFile("input.txt")

	//using -1 to replace all occurrences
	new_contents := strings.Replace(string(contents), " -> ", ",", -1)
	rows_data := strings.Split(new_contents, "\n")
	results := []string{}
	for i := 0; i < len(rows_data)-1; i++ {
		row_data := strings.Split(rows_data[i], ",")
		x1, _ := strconv.Atoi(row_data[0])
		x2, _ := strconv.Atoi(row_data[2])
		y1, _ := strconv.Atoi(row_data[1])
		y2, _ := strconv.Atoi(row_data[3])

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
	}
	dictionary := make(map[string]int)

	for _, item := range results {

		_, exists := dictionary[item]

		if exists {
			dictionary[item] += 1
		} else {
			dictionary[item] = 1
		}

	}
	fmt.Println(dictionary)
	results_counter := 0
	for _, value := range dictionary {
		if value > 1 {
			results_counter += 1
		}
	}
	fmt.Println(results_counter)

}
