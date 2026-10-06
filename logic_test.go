package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("Fyne")
	want := "Hello, Fyne!"
	if got != want {
		t.Errorf("Greeting() = %q, want %q", got, want)
	}
}

func TestSumRange(t *testing.T) {
	got := SumRange(1, 10)
	want := 55
	if got != want {
		t.Errorf("SumRange(1, 10) = %d, want %d", got, want)
	}
}
