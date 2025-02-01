package greeting

import (
	"fmt"
	"io"
)

type Greetie struct {
	Name string
}

func (g Greetie) Greet(w io.Writer) {
	fmt.Fprintf(w, "hello, %s", g.Name)
}

func (g *Greetie) ReadName(w io.Writer, r io.Reader) {
	fmt.Fprintf(w, "name, fucko?")
	fmt.Fscanln(r, &g.Name)
}
