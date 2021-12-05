package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countBits(line string, counter *[]int) {
	for i := 0; i < len(*counter); i++ {
		bit := string(line[i])
		switch bit {
		case "0":
			(*counter)[i]--
		case "1":
			(*counter)[i]++
		default:
			panic("Invalid bit")
		}
	}
}

func sliceToString(ints []int) string {
	stringVals := make([]string, len(ints))

	for ind, val := range ints {
		stringVals[ind] = strconv.Itoa(val)
	}

	return strings.Join(stringVals, "")
}

func binaryStringNot(binaryString string) string {
	var result string
	for _, bit := range binaryString {
		if bit == '1' {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}

func D3P1() {
	// powerConsumption = gammarRate * epsilonRate
	// gammaRate =  most common bit in the corresponding position of all numbers
	// epsilon =  least common bit in the corresponding position of all numbers

	file, err := os.Open("./data/day3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Will store a counter value for each bit potision
	// If >0 then the bit is most of the times 1
	// If <0 then the bit is most of the times 0
	// If it is 0 then we have tie, which is not specified in the problem
	var counter []int
	var length int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if counter == nil {
			// Initialize the counter with the length of the first line
			length = len(line)
			counter = make([]int, length)
		}
		countBits(line, &counter)
	}

	// Calculate the gamma and epsilon rates
	gamma := make([]int, length)
	// If we were confident that the data would fit certain types
	// we could use bit types and masks
	for i, value := range counter {
		if value > 0 {
			gamma[i] = 1
		} else {
			gamma[i] = 0
		}
	}

	binaryGamma := sliceToString(gamma)
	binaryEpsilon := binaryStringNot(binaryGamma)

	var gammaValue uint64
	var epsilonValue uint64

	if value, err := strconv.ParseUint(binaryGamma, 2, len(gamma)); err != nil {
		panic(err)
	} else {
		gammaValue = value
	}
	if value, err := strconv.ParseUint(binaryEpsilon, 2, len(gamma)); err != nil {
		panic(err)
	} else {
		epsilonValue = value
	}

	fmt.Println("[D3 - Part 1]: Power Consumption: ", epsilonValue*gammaValue)
}

func countBits2(keys map[string]string, position int) int {
	var counter int
	for key, _ := range keys {
		bit := string(key[position])
		switch bit {
		case "0":
			counter--
		case "1":
			counter++
		default:
			panic("Invalid bit")
		}
	}
	return counter
}

func deleteKeys(keys map[string]string, position int, unwanted string) {
	for _, key := range keys {
		if string(key[position]) == unwanted {
			delete(keys, key)
		}
	}
}

func D3P2() {
	file, err := os.Open("./data/day3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var counter []int
	var length int

	oxygenKeys := map[string]string{}
	co2Keys := map[string]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if counter == nil {
			// Initialize the counter with the length of the first line
			length = len(line)
			counter = make([]int, length)
		}
		// values = append(values, line)
		oxygenKeys[line] = line
		co2Keys[line] = line
	}

	var index int
	/*
		Oxygen Calculation
	*/
	for {
		positionSum := countBits2(oxygenKeys, index)
		if positionSum >= 0 {
			// Delete all keys that start with 0
			deleteKeys(oxygenKeys, index, "0")
		} else {
			// Delete all keys that start with 1
			deleteKeys(oxygenKeys, index, "1")
		}
		if len(oxygenKeys) == 1 {
			break
		}
		index++
	}

	/*
		CO2 Calculation
	*/
	index = 0
	for {
		positionSum := countBits2(co2Keys, index)
		if positionSum >= 0 {
			// Delete all keys that start with 1
			deleteKeys(co2Keys, index, "1")
		} else {
			// Delete all keys that start with 0
			deleteKeys(co2Keys, index, "0")
		}
		if len(co2Keys) == 1 {
			break
		}
		index++
	}
	var result uint64

	for key1 := range co2Keys {
		for key2 := range oxygenKeys {
			value1, err := strconv.ParseUint(key1, 2, len(key1))
			if err != nil {
				panic(err)
			}
			value2, err := strconv.ParseUint(key2, 2, len(key2))
			if err != nil {
				panic(err)
			}
			result = value1 * value2
		}
	}

	fmt.Println("[D3 - Part 2]: Oxygen * CO2:", result)
}
