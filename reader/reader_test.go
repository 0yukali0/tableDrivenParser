package reader

import (
	"testing"
)

func TestNewFileReader(t *testing.T) {
	f := NewFileReader()
	if f != nil {
		t.Log("New siccess")
	}
}

func TestRead(t *testing.T) {
	f := NewFileReader()
	f.Read("input.txt")
	if f.err != nil {
		t.Error("File Read fail")
	} else {
		t.Log(f.context)
	}
}
