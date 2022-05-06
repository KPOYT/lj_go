package download

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFromURL(t *testing.T) {
	URL := "https://raw.githubusercontent.com/lambdajack/go/main/internal/test/data/download/robots.txt"

	// TEST SPECIFIED BASE
	spDir := filepath.Join(t.TempDir(), "specified-base")

	err := FromURL(URL, spDir, "we_love_robots.txt")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.OpenFile(filepath.Join(spDir, "we_love_robots.txt"), os.O_RDONLY, os.ModePerm); err != nil {
		t.Fatal(err)
	}

	f, err := os.OpenFile(filepath.Join(spDir, "we_love_robots.txt"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var b [128]byte
	n, err := f.Read(b[:])
	if err != nil {
		t.Fatal(err)
	}
	if string(b[:n]) != "User-agent: *" {
		t.Fatalf("robots.txt is not correct; wanted 'User-agent: *', got %v", string(b[:n]))
	}

	time.Sleep(time.Second * 1)

	// TEST CALCULATED BASE
	clDir := filepath.Join(t.TempDir(), "calculated-base")

	err = FromURL(URL, clDir, "")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.OpenFile(filepath.Join(clDir, "robots.txt"), os.O_RDONLY, os.ModePerm); err != nil {
		t.Fatal(err)
	}

	f, err = os.OpenFile(filepath.Join(clDir, "robots.txt"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var c [128]byte
	n, err = f.Read(c[:])
	if err != nil {
		t.Fatal(err)
	}
	if string(c[:n]) != "User-agent: *" {
		t.Fatalf("robots.txt is not correct; wanted 'User-agent: *', got %v", string(b[:n]))
	}
}
