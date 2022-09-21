package adder_test

import (
	"testing"

	"github.com/theghostmac/sonarqube-code-test/adder"
)
func TestAddNumbers(t *testing.T) {
	solution := adder.AddNumbers(6, 5)
	if solution != 11 {
		t.Error("result is incorrect: expected 5, got", solution)
	}
}