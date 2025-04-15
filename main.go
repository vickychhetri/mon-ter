package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/vickychhetri/mon-ter/repl"
)

func main() {
	// Example usage of the lexer and token package
	// input := `let five=5;
	// let ten = 10;
	// let add = fn(x,y){
	// x + y;
	// };let result = add(five,ten);
	// !-/*5;
	// 5 < 10 > 5;
	// if 5 < 10 {
	// 	return true;
	// } else {
	// 	return false;
	// }
	// 	10 == 10;
	// 	10 != 9;

	// `
	// l := lexer.New(input)

	// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
	// 	fmt.Printf("[Type: %s ], [Literal: %s]\n", tok.Type, tok.Literal)
	// }

	user, err := user.Current()

	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is mon-ter programming language!\n", user.Username)

	fmt.Printf("Feel free to type in commands \n")

	repl.Start(os.Stdin, os.Stdout)
}
