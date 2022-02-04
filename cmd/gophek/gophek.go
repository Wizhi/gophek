package main

import (
	"flag"
	"os"
	"strings"

	"github.com/wizhi/gophek"
)

func main() {
	flag.Parse()

	gophek.Eval(strings.Join(flag.Args(), ""), os.Stdin, os.Stdout)
}
