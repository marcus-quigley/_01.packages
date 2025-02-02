package greeting_test

import (
	"bytes"
	"errors"
	"testing"
	"testing/iotest"

	greeting "github.com/mq/packages"
)

// func TestDisplayName(t *testing.T) {

// 	want := "hello, marcus"
// 	buf := &bytes.Buffer{}
// 	g := greeting.New()
// 	g.Output = buf
// 	g.Name = "marcus"

// 	g.Greet()

// 	got := buf.String()
// 	if got != want {
// 		t.Errorf("got %s, want %s", got, want)
// 	}
// }

// func TestReadName(t *testing.T) {
// 	want := "marcus"
// 	g := greeting.New()
// 	g.Input = bytes.NewBufferString("marcus\n")
// 	g.ReadName()

// 	got := g.Name
// 	if got != want {
// 		t.Errorf("got %s, want %s", got, want)
// 	}
// }

func TestGreetUser_PromptsUserForANameAndRendersGreeting(t *testing.T) {
	t.Parallel()
	input := bytes.NewBufferString("Greg")
	output := new(bytes.Buffer)
	greeting.GreetUser(output, input)
	got := output.String()
	want := "name, fucko?\nhello, Greg\n"
	if want != got {
		t.Fatalf("wanted %q but got %q", want, got)
	}
}

func TestGreetUser_PrintsHelloYouOnReadError(t *testing.T) {
	t.Parallel()
	input := iotest.ErrReader(errors.New("bad reader"))
	output := new(bytes.Buffer)
	greeting.GreetUser(output, input)
	got := output.String()
	want := "name, fucko?\nhello, \n"
	if want != got {
		t.Fatalf("wanted %q but got %q", want, got)
	}
}
