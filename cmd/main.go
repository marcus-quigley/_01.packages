package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/mq/packages/greeting"
	"github.com/mq/packages/hello"
	"github.com/mq/packages/pipeline"
	"github.com/mq/packages/shell"
	"github.com/mq/packages/writer"
)

func main() {

	pipeline.FromString("hello, world\n").Stdout()
}

func main7() {
	// read in file, get ip from 1st column of each line
	//total them
	//sort in desc
	//show first 10

	f, e := os.Open("log.txt")
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		return
		//os.Exit(1)
	}
	defer f.Close()
	input := bufio.NewScanner(f)

	uniques := map[string]int{}
	var args []string
	for input.Scan() {
		l := input.Text()
		args = strings.Fields(l)
		if len(args) > 0 {
			uniques[args[0]]++
		}
	}

	type freq struct {
		addr  string
		count int
	}
	freqs := make([]freq, 0, len(uniques))
	for a, c := range uniques {
		freqs = append(freqs, freq{a, c})
	}

	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].count > freqs[j].count
	})
	fmt.Printf("%-16s%s\n", "address", "requests")
	for i, f := range freqs {
		if i > 9 {
			break
		}
		fmt.Printf("%-16s%d\n", f.addr, f.count)

	}
}

func main6() {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		cmd, e := shell.CreateCommand(line)
		if e != nil {
			fmt.Printf("error creating command %v\n", e)
			continue
		}
		out, e := cmd.CombinedOutput()
		if e != nil {
			fmt.Printf("error after creating command %v\n", e)
		}
		fmt.Printf("%s\n", out)
		fmt.Print("> ")
	}
	fmt.Println("see ya...")
}

func main5() {
	os.Exit(writer.Main())
	//os.Exit(count.Main())
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
// 	// 	lines++s
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
