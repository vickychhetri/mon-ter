package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/vickychhetri/mon-ter/lexer"
	"github.com/vickychhetri/mon-ter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		fmt.Println(line)

		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
