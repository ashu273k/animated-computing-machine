package main

import "testing"

func TestFrames(t *testing.T) {
	got := frames()
	if len(got) == 0 {
		t.Fatal("frames() returned no frames")
	}
}
