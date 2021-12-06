package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Winner struct {
	sum         int
	bingoNumber int
	position    int
}

func main() {
	winners := make(map[int]Winner)
	winners_counter := 0
	contents, _ := os.ReadFile("input.txt")
	var boards [][]int
	data := strings.Split(string(contents), "\n")
	for i := 2; i < len(data); i += 6 {
		full_board := data[i : i+5]
		full_board_int := strings.Join(full_board, " ")
		full_board_int = strings.Replace(full_board_int, "  ", " ", -1)
		full_board_int_arr := strings.Split(full_board_int, " ")
		ints := make([]int, len(full_board_int_arr))

		for i, s := range full_board_int_arr {
			ints[i], _ = strconv.Atoi(s)
		}
		boards = append(boards, ints)
	}
	fmt.Println("boards", boards)
	bingoNumbers := strings.Split(data[0], ",")
	for index, number := range bingoNumbers {
		for i := 0; i < len(boards); i++ {
			number, _ := strconv.Atoi(number)
			fmt.Println("number", number, "index", index)
			winners_counter = checkNumber(boards[i], number, i, winners_counter, winners)
		}
	}

	fmt.Println("winnersArray", winners)

	for _, val := range winners {
		fmt.Println("val", val)
	}

}

func checkNumber(board []int, bingoNumber int, boardNumber int, winners_counter int, winners map[int]Winner) int {
	counter := 0
	fmt.Println("board in checkNumber", board)
	for index, items := range board {
		if items == bingoNumber {
			board[index] = -1
			counter += 1
		}
		if items == -1 {
			counter += 1
		}
		if checkForWin(board) {

			sum := 0
			for _, val := range board {
				if val != -1 {
					sum += val
				}
			}
			if _, ok := winners[boardNumber]; ok == false {
				fmt.Println(winners_counter)
				win := Winner{sum, bingoNumber, winners_counter}
				winners[boardNumber] = win
				winners_counter += 1
				return winners_counter
			}
		}
	}
	return winners_counter

}

func checkForWin(board []int) bool {
	total_vertically := 0
	for i := 0; i < 5; i++ {
		total_vertically = board[i] + board[i+5] + board[i+10] + board[i+15] + board[i+20]
		if total_vertically == -5 {
			return true
		}
	}
	total_horizontally := 0
	for i := 0; i < 25; i += 5 {
		total_horizontally = board[i] + board[i+1] + board[i+2] + board[i+3] + board[i+4]
		if total_horizontally == -5 {
			return true
		}
	}
	return false
}
