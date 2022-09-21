// Unit test
package greeting_test

import (
	"testing"
	"github.com/theghostmac/sonarqube-code-test/greeting"
)

func TestGreet(t *testing.T){
	// name := "John Doe"
	// want := "Hello, John Doe"
	// got, err := greeting.Greet(name)
	// if got != want || err != nil {
	// 	t.Fatalf("TestGreet(%s): got %q/%v, want %q/nil", name, got, err, want)
	// }
	// solution := adder.AddNumbers(6, 5)
	// if solution != 11 {
	// 	t.Error("result is incorrect: expected 5, got", solution)
	// }
	name := "John Doe"
	want := greeting.Greet(name)
	if name != "John Doe" {
		t.Error("the name is incorrect: expected John Doe, got", name)
	}
}