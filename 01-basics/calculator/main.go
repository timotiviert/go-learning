package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(0)
	file := flag.String("file", "", "input file")
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

	op := args[1]

	b, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		log.Fatal("invalid argument: b must be a number")
	}

	result, err := calculate(a, b, op)

	if err != nil {
		log.Fatal(err)
	}

	var out io.Writer = os.Stdout
	if *file != "" {
		f, err := os.Create(*file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		out = f
	}

	fmt.Fprintln(out, result)
}

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "x", "*":
		return a * b, nil
	case "/":
		if b == 0.0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	case "%":
		return math.Mod(a, b), nil
	default:
		return 0, errors.New("invalid argument: operation must be one of +, -, x, / or %")
	}
}
