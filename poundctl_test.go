package poundctl

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseStatusXml(t *testing.T) {
	xml, err := ioutil.ReadFile("example.xml")
	if err != nil {
		t.Fatal(err)
	}

	status, err := ParseStatusXml(xml)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", status)
}
