package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunPartOneSample() {
	result, err := runPartOne("day01/sample_input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Result", result)
}

func RunPartOne() {
	result, err := runPartOne("day01/input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Result", result)
}

func runPartOne(filePath string) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var maxCal int
	var currentCal int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			if currentCal > maxCal {
				maxCal = currentCal
			}
			currentCal = 0
			continue
		}

		cal, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			return 0, err
		}

		currentCal = currentCal + int(cal)
	}

	if currentCal > maxCal {
		maxCal = currentCal
	}

	return int64(maxCal), nil
}

func RunPartTwoSample() {
	result, err := runPartTwo("day01/sample_input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Result", result)
}

func RunPartTwo() {
	result, err := runPartTwo("day01/input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Result", result)
}

func runPartTwo(filePath string) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	maxCalList := []int64{0, 0, 0}
	var currentCal int64

	insertToList := func(a int64) {
		if maxCalList[2] > a {
			return
		}
		maxCalList[2] = a
		if maxCalList[1] > a {
			return
		}
		maxCalList[1], maxCalList[2] = maxCalList[2], maxCalList[1]
		if maxCalList[0] > a {
			return
		}
		maxCalList[0], maxCalList[1] = maxCalList[1], maxCalList[0]
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			insertToList(currentCal)
			currentCal = 0
			continue
		}

		cal, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			return 0, err
		}

		currentCal = currentCal + int64(cal)
	}

	insertToList(currentCal)

	var sum int64
	for _, maxCalList := range maxCalList {
		sum = sum + maxCalList
	}

	return sum, nil
}
