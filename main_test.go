package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"
)

const (
	FILE_IN  = "test_in"
	FILE_OUT = "test_out"
)

func TestMain(t *testing.T) {
	in, err := os.Open(FILE_IN)
	if err != nil {
		t.Error(err)
	}
	reader = bufio.NewReader(in)
	buf := new(bytes.Buffer)
	writer = bufio.NewWriter(buf)
	main()
	out, err := os.ReadFile(FILE_OUT)
	if err != nil {
		t.Error(err)
	}
	if strings.TrimSpace(string(out)) != strings.TrimSpace(buf.String()) {
		t.Errorf("Expected:\n%s\n=============\nBut got:\n%s\n=============\n", out, buf.String())
	}
}
