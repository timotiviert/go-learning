package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	//"os"
	//"strconv"
)

const (
	methodAddition       = "+"
	methodSubtraction    = "-"
	methodMultiplication = "x"
	methodDivision       = "/"
	methodModulus        = "%"
)

func exit(s string) {
	fmt.Println(s)
	os.Exit(1)
}

func main() {
	//var file string
	//flag.StringVar(&file, "file", "-", "output file name")
	flag.Parse()

	//fmt.Println("file:", file)

	//for _, arg := range flag.Args() {
	//	fmt.Printf("%s\n", arg)
	//}

	if flag.NArg() != 3 {
		exit("usage: calculator -file <file> <a> <operand> <b>")
	}

	// Parsing of a, method and b.
	a, err := strconv.ParseFloat(flag.Arg(0), 64)
	if err != nil {
		exit("invalid argument: a must be a number")
	}
	method := flag.Arg(1)
	b, err := strconv.ParseFloat(flag.Arg(2), 64)
	if err != nil {
		exit("invalid argument: b must be a number")
	}

	// The result.
	var c float64

	switch method {
	case methodAddition:
		c = a + b
	case methodSubtraction:
		c = a - b
	case methodMultiplication:
		c = a * b
	case methodDivision:
		if b == 0.0 {
			exit("division by zero")
		}
		c = a / b
	case methodModulus:
		c = a * b
	default:
		exit("invalid argument: method must be +, -, x, / or %")
	}

	fmt.Println(c)

}
