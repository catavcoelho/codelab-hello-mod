package hellomod

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	want := "Hello, Gopher!"
	if got != want {
		t.Errorf("Greet(\"Gopher\") = %q; want %q", got, want)
	}
}

