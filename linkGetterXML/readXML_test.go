package main

import (
	"testing"
)

func TestReadXML(t *testing.T) {
	xml := readXml("./testFiles/plant_catalog.xml")
	value := "<CATALOG>"

	if xml[1] != value {

		t.Errorf("Read was incorrect, got: %s, wanted %s", xml[1], value)
	}

}

// write the test with a struct for the answers so it can test a mutlitude of different files
