package main

import (
	"fmt"
	"log"

	"github.com/haunt98/evaluator/pkg/evaluate"
	"github.com/haunt98/evaluator/pkg/parser"
)

func main() {
	input := `$x == 1 and $y == "a"`

	// parse section
	p := parser.NewParser(input)
	expr, err := p.Parse()
	if err != nil {
		log.Fatal(err)
	}

	// evaluate section
	args := map[string]interface{}{
		"x": int64(1),
		"y": "a",
	}

	result, err := evaluate.Evaluate(expr, args)
	if err != nil {
		log.Fatal(expr)
	}

	fmt.Println(result)
}
