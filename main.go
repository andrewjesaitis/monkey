package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/andrewjesaitis/monkey/evaluator"
	"github.com/andrewjesaitis/monkey/lexer"
	"github.com/andrewjesaitis/monkey/object"
	"github.com/andrewjesaitis/monkey/parser"
	"github.com/andrewjesaitis/monkey/repl"
)

func main() {
	if len(os.Args) > 2 {
		panic("Please pass a single input file to evaluate or pass no argument to access the repl.")
	} else if len(os.Args) == 2 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("Can't read file: ", os.Args[1])
			panic(err)
		}
		l := lexer.New(string(data))
		p := parser.New(l)
		program := p.ParseProgram()

		errors := p.Errors()
		if len(errors) != 0 {
			for _, msg := range errors {
				fmt.Printf(msg + "\n")
			}
			panic("Could not parse input file.")
		}
		evaluator.Eval(program, object.NewEnvironment())
	} else {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Monkey programming language!\n",
			user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
