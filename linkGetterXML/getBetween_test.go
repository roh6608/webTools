package main

import (
	"testing"
)

func TestGetBetween(t *testing.T) {

	data := []struct {
		x []string
		y []string
	}{
		{[]string{"<Content>Line 1</Content>"}, []string{"Line 1"}},
		{[]string{"<Content>Line 1</Content>", "<Content>Line 2</Content>"}, []string{"Line 1", "Line 2"}},
		{[]string{"<Content>Line 1</Content>", "<NotContent>Don't Grab</NotContent>", "<Content>Line 3</Content>"},
			[]string{"Line 1", "Line 3"}},
	}

	for _, table := range data {
		testGot := getBetween(table.x, "Content")

		if equalSlice(testGot, table.y) == false {
			t.Errorf("getBetween was incorrect, got: %s, wanted: %s", testGot, table.y)
		}

	}
}

// write more wide ranging tests utilising structs and loops to test more cases
// Currently function works if all lines have the tag being searched for, however doesnt if other tags are present
// have current test as first in the struct and then write in more tests as they are needed
