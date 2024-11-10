package greetings

import (
	"regexp"
	"testing"
)


func TestGreet(t *testing.T){
    name := "Gael"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg,err := Greet("Gael")

    if !want.MatchString(msg) || err != nil{
        t.Fatalf(`Hello("Gael") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestGreetEmpty(t *testing.T) {
    msg, err := Greet("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
