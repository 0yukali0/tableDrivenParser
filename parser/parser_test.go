package parser

import (
	"testing"
)

func TestNewParser(t *testing.T) {
	p := NewParser()
	if p != nil {
		t.Log("New parser success")
	}
}
