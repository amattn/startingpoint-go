package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

// https://groups.google.com/forum/#!msg/golang-nuts/a9PitPAHSSU/ziQw1-QHw3EJ
const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

var show_h bool
var show_help bool
var show_version bool

func init() {
	flag.BoolVar(&show_h, "h", false, "show help message and exit(0)")
	flag.BoolVar(&show_help, "help", false, "show help message and exit(0)")
	flag.BoolVar(&show_version, "version", false, "show version info and exit(0)")
}

func main() {
	log.Printf("> startingpoint-go\n")
	log.Printf(">> gostart version %v (build %d)\n", Version(), BuildNumber())
	log.Printf(">> Go Version: %v %v/%v\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	// command line flags:
	flag.Parse()

	if show_version {
		os.Exit(0)
	}

	if show_h || show_help {
		flag.Usage()
		os.Exit(0)
	}

	// do stuff

	fmt.Println("hello world")
}
