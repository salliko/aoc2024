package mullover

import (
	"os"
	"regexp"
	"strings"

	"github.com/salliko/aoc2024/internal/reports"
)

func countMull(data []byte) (int, error) {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	fVal := re.FindAll(data, -1)
	result := 0
	for _, val := range fVal {
		cleanStr := strings.Trim(string(val[4:]), ")")
		numbers, err := reports.StringListToInt(strings.Split(cleanStr, ","))
		if err != nil {
			return 0, err
		}
		result += numbers[0] * numbers[1]

	}
	return result, nil
}

func Mull(basePath string) (int, error) {
	data, err := os.ReadFile(basePath)
	if err != nil {
		return 0, err
	}

	result, err := countMull(data)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func MullDontDo(basePath string) (int, error) {
	data, err := os.ReadFile(basePath)
	if err != nil {
		return 0, err
	}

	re := regexp.MustCompile(`don't\(\)(?:.*?|\n)*?do\(\)|don't\(\)(?:.*?|\n)*?$`)
	values := re.FindAll(data, -1)
	for _, value := range values {
		data = []byte(strings.Replace(string(data), string(value), "", -1))
	}

	result, err := countMull(data)
	if err != nil {
		return 0, err
	}

	return result, nil
}
