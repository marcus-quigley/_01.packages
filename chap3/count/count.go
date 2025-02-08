package count

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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
	// for _,f:= range c.files{
	// 	f.(.io.closer)
	// }
	return i
}

func (c Counter) Words() int {
	i := 0
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		l := scanner.Text()
		i += strings.Count(l, " ") + 1
	}
	return i
}

func (c Counter) Bytes() int {
	i := 0
	scanner := bufio.NewScanner(c.input)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		i += len(scanner.Bytes())
	}
	return i
}

func Main() int {
	lineMode := flag.Bool("lines", false, "Count lines, not words")
	byteMode := flag.Bool("bytes", false, "Count bytes, not words")
	flag.Usage = func() {
		fmt.Printf("Usage: %s [-lines | -bytes] [files...]\n", os.Args[0])
		fmt.Println("Counts words (or lines or bytes) in named files or standard input.\nFlags:")
		flag.PrintDefaults()
	}
	flag.Parse()
	c, err := NewCounter(
		WithInputFromArgs(flag.Args()),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	switch {
	case *lineMode && *byteMode:
		fmt.Fprintln(os.Stderr, "Please specify either '-lines' or '-bytes', but not both.")

	case *lineMode:
		fmt.Println(c.Lines())
	case *byteMode:
		fmt.Println("sh 1")

	default:
		fmt.Println(c.Words())
	}
	return 0
}
