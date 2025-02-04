package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type option func(*Counter) error

type Counter struct {
	input io.Reader
}

func NewCounter(opts ...option) (*Counter, error) {
	c := &Counter{
		input: os.Stdin,
	}
	for _, o := range opts {
		e := o(c)
		if e != nil {
			return nil, e
		}
	}
	return c, nil
}

func WithInput(input io.Reader) option {
	return func(c *Counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(c *Counter) error {
		if args == nil || len(args) < 1 {
			return nil
		}
		f, e := os.Open(args[0])
		if e != nil {
			return e
		}
		//input := strings.NewReader(strings.Join(args, ""))
		c.input = f
		return nil
	}
}

func (c Counter) Lines() int {
	i := 0
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		i += 1
	}
	return i
}

func Main() int {
	c, err := NewCounter(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Println(c.Lines())
	return 0
}
