package main

import (
	"log"
	"path/filepath"

	mullover "github.com/salliko/aoc2024/internal/mull_over"
	"github.com/salliko/aoc2024/internal/reports"
)

func main() {
	testFilePath, err := filepath.Abs("test/testdata/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	countSafeReports, err := reports.CountSafeReportsTolerateASingleBadLevel(testFilePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(countSafeReports)

	testFilePath, err = filepath.Abs("test/testdata/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	coutMull, err := mullover.Mull(testFilePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("mull: %d", coutMull)

	testFilePath, err = filepath.Abs("test/testdata/day3/input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	coutMull, err = mullover.MullDontDo(testFilePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("mull dont: %d", coutMull)
}
