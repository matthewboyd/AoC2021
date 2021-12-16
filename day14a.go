package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	result_all := ""
	polymer_pairs := map[string]string{}
	polymere_template := populateDictionary(polymer_pairs)
	polymer_pairs_arr := strings.Split(polymere_template, "")
	for j := 0; j < 40; j++ {
		previous := ""
		for i := 0; i < len(polymer_pairs_arr)-1; i++ {
			test := polymer_pairs_arr[i : i+2]
			result := polymer_pairs[strings.Join(test, "")]
			if previous != test[0] {
				result_all += test[0] + result + test[1]
			} else {
				result_all += result + test[1]
			}
			previous = test[1]
		}
		fmt.Println("day" + strconv.Itoa(j))
		polymer_pairs_arr = strings.Split(result_all, "")
		result_all = ""
	}

	countOccurences := make(map[string]int)

	for _, val := range polymer_pairs_arr {
		if _, ok := countOccurences[val]; ok {
			countOccurences[val] += 1
		} else {
			countOccurences[val] = 0
		}
	}

	fmt.Println(countOccurences)
}

func populateDictionary(polymer_pairs map[string]string) string {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	counter := 0
	var polymer_template string
	for scanner.Scan() {
		if counter == 0 {
			polymer_template = scanner.Text()
		}
		if counter > 1 {
			line := strings.Split(scanner.Text(), " ")
			polymer_pairs[line[0]] = line[2]
		}
		counter++
	}
	return polymer_template
}
