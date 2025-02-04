package match_test

import (
	"bytes"
	"testing"

	"github.com/mq/packages/match"
)

func TestContainsLines_ReturnCorrect(t *testing.T) {
	input := bytes.NewBufferString("hilda\nhi\nhit me\n")
	output := &bytes.Buffer{}
	m, e := match.NewMatcher(
		match.WithInput(input),
		match.WithOutput(output),
	)
	if e != nil {
		t.Fatalf("cant create coiuntie: %v", e.Error())
	}
	want := 3
	list := m.Matches("hi")
	got := len(list)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
