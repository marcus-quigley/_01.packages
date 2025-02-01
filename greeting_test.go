package greeting_test

import (
	"bytes"
	"testing"

	greeting "github.com/mq/packages"
)

func TestDisplayName(t *testing.T) {

	want := "hello, marcus"
	g := greeting.Greetie{
		Name: "marcus",
	}
	buf := &bytes.Buffer{}

	g.Greet(buf)

	got := buf.String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestReadName(t *testing.T) {
	want := "marcus"
	g := greeting.Greetie{}
	input := bytes.NewBufferString("marcus\n")
	g.ReadName(&bytes.Buffer{}, input)

	got := g.Name
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
