package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/mq/packages/chap3/count"
	"github.com/mq/packages/greeting"
	"github.com/mq/packages/hello"
)

func main() {
	count.Main()
}

func main4() {
	input := bytes.NewBufferString("this\nis\nit\n")
	match := "is"
	lines := []string{}
	// m:= NewMatch
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		l := scanner.Text()
		if strings.Contains(l, match) {
			lines = append(lines, l)
		}
	}
	fmt.Println(lines)
}

// func main3() {

// 	// lines := 0
// 	// input := bufio.NewScanner(os.Stdin)
// 	// for input.Scan() {
// 	// 	lines++
// 	// }
// 	// fmt.Println(lines)
// 	c, _ := count.NewCountie(
// 		count.WithInput(bytes.NewBufferString("this\nis\nit\n")),
// 	)
// 	fmt.Fprintf(os.Stdout, "lines: %d", c.Lines())
// }

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
