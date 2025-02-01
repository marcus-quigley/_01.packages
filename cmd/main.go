package main

import (
	"fmt"
)

func main() {
	// buf := &bytes.Buffer{}
	name := ""
	fmt.Print("whats your name, bitch?")
	_, e := fmt.Scanln(&name) //.Fscan(os.Stdin, buf)
	if e != nil {
		fmt.Printf("error entering aname, %v\n", e.Error())
		return
	}
	// name := buf.String()
	fmt.Printf("hello, %s", name)
}
