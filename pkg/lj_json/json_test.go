package lj_json

import (
	"testing"
)

type correctTestStruct struct {
	Test  string   `json:"test"`
	Names []string `json:"names"`
}

type malformedTestStruct struct {
	Test  string   `json:"pest"`
	Names []string `json:"names"`
}

type extraTestStruct struct {
	Test  string   `json:"test"`
	Names []string `json:"names"`
	Extra string   `json:"extra"`
}

func TestDecodeTo(t *testing.T) {

	testJSONPath := "../../internal/test/data/json/test.json"

	c := correctTestStruct{}
	err := DecodeTo(&c, testJSONPath, 0)
	if err != nil {
		t.Fatal(err)
	}

	m := malformedTestStruct{}
	err = DecodeTo(&m, testJSONPath, 0)
	if err == nil {
		t.Fatal(err)
	}

	e := extraTestStruct{}
	err = DecodeTo(&e, testJSONPath, 0)
	if err != nil {
		t.Fatal(err)
	}
	if e.Extra != "" {
		t.Fatal("'extra' field should be empty")
	}

	// Intentionally not set .json extension to keep ide happy when writing bad JSON to the file.
	multiJSONPath := "../../internal/test/data/json/multi"

	multi := correctTestStruct{}
	err = DecodeTo(&multi, multiJSONPath, 0)
	if err.Error() != "../../internal/test/data/json/multi contains multiple json objects in one file when only one is allowed" {
		t.Fatal(err)
	}
}
