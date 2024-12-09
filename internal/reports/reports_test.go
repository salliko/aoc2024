package reports

import (
	"log"
	"path/filepath"
	"testing"
)

func TestCountSafeReports(t *testing.T) {
	testPathData, err := filepath.Abs("../../test/testdata/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rightAnswer := 660
	answer, err := CountSafeReports(testPathData)
	if err != nil {
		log.Fatal(err)
	}
	if answer != rightAnswer {
		t.Errorf("Result was incorrect, got: %d, want: %d", answer, rightAnswer)
	}
}

func TestCountSafeReportsTolerateASingleBadLevel(t *testing.T) {
	testPathData, err := filepath.Abs("../../test/testdata/day2/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	expected := 4
	answer, err := CountSafeReportsTolerateASingleBadLevel(testPathData)
	if err != nil {
		log.Fatal(err)
	}
	if answer != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d", answer, expected)
	}
}
