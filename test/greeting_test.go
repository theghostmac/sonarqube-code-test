package main_test
// Unit test

import (
	"testing"
	"github.com/theghostmac/sonarqube-code-test/app"
)

func TestGreet(t *testing.T){
	name := "John Doe"
	want := "Hello, John Doe"
	got, err := Greet(name)
	if got != want || err != nil {
		t.Fatalf("TestGreet(%s): got %q/%v, want %q/nil", name, got, err, want)
	}
}