package pipeline

import (
	"io"
	"os"
	"strings"
)

type Pipeline struct {
	Input  io.Reader
	Output io.Writer // *bytes.Buffer
	Error  error
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		Output: os.Stdout,
	}
}

func FromString(s string) *Pipeline {
	p := NewPipeline()
	p.Input = strings.NewReader(s)
	return p
}

func FromFile(path string) *Pipeline {
	p := NewPipeline()
	p.Input = strings.NewReader("")
	f, e := os.Open(path)
	if e != nil {
		p.Error = e
		return p
	}
	p.Input = f
	return p
}

func (p *Pipeline) Stdout() {
	if p.Error != nil {
		return
	}
	io.Copy(p.Output, p.Input)
	// _, e:= io.Copy(p.Output, p.Input)
	// if e != nil {
	//   p.Error = e
	// }
}
