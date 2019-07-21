package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "./sum/sum"
import "./mult/mult"

func main() {
	fmt.Println("Calculator")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type your expression: ")
	text, _ := reader.ReadString('\n')
	s := strings.Split(text, " ")
	sa, op, sb := s[0], s[1], s[2]
	
	a, _ := strconv.Atoi(sa)
	b, _ := strconv.Atoi(sb)

	switch op {
		case "+": fmt.Printf("Sum %d + %d is %d", a, b, sum.Sum(a, b))
		case "*": fmt.Printf("Mult %d * %d is %d", a, b, mult.Mult(a, b))
		default: fmt.Println("Unsupported operation!")
	}
	fmt.Printf("\n")
}