package match

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

type Matcher struct {
	input  io.Reader
	output io.Writer
}

type option func(*Matcher) error

func NewMatcher(opts ...option) (*Matcher, error) {
	m := &Matcher{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, o := range opts {
		e := o(m)
		if e != nil {
			return nil, e
		}
	}
	return m, nil
}

func WithInput(r io.Reader) option {
	return func(m *Matcher) error {
		if r == nil {
			return errors.New("nil input reader")
		}
		m.input = r
		return nil
	}
}

func WithOutput(w io.Writer) option {
	return func(m *Matcher) error {
		if w == nil {
			return errors.New("nil output reader")
		}
		m.output = w
		return nil
	}
}

func (m Matcher) Matches(v string) []string {
	lines := []string{}
	scanner := bufio.NewScanner(m.input)
	for scanner.Scan() {
		l := scanner.Text()
		if strings.Contains(l, v) {
			lines = append(lines, l)
		}
	}
	return lines
}
