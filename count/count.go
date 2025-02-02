package count

import (
	"bufio"
	"io"
)

type Countie struct {
	Input io.Reader
}

func NewCountie() *Countie {
	return &Countie{}
}

func (c Countie) Lines() int {
	i := 0
	scanner := bufio.NewScanner(c.Input)
	for scanner.Scan() {
		i += 1
	}
	return i
}
