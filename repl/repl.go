package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/gavwyh/go-interpreter/token"
	"github.com/gavwyh/go-interpreter/lexer"
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
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.nextToken() {
			fmt.Fprintf(out, "+v\n", tok)
		}
	}
}