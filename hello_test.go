package main

import "testing"

func TestMain(t *testing.T) {
	if true {
		t.Error("expected false", true)
	}
}
