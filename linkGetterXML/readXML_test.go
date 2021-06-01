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

// write this test properly as getBetween
