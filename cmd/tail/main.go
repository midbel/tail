package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/midbel/tail"
)

func main() {
	var (
		index  = flag.Bool("x", false, "print line number")
		lines  = flag.Int("n", 10, "display last N lines")
		header bool
	)
	flag.Parse()

	header = flag.NArg() > 1
	for _, a := range flag.Args() {
		tailFile(a, *lines, header, *index)
	}
}

func tailFile(file string, lines int, header, index bool) {
	r, err := tail.Tail(file, lines)
	if err != nil {
		return
	}
	defer r.Close()

	if header {
		fmt.Fprintf(os.Stdout, ">>> %s <<<\n", file)
	}
	var (
		line = 1
		scan = bufio.NewScanner(r)
	)
	for scan.Scan() {
		if index {
			fmt.Fprintf(os.Stdout, "%d: ", line)
		}
		fmt.Fprintln(os.Stdout, scan.Text())
		line++
	}
	if line > 0 && header {
		fmt.Fprintln(os.Stdout)
	}
}
