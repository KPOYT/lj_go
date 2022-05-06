package json

import "log"

type jsonStruct struct {
	Test string `json:"test"`
	Names []string `json:"names"`
}

//lint:ignore U1000 unused example;
func exDecodeTo() {

	j := jsonStruct{}

	err := DecodeTo(&j, "../../internal/test/data/json/test.json")
	if err != nil {
		log.Fatal(err)
	}
}