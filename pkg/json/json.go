package json

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func DecodeTo(i interface{}, jsonPath string) error {
	jr, err := os.Open(jsonPath)
	if err != nil {
		return err
	}
	defer jr.Close()

	dec := json.NewDecoder(jr)
	dec.DisallowUnknownFields()

	err = dec.Decode(i)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("supplied multiple json objects in one file when only one is allowed")
	}

	return nil
}