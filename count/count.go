package count

import (
	"bufio"
	"errors"
	"io"
)

type option func(*Countie) error

type Countie struct {
	input io.Reader
}

func NewCountie(opts ...option) (*Countie, error) {
	c := &Countie{}
	for _, o := range opts {
		e := o(c)
		if e != nil {
			return nil, e
		}
	}
	return c, nil
}

func WithInput(input io.Reader) option {
	return func(c *Countie) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
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
