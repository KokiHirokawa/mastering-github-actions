package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(18)
	if result != "Even" {
		t.Errorf("expected = Even, actual %s", result)
	}
}