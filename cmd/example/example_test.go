package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	hoge := example("hoge")

	if hoge != "hoge" {
		t.Fatal("failed test")
	}
}
