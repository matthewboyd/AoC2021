package main

import (
	"fmt"
)

func main() {

	fmt.Println("Day 6b")
	input := []int{2, 1, 1, 1, 1, 1, 1, 5, 1, 1, 1, 1, 5, 1, 1, 3, 5, 1, 1, 3, 1, 1, 3, 1, 4, 4, 4, 5, 1, 1, 1, 3, 1, 3, 1, 1, 2, 2, 1, 1, 1, 5, 1, 1, 1, 5, 2, 5, 1, 1, 2, 1, 3, 3, 5, 1, 1, 4, 1, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 5, 1, 2, 1, 1, 1, 1, 5, 1, 1, 1, 1, 1, 5, 1, 1, 1, 4, 5, 1, 1, 3, 4, 1, 1, 1, 3, 5, 1, 1, 1, 2, 1, 1, 4, 1, 4, 1, 2, 1, 1, 2, 1, 5, 1, 1, 1, 5, 1, 2, 2, 1, 1, 1, 5, 1, 2, 3, 1, 1, 1, 5, 3, 2, 1, 1, 3, 1, 1, 3, 1, 3, 1, 1, 1, 5, 1, 1, 1, 1, 1, 1, 1, 3, 1, 1, 1, 1, 3, 1, 1, 4, 1, 1, 3, 2, 1, 2, 1, 1, 2, 2, 1, 2, 1, 1, 1, 4, 1, 2, 4, 1, 1, 4, 4, 1, 1, 1, 1, 1, 4, 1, 1, 1, 2, 1, 1, 2, 1, 5, 1, 1, 1, 1, 1, 5, 1, 3, 1, 1, 2, 3, 4, 4, 1, 1, 1, 3, 2, 4, 4, 1, 1, 3, 5, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 5, 3, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 4, 4, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 3, 1, 4, 1, 1, 2, 2, 2, 1, 1, 2, 1, 1}
	countOfFish := make(map[int]int)
	countOfFish[0], countOfFish[1], countOfFish[2], countOfFish[3], countOfFish[4], countOfFish[5], countOfFish[6], countOfFish[7], countOfFish[8] = 0, 0, 0, 0, 0, 0, 0, 0, 0
	for _, v := range input {
		if val, ok := countOfFish[v]; ok {
			fmt.Println(val)
			countOfFish[v] += 1
		} else {
			countOfFish[v] = 1
		}
	}
	fmt.Println(countOfFish)
	amountOfZeros := 0
	for i := 0; i < 256; i++ {
		fmt.Println("fish in 8", countOfFish[8])
		countOfFish[8] += amountOfZeros
		countOfFish[6] += amountOfZeros
		temp := countOfFish[7]
		if amountOfZeros > 0 {
			countOfFish[7] = countOfFish[8] - amountOfZeros
		} else {
			countOfFish[7] = countOfFish[8]
		}
		countOfFish[8] = amountOfZeros
		temp2 := countOfFish[6]

		countOfFish[6] = temp + amountOfZeros
		temp = countOfFish[5]
		if amountOfZeros > 0 {
			countOfFish[5] = temp2 - amountOfZeros
		} else {
			countOfFish[5] = temp2
		}
		// countOfFish[5] = temp2
		temp2 = countOfFish[4]
		countOfFish[4] = temp
		temp = countOfFish[3]
		countOfFish[3] = temp2
		temp2 = countOfFish[2]
		countOfFish[2] = temp
		temp = countOfFish[1]
		countOfFish[1] = temp2
		countOfFish[0] = temp
		amountOfZeros = countOfFish[0]
		fmt.Println("numberOfZeros", amountOfZeros)

	}
	fmt.Println(countOfFish)
	final_sum := 0
	for _, v := range countOfFish {
		final_sum += v
	}
	fmt.Println(final_sum)
}
