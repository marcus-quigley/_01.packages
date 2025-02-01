package main

import (
	"fmt"
	"os"

	greeting "github.com/mq/packages"
)

func main() {
	// buf := &bytes.Buffer{}
	// name := ""
	// fmt.Print("whats your name, bitch?")
	// _, e := fmt.Scanln(&name) //.Fscan(os.Stdin, buf)

	// name, e := greeting.ReadName()
	// if e != nil {
	// 	fmt.Printf("error entering aname, %v\n", e.Error())
	// 	return
	// }

	//greeting.Greet(os.Stdout, name)
	// fmt.Printf("hello, %s", name)
	g := greeting.Greetie{}
	g.ReadName(os.Stdout, os.Stdin)
	g.Greet(os.Stdout)
}

func main1() {
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
