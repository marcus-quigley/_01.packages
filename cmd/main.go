package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/mq/packages/count"
	"github.com/mq/packages/greeting"
	"github.com/mq/packages/hello"
)

func main() {

	// lines := 0
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	lines++
	// }
	// fmt.Println(lines)
	c := count.NewCountie(
		count.WithInput(bytes.NewBufferString("this\nis\nit\n")),
	)
	fmt.Fprintf(os.Stdout, "lines: %d", c.Lines())
}

func main2() {
	greeting.GreetUser(os.Stdout, os.Stdin)

	//hello.PrintTo(os.Stdout)
	hello.Main()
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
