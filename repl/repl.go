package repl

import (
  "fmt"
  "os"
  "bufio"
  "baboon/lexer"
  "baboon/token"
)

const PROMPT string = ">> "

func Start() {
  fmt.Println("Welcome to the baboon interpreter!")
  fmt.Printf(PROMPT)

  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    l := lexer.New(scanner.Text())

    for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
      fmt.Printf("%+v\n", t)
    }

    fmt.Printf(PROMPT)
  }
  fmt.Println()
}
