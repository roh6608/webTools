package main

import (
	"testing"
)

func TestGetBetween(t *testing.T) {
	inputData := make([]string, 2)
	inputData[0] = "<Content>Line 1</Content>"
	inputData[1] = "<Content>Line 2</Content>"

	testWant := make([]string, 2)
	testWant[0] = "Line 1"
	testWant[1] = "Line 2"

	testGot := getBetween(inputData, "Content")

	if testGot[0] != testWant[0] || testGot[1] != testWant[1] {
		t.Errorf("getBetween was incorrect, got: %s, wanted: %s", testGot, testWant)
	}

}

// write more wide ranging tests utilising structs and loops to test more cases
// Currently function works if all lines have the tag being searched for, however doesnt if other tags are present
// have current test as first in the struct and then write in more tests as they are needed
