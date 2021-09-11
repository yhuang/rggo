package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4

	res := count(b, map[string]bool{})

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

// TestCountLines tests the count function set to count lines
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

	exp := 3

	res := count(b, map[string]bool{"lines": true})

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

// TestCountBytes tests the count function set to count bytes
func TestCountByte(t *testing.T) {
	b := bytes.NewBufferString("abcdefghijk")
	exp := 11

	res := count(b, map[string]bool{"bytes": true})

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
