package count

import (
	"bufio"
	"io"
)

type option func(*Countie)

type Countie struct {
	input io.Reader
}

func NewCountie(opts ...option) *Countie {
	c := &Countie{}
	for _, o := range opts {
		o(c)
	}
	return c
}

func WithInput(input io.Reader) option {
	return func(c *Countie) {
		c.input = input
	}
}

func (c Countie) Lines() int {
	i := 0
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		i += 1
	}
	return i
}
