package reports

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func StringListToInt(numbers []string) ([]int, error) {
	listNums := make([]int, 0)

	for _, snum := range numbers {
		num, err := strconv.Atoi(string(snum))
		if err != nil {
			return nil, err
		}
		listNums = append(listNums, num)
	}

	return listNums, nil
}

func isSafeReport(numbers []int) bool {
	if numbers[0] == numbers[len(numbers)-1] {
		return false
	}

	if numbers[0] > numbers[len(numbers)-1] {
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] <= numbers[i+1] {
				return false
			}
			if int(math.Abs(float64(numbers[i])-float64(numbers[i+1]))) > 3 {
				return false
			}
		}
	} else {
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] >= numbers[i+1] {
				return false
			}
			if int(math.Abs(float64(numbers[i])-float64(numbers[i+1]))) > 3 {
				return false
			}
		}
	}

	return true
}

func CountSafeReports(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeReports := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers, err := StringListToInt(strings.Split(line, " "))
		if err != nil {
			return 0, err
		}

		if isSafeReport(numbers) {
			safeReports++
		}
	}
	return safeReports, nil
}

func CountSafeReportsTolerateASingleBadLevel(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeReports := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers, err := StringListToInt(strings.Split(line, " "))
		if err != nil {
			return 0, err
		}

		if isSafeReport(numbers) {
			safeReports++
		} else {
			for i := 0; i < len(numbers); i++ {
				recoverNum := slices.Clone(numbers)
				recoverNum = append(recoverNum[:i], recoverNum[i+1:]...)
				if isSafeReport(recoverNum) {
					safeReports++
					break
				}
			}
		}
	}
	return safeReports, nil
}
