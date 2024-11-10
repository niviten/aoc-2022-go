package day02

import (
	"bufio"
	"fmt"
	"os"
)

func RunPartOneSample() {
	result, err := runPartOne("day02/sample_input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Result", result)
}

func RunPartOne() {
	result, err := runPartOne("day02/input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Result", result)
}

func runPartOne(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 3 {
			continue
		}

		a := string(line[0])
		b := string(line[2])

		score := 0

		if a == "A" {
			if b == "X" {
				score = score + 3 + 1
			} else if b == "Y" {
				score = score + 6 + 2
			} else if b == "Z" {
				score = score + 0 + 3
			}
		} else if a == "B" {
			if b == "X" {
				score = score + 0 + 1
			} else if b == "Y" {
				score = score + 3 + 2
			} else if b == "Z" {
				score = score + 6 + 3
			}
		} else if a == "C" {
			if b == "X" {
				score = score + 6 + 1
			} else if b == "Y" {
				score = score + 0 + 2
			} else if b == "Z" {
				score = score + 3 + 3
			}
		}

		totalScore = totalScore + score
	}

	return totalScore, nil
}

func RunPartTwoSample() {
	result, err := runPartTwo("day02/sample_input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Result", result)
}

func RunPartTwo() {
	result, err := runPartTwo("day02/input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Result", result)
}

func runPartTwo(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 3 {
			continue
		}

		a := string(line[0])
		b := string(line[2])

		score := 0

		if a == "A" {
			if b == "X" {
				score = score + 0 + 3
			} else if b == "Y" {
				score = score + 3 + 1
			} else if b == "Z" {
				score = score + 6 + 2
			}
		} else if a == "B" {
			if b == "X" {
				score = score + 0 + 1
			} else if b == "Y" {
				score = score + 3 + 2
			} else if b == "Z" {
				score = score + 6 + 3
			}
		} else if a == "C" {
			if b == "X" {
				score = score + 0 + 2
			} else if b == "Y" {
				score = score + 3 + 3
			} else if b == "Z" {
				score = score + 6 + 1
			}
		}

		totalScore = totalScore + score
	}

	return totalScore, nil
}
