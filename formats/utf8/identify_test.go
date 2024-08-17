package utf8

import (
	"bytes"
	"strings"
	"testing"
)

func TestIdentifyMatch(t *testing.T) {
	r := strings.NewReader("foo")

	match, err := Identify(r)

	if err != nil {
		t.Fatalf("Read error: %v", err)
	} else if !match {
		t.Fatalf("Doesn't match")
	}
}

func TestIdentifyNoMatch(t *testing.T) {
	r := bytes.NewReader([]byte{'f', 'o', 'o', ' ', 128})

	match, err := Identify(r)

	if err != nil {
		t.Fatalf("Read error: %v", err)
	} else if match {
		t.Fatal("Invalid UTF-8 not detected")
	}
}
