package oshash

import (
	"os"
	"testing"
)

func fromFileTester(filename string, expected uint64, t *testing.T) {
	f, err := os.Open(filename)
	if err != nil {
		t.Error(err)
	}
	h, err := FromFile(f)
	if err != nil {
		t.Error(err)
	}
	if h != expected {
		t.Errorf("Expected %v, got %v", expected, h)
	}
}
func TestFromFile(t *testing.T) {
	fromFileTester("test/parrots.jpg", uint64(17604422328474205166), t)

	// Tests for the OpenSubtitles reference files. These files are not
	// included in the GitHub repository due to their large filesize, and so
	// are commented out here, but you can download the files (and extract
	// the 4 GB dummy.bin file from its RAR archive) for more thorough testing.
	//
	// http://www.opensubtitles.org/addons/avi/breakdance.avi
	// fromFileTester("test/breakdance.avi", uint64(10242414353417707026), t)
	//
	// http://www.opensubtitles.org/addons/avi/dummy.rar
	// fromFileTester("test/dummy.bin", uint64(7059239720196713467), t)
}

func TestFromFilepath(t *testing.T) {
	expected := uint64(17604422328474205166)
	h, err := FromFilepath("test/parrots.jpg")
	if err != nil {
		t.Error(err)
	}
	if h != expected {
		t.Errorf("Expected %v, got %v", expected, h)
	}
}

func TestFromBytes(t *testing.T) {
	f, err := os.Open("test/parrots.jpg")
	if err != nil {
		t.Error(err)
	}
	info, err := os.Stat("test/parrots.jpg")
	if err != nil {
		t.Error(err)
	}
	buf := make([]byte, info.Size())
	f.Read(buf)
	h, err := FromBytes(buf)
	if err != nil {
		t.Error(err)
	}
	expected := uint64(17604422328474205166)
	if h != expected {
		t.Errorf("Expected %v, got %v", expected, h)
	}
}

func TestSmallDataErrors(t *testing.T) {
	buf := []byte("just a bit of data")
	_, err := FromBytes(buf)
	if err != ErrDataTooSmall {
		t.Error(err)
	}
}
