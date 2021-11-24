package greetings

import (
	"regexp"
	"testing"
)

// test HelloName calls greetings.Hello with a name
// check for valid return values
func TestHelloName(t *testing.T) {
	name := "Cam"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Cam")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Cam") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// test helloName with an empty string
// check for valid return values
func TestEmptyHelloName(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
