package example

import (
	"testing"
)

func TestExample(t *testing.T) {
	hoge := Echo("hoge")

	if hoge != "hoge" {
		t.Fatal("failed test")
	}
}
