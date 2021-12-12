package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file_input, _ := os.Open("input.txt")
	defer file_input.Close()
	scanner := bufio.NewScanner(file_input)
	// sum := 0
	// optionally, resize scanner's capacity for lines over 64K, see next example
	sum_total := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		input := scanner.Text()
		input_split := strings.Split(input, "|")
		real_input := strings.Split(strings.TrimRight(input_split[0], " "), " ")
		real_code_to_decode := strings.Split(strings.TrimLeft(input_split[1], " "), " ")
		input_map_known := make(map[string]string)
		input_map_unknown := make(map[string]int)
		for _, val := range real_input {
			switch {
			case len(val) == 2:
				input_map_known["1"] = val
			case len(val) == 3:
				input_map_known["7"] = val
			case len(val) == 7:
				input_map_known["8"] = val
			case len(val) == 4:
				input_map_known["4"] = val
			default:
				input_map_unknown[val] = len(val)
			}
		}

		for i, val := range input_map_unknown {
			for index, value := range input_map_known {
				if index == "7" && val == 6 {
					split_value := strings.Split(value, "")
					for iterator := range split_value {
						if !strings.Contains(i, split_value[iterator]) {
							input_map_known["6"] = i
							delete(input_map_unknown, i)
						}

					}
				}
				if index == "1" && val == 5 {
					split_value := strings.Split(value, "")
					contains := false
					for iterator := range split_value {
						if !strings.Contains(i, split_value[iterator]) {
							contains = true
							break
						}
					}
					if contains == false {
						input_map_known["3"] = i
						delete(input_map_unknown, i)
					}
				}
			}
		}
		number_left_out := ""
		for _, val := range input_map_known["4"] {
			if strings.Index(input_map_known["3"], string(val)) == -1 {
				number_left_out = string(val)
			}
		}
		for i, val := range input_map_unknown {
			if val == 5 && strings.Contains(i, number_left_out) {
				input_map_known["5"] = i
				delete(input_map_unknown, i)
			}
			if val == 5 && !strings.Contains(i, number_left_out) {
				input_map_known["2"] = i
				delete(input_map_unknown, i)
			}
		}

		for _, val := range input_map_known["6"] {
			if strings.Index(input_map_known["5"], string(val)) == -1 {
				number_left_out = string(val)
			}
		}

		for i, _ := range input_map_unknown {
			if strings.Contains(i, number_left_out) {
				input_map_known["0"] = i
				delete(input_map_unknown, i)
			}
			if !strings.Contains(i, number_left_out) {
				input_map_known["9"] = i
				delete(input_map_unknown, i)
			}
		}
		sum_per_line := ""
		for _, val := range real_code_to_decode {
			if len(val) == 2 {
				fmt.Println("digit 1")
				sum_per_line += "1"
			} else if len(val) == 7 {
				fmt.Println("digit 8")
				sum_per_line += "8"
			} else if len(val) == 3 {
				fmt.Println("digit 7")
				sum_per_line += "7"
			} else if len(val) == 4 {
				fmt.Println("digit 4")
				sum_per_line += "4"
			} else if len(val) == 5 {
				contains := 0
				for _, value := range input_map_known["3"] {
					if !strings.Contains(val, string(value)) {
						break
					} else {
						contains += 1
					}
				}
				if contains == 5 {
					fmt.Println("digit 3")
					sum_per_line += "3"
				} else {
					contains := 0
					for _, value := range input_map_known["2"] {
						if !strings.Contains(val, string(value)) {
							break
						} else {
							contains += 1
						}
					}
					if contains == 5 {
						fmt.Println("digit 2")
						sum_per_line += "2"
					} else {
						contains := 0
						for _, value := range input_map_known["5"] {
							if !strings.Contains(val, string(value)) {
								break
							} else {
								contains += 1
							}
						}
						if contains == 5 {
							fmt.Println("digit 5")
							sum_per_line += "5"
						}
					}
				}
			} else if len(val) == 6 {
				fmt.Println("number 6:/9", input_map_known["9"])
				contains := 0
				for _, value := range input_map_known["0"] {
					if !strings.Contains(val, string(value)) {
						break
					} else {
						contains += 1
					}
				}
				if contains == 6 {
					fmt.Println("digit 0")
					sum_per_line += "0"
				} else {
					contains := 0
					for _, value := range input_map_known["9"] {
						if !strings.Contains(val, string(value)) {
							break
						} else {
							contains += 1
						}
					}
					if contains == 6 {
						fmt.Println("digit 9")
						sum_per_line += "9"
					} else {
						contains := 0
						for _, value := range input_map_known["6"] {
							if !strings.Contains(val, string(value)) {
								break
							} else {
								contains += 1
							}
						}
						if contains == 6 {
							fmt.Println("digit 6")
							sum_per_line += "6"
						}
					}

					//do nothing
				}
			}

		}
		fmt.Println("sum per line", sum_per_line)
		number, _ := strconv.Atoi(sum_per_line)
		sum_total += number
		fmt.Println("sum", sum_total)
	}
}
