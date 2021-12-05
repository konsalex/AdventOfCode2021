package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type commands string

const (
	forward commands = "forward"
	down    commands = "down"
	up      commands = "up"
)

func executeCommand(command string, currentPosition *Position) {

	splitted := strings.Split(command, " ")

	commandStripped := splitted[0]
	value, err := strconv.Atoi(splitted[1])
	if err != nil {
		panic(err)
	}

	switch commandStripped {
	case string(forward):
		currentPosition.x += value
	case string(down):
		currentPosition.y += value
	case string(up):
		currentPosition.y -= value
	default:
		panic("Unknown command")
	}

}

func executeAdvancedCommand(command string, currentPosition *AugmentedPosition) {

	splitted := strings.Split(command, " ")

	commandStripped := splitted[0]
	value, err := strconv.Atoi(splitted[1])
	if err != nil {
		panic(err)
	}

	switch commandStripped {
	case string(forward):
		currentPosition.x += value
		currentPosition.y += (currentPosition.aim * value)
	case string(down):
		currentPosition.aim += value
	case string(up):
		currentPosition.aim -= value
	default:
		panic("Unknown command")
	}

}

type Position struct {
	x int
	y int
}

type AugmentedPosition struct {
	x   int
	y   int
	aim int
}

func D2P1() {
	file, err := os.Open("./data/day2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	position := Position{0, 0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		executeCommand(scanner.Text(), &position)
	}
	fmt.Println("[D2 - Part 1]: Position X*Y:", position.x*position.y)
}

func D2P2() {
	file, err := os.Open("./data/day2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	position := AugmentedPosition{0, 0, 0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		executeAdvancedCommand(scanner.Text(), &position)
	}
	fmt.Println("[D2 - Part 2]: Position X*Y:", position.x*position.y)
}
