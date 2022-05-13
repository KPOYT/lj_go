package lj_validate

import "testing"

func TestValidateHTML(t *testing.T) {
	
	// good html
	err := ValidateHTML("<html><body><p>Hello World</p></body></html>")
	if err != nil {
		t.Error(err)
	}

	// bad html
	err = ValidateHTML("<html<body><p>Hello World</p></body></html>")
	if err == nil {
		t.Error(err)
	}

}