package main

import (
	"flag"
	"fmt"
	"log"
	//"os"
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

func main() {
	log.SetFlags(0)
	flag.Parse()

	args := flag.Args()

	if len(args) != 3 {
		log.Fatal("usage: calculator -file <file> <a> <operand> <b>")
	}

	// Parsing of a, method and b.
	a, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		log.Fatal("invalid argument: a must be a number")
	}

	method := args[1]

	b, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		log.Fatal("invalid argument: b must be a number")
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
			log.Fatal("division by zero")
		}
		c = a / b
	case methodModulus:
		c = a * b
	default:
		log.Fatal("invalid argument: method must be +, -, x, / or %")
	}

	fmt.Println(c)
}
