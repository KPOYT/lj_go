package lj_validate

import "log"

//lint:ignore U1000 unused example;
func exValidateHTML() {
	err := ValidateHTML("<html><body><p>Hello World</p></body></html>")
	if err != nil {
		log.Fatal(err)
	}
}
