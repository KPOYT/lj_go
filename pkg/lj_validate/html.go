package lj_validate

import (
	"encoding/xml"
	"io"
	"strings"
)

// ValidateHTML validates an HTML string for well-formedness.
func ValidateHTML(htmlToValidate string) error {
	r := strings.NewReader(htmlToValidate)
	d := xml.NewDecoder(r)

	d.Strict = false
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity
	for {
		_, err := d.Token()
		switch err {
		case io.EOF:
			return nil
		case nil:
		default:
			return err
		}
	}
}
