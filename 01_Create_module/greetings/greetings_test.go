/* use `go test -v` to test and show testing info in Terminal. */
package greetings

import (
	"regexp"
	"testing"
)

// Use testing pkg to call greetings.Hello with a name,\
// checking for if valid.
func TestHelloName(t *testing.T) {
	name := "Darling"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Darling") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// Test Hello Empty input calls greetings.Hello with an empty string.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
