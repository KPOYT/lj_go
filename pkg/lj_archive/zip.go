package lj_archive

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Unzip extracts a .zip archive into the specified destination directory.
// It returns a string slice of the filepaths for each extracted file, relative to the destination root.
// If any of the destination files or folders already exists, they will be overwritten assuming the process has write permissions to the destination.
//
// The destination parent will be named after the input .zip file.
// e.g. if the input .zip file is named "test.zip", and the destination is "/data/project", the final destination will be "/data/project/test".
func Unzip(source, destination string) ([]string, error) {

	source = filepath.Clean(source)
	destination = filepath.Clean(destination)

	z, err := zip.OpenReader(source)
	if err != nil {
		return nil, err
	}
	defer z.Close()

	err = os.MkdirAll(destination, os.ModePerm)
	if err != nil {
		return nil, err
	}

	var extractedFiles []string
	for _, f := range z.File {
		err := write(destination, f)
		if err != nil {
			return nil, err
		}

		extractedFiles = append(extractedFiles, f.Name)
	}

	return extractedFiles, nil
}

func write(destination string, f *zip.File) error {
	zz, err := f.Open()
	if err != nil {
		return err
	}
	defer zz.Close()

	path := filepath.Join(destination, f.Name)

	if f.FileInfo().IsDir() {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(f, zz)
		if err != nil {
			return err
		}
	}

	return nil
}
