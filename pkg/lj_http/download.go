package lj_http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type downloadInformation struct {
	destFileName           string
	counter, contentLength int64
}

func (d *downloadInformation) Write(p []byte) (n int, err error) {
	d.counter += int64(len(p))
	if d.contentLength != 0 {
		percentage := d.counter * 100 / d.contentLength
		fmt.Printf("\rDownloading %v of %v bytes (%v%%) of %v", d.counter, d.contentLength, percentage, d.destFileName)
	} else {
		fmt.Printf("\rDownloading %v bytes of %v", d.counter, d.destFileName)
	}
	return int(d.counter), nil
}

// Downloads a file from a URL and saves it to a local file.
// Download progress is reported in bytes to stdout. If possible, the content length will be used to calculate the percentage completion.
//
// If an empty destFileName string is provided, the file name will be calculated from the base of the provided URL.
func Download(URL, destFolder, destFileName string) error {
	err := os.MkdirAll(destFolder, os.ModePerm)
	if err != nil {
		return err
	}

	if destFileName == "" {
		destFileName = filepath.Base(URL)
	}

	writeFile := filepath.Join(destFolder, destFileName)

	f, err := os.Create(writeFile)
	if err != nil {
		return err
	}

	client := &http.Client{}

	r, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code (%v) received from %v", resp.StatusCode, URL)
	}

	client.Timeout = 0

	di := downloadInformation{
		destFileName:  destFileName,
		counter:       0,
		contentLength: resp.ContentLength,
	}

	tee := io.TeeReader(resp.Body, &di)
	_, err = io.Copy(f, tee)
	if err != nil {
		return err
	}

	return nil
}
