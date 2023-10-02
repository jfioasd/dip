package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"math/big"
)

func run_dip(code string, stack []*big.Int) []*big.Int {
	ip := 0
	for ip < len(code) {
		switch code[ip] {
		case '0':
			stack = append(stack, big.NewInt(0))

		case '\'':
			stack[len(stack) - 1] .Add(
				stack[len(stack) - 1],
				big.NewInt(1))

		case ';':
			y := stack[len(stack) - 1]
			copy(stack[1:], stack)
			stack[0] = y
		
        case '.':
			fmt.Print(stack[len(stack) - 1], " ")
		
		case '!':
			fmt.Println(stack)
		
		case '(':
			lvl := 1
			l_bound := ip + 1
			for ip < len(code) && lvl > 0 {
				ip ++
				if code[ip] == '(' {
					lvl ++
				}
				if code[ip] == ')' {
					lvl --
				}
			}

			for stack[len(stack) - 1].Cmp(big.NewInt(0)) != 0 {
				stack[len(stack) - 1] .Sub (
					stack[len(stack) - 1],
				    big.NewInt(1))
				stack = run_dip(code[l_bound:ip], stack)
			}

			stack = stack[:len(stack) - 1]
		}
		ip ++
	}
	return stack
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <file>\n", os.Args[0])
		return
	}
	prog, _ := ioutil.ReadFile(os.Args[1])
	run_dip(string(prog), []*big.Int{})
}
