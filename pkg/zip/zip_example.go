package zip

import (
	"fmt"
	"log"
)

func exUnzip() {

	files, err := Unzip("path/to/zip.zip", "path/to/out/dir")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)
}