package zip

import (
	"os"
	"path/filepath"
	"testing"
)

// The ziptest.zip (internal/test/data/zip/ziptest.zip) tree:
// ziptest.zip/
// ├── nested
// │   └── nested.txt
// └── top.txt
// 1 directory, 2 files

// This test will unzip the testzip.zip file into a temporary directory, and validate that the correct files have been extracted and placed correctly.

func TestUnzip(t *testing.T) {
	tmpDir := t.TempDir()

	files, err := Unzip("../../internal/test/data/zip/ziptest.zip", tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	expectedFiles := []string{
		"ziptest",
		"ziptest/top.txt",
		"ziptest/nested",
		"ziptest/nested/nested.txt",
	}
	var unzipOut bool
	for _, f := range expectedFiles {
		unzipOut = checkUnzipOutput(files, f)
	}

	if !unzipOut {
		t.Fatal("unzip output is not correct")
	}

	top := filepath.Join(tmpDir, "ziptest", "top.txt")
	if _, err := os.OpenFile(top, os.O_RDONLY, os.ModePerm); err != nil {
		t.Fatal(err)
	}

	nested := filepath.Join(tmpDir, "ziptest", "nested", "nested.txt")
	if _, err := os.OpenFile(nested, os.O_RDONLY, os.ModePerm); err != nil {
		t.Fatal(err)
	}

}

func checkUnzipOutput(files []string, path string) bool {
	for _, f := range files {
		if f == path {
			return true
		}
	}
	return false
}
