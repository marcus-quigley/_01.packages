package greeting

import (
	"bufio"
	"fmt"
	"io"
)

// type Greetie struct {
// 	Name   string
// 	Output io.Writer
// 	Input  io.Reader
// }

// func Run() {
// 	g := New()
// 	g.ReadName()
// 	g.Greet()
// }

// func New() Greetie {
// 	return Greetie{
// 		Output: os.Stdout,
// 		Input:  os.Stdin,
// 	}
// }

func GreetUser(w io.Writer, r io.Reader) {
	name := ""
	fmt.Fprintln(w, "name, fucko?")
	input := bufio.NewScanner(r)
	if input.Scan() {
		name = input.Text()
	}
	fmt.Fprintf(w, "hello, %s\n", name)

}

// func (g Greetie) Greet() {
// 	fmt.Fprintf(g.Output, "hello, %s", g.Name)
// }

// func (g *Greetie) ReadName() {
// 	fmt.Fprintf(g.Output, "name, fucko?")
// 	fmt.Fscanln(g.Input, &g.Name)
// }
