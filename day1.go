package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type actions string

const (
	increased actions = "increased"
	decreased actions = "decreased"
)

func diff(val1 int, val2 int) actions {
	if val2 > val1 {
		return increased
	}
	return decreased
}

func findArraySum(arr *[]int, start int, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		sum += (*arr)[i]
	}
	return sum
}

func D1P1() {
	/*
		"In order to save Christmas, you'll need to get all fifty stars by December 25th."
	*/
	file, err := os.Open("./data/day1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var previous *int
	var counter int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if previous == nil {
			previous = &number
			continue
		}
		if diff(*previous, number) == increased {
			counter += 1
		}
		*previous = number
	}
	fmt.Println("[D1 - Part 1]: Increased counter:", counter)
}

func D1P2() {
	file, err := os.Open("./data/day1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var index int
	var counter int
	// Sliding window of 4 values as the middle value is always the same
	values := make([]int, 4)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if index < 4 {
			values[index%4] = number
		} else {
			// Shift the array (pop first)
			values = append(values[1:], number)
		}

		if index >= 3 {
			/** Now we can start compering the sliding windows */
			sum1 := findArraySum(&values, 0, 2)
			sum2 := findArraySum(&values, 1, 3)
			if sum2 > sum1 {
				counter += 1
			}
		}
		index += 1
	}
	fmt.Println("[D1 - Part 2]: Increased counter:", counter)
}
