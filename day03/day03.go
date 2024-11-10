package day03

import (
	"bufio"
	"fmt"
	"os"
)

func RunPartOneSample() {
	result, err := runPartOne("day03/sample_input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Result", result)
}

func RunPartOne() {
	result, err := runPartOne("day03/input")
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

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		n := len(line)

		var ch rune
		for i := 0; i < n/2; i++ {
			found := false
			for j := n / 2; j < n; j++ {
				if line[i] == line[j] {
					ch = rune(line[i])
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		if ch >= 'a' && ch <= 'z' {
			sum = sum + int(ch-'a') + 1
		} else if ch >= 'A' && ch <= 'Z' {
			sum = sum + int(ch-'A') + 27
		}
	}

	return sum, nil
}

func RunPartTwoSample() {
	result, err := runPartTwo("day03/sample_input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Result", result)
}

func RunPartTwo() {
	result, err := runPartTwo("day03/input")
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

	sum := 0

	lines := make([]string, 3)
	idx := 0

	for scanner.Scan() {
		line := scanner.Text()

		lines[idx] = line
		idx = idx + 1

		if idx == 3 {
			m1 := make(map[rune]bool)

			for _, r := range lines[0] {
				m1[r] = true
			}

			m2 := make(map[rune]bool)
			for _, r := range lines[1] {
				if m1[r] {
					m2[r] = true
				}
			}

			var ch rune
			for _, r := range lines[2] {
				if m2[r] {
					ch = r
					break
				}
			}

			if ch >= 'a' && ch <= 'z' {
				sum = sum + int(ch-'a') + 1
			} else if ch >= 'A' && ch <= 'Z' {
				sum = sum + int(ch-'A') + 27
			}

			idx = 0
		}
	}

	return sum, nil
}
