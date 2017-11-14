package main

import (
	"testing"
	"os"
)

func TestDefEnv(t *testing.T) {
	os.Setenv("ADDRESS", "localhost")
	if defEnv("ADDRESS", "argle") != "localhost" {
		t.Error("defEnv('ADDRESS') fail")
	}
	if defEnv("FOOBAR", "barfoo") != "barfoo" {
		t.Error("defEnv('FOOBAR') fail")
	}
}
