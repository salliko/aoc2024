package mullover

import (
	"log"
	"path/filepath"
	"testing"
)

func TestMull(t *testing.T) {
	testPathData, err := filepath.Abs("../../test/testdata/day3/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	expected := 161
	answer, err := Mull(testPathData)
	if err != nil {
		log.Fatal(err)
	}
	if answer != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d", answer, expected)
	}
}

func TestMullDontDo(t *testing.T) {
	testPathData, err := filepath.Abs("../../test/testdata/day3/test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	expected := 48
	answer, err := MullDontDo(testPathData)
	if err != nil {
		log.Fatal(err)
	}
	if answer != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d", answer, expected)
	}
}
