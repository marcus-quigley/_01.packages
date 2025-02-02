package hello

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	Output io.Writer
}

func Main() {
	NewPrinter().Print()
}

func NewPrinter() *Printer {
	return &Printer{
		Output: os.Stdout,
	}
}

func (p *Printer) Print() {
	fmt.Fprintf(p.Output, "hello, world\n")
}

// func PrintTo(w io.Writer) {
// 	fmt.Fprint(w, "hello, world\n")
// }
