package main

import "testing"

func assertMatch(t *testing.T, pattern, name string) {
	if !match(pattern, name) {
		t.Error(pattern, name)
	}
}

func assertNoMatch(t *testing.T, pattern, name string) {
	if match(pattern, name) {
		t.Error(pattern, name)
	}
}

func TestMatch(t *testing.T) {
	assertMatch(t, "*.go", "main.go")
	assertNoMatch(t, "*.go", "main.go2")
	assertMatch(t, "Dockerfile*", "Dockerfile.ruler")
}
