package greeting

import (
	"fmt"
	"io"
	"os"
)

type Greetie struct {
	Name   string
	Output io.Writer
	Input  io.Reader
}

func Run() {
	g := New()
	g.ReadName()
	g.Greet()
}

func New() Greetie {
	return Greetie{
		Output: os.Stdout,
		Input:  os.Stdin,
	}
}

func (g Greetie) Greet() {
	fmt.Fprintf(g.Output, "hello, %s", g.Name)
}

func (g *Greetie) ReadName() {
	fmt.Fprintf(g.Output, "name, fucko?")
	fmt.Fscanln(g.Input, &g.Name)
}
