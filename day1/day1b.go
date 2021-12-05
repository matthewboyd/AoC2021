package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("couldn't open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	increased := 0
	total := 0
	var a string
	var b string
	var c string
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		switch counter {
		case 0:
			a = scanner.Text()
			counter += 1
			break
		case 1:
			b = scanner.Text()
			counter += 1
			break
		case 2:
			c = scanner.Text()
			counter += 1
			break
		case 3:
			fmt.Println("in here")
			if scanner.Text() == "" {
				return
			}
			a_int, _ := strconv.Atoi(a)
			b_int, _ := strconv.Atoi(b)
			c_int, _ := strconv.Atoi(c)

			if total != 0 {
				if total < a_int+b_int+c_int {
					fmt.Println("in the increase")
					increased += 1
				}
			}
			total = a_int + b_int + c_int
			a = b
			b = c
			c = scanner.Text()
			fmt.Println("a =", c, "b =", b, "c =", c)
			counter = 3
			fmt.Println("total = ", total)
			break
		}
	}

	a_int, _ := strconv.Atoi(a)
	b_int, _ := strconv.Atoi(b)
	c_int, _ := strconv.Atoi(c)
	if total < a_int+b_int+c_int {
		increased += 1
	}
	total = a_int + b_int + c_int
	fmt.Println(total)
	// b_int, _ := strconv.Atoi(b)
	// c_int, _ := strconv.Atoi(c)
	// if total < a_int+b_int+c_int {
	// fmt.Println("in the increase")
	// increased += 1
	// fmt.Println("total = ", total)
	// }
	fmt.Println("amount increased", increased)
}
