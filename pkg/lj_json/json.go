package lj_json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Decodes a json file to an interface. Interface fields must match the json to be decoded.
// Extra fields which are not included in the json file will be ignored.
//
// If the json file contains multiple objects, only the first one will be decoded.
//
// If the json file is empty, an error will be returned.
//
// maxReadSize is the maximum size in bytes of the json file to read.
// If the file is larger than this, an error is returned. Set to 0 to read the whole file regardless of its size.
func DecodeTo(i interface{}, jsonPath string, maxReadSize int64) error {
	jr, err := os.Open(jsonPath)
	if err != nil {
		return err
	}
	defer jr.Close()

	s, err := jr.Stat()
	if err != nil {
		return err
	}
	
	if s.Size() == 0 {
		return fmt.Errorf("%s is empty", jsonPath)
	}

	if maxReadSize > 0 && s.Size() > maxReadSize {
		return fmt.Errorf("%s is too large. file size is greater than %d", jsonPath, maxReadSize)
	}

	dec := json.NewDecoder(jr)
	dec.DisallowUnknownFields()

	err = dec.Decode(i)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return fmt.Errorf("%s contains multiple json objects in one file when only one is allowed", jsonPath)
	}

	return nil
}
