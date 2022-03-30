package main

import (
	"flag"
	"os"
	"fmt"
)

func runhelp() {
    fmt.Fprintf(os.Stderr, "Usage of %s: \n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}