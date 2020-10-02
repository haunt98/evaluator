package evaluate

import (
	"testing"

	"github.com/haunt98/evaluator/pkg/parser"
)

//goos: linux
//goarch: amd64
//pkg: evaluator/pkg/evaluate
//BenchmarkEvaluate
//BenchmarkEvaluate-4   	 5635582	       244 ns/op
func BenchmarkEvaluate(b *testing.B) {
	input := `!($x == true or $y != 1) and $z == "a" or $t in [true, 1, "a"]`
	p := parser.NewParser(input)

	args := map[string]interface{}{
		"x": true,
		"y": int64(1),
		"z": "a",
		"t": true,
	}

	expr, _ := p.Parse()
	var result bool
	var err error
	for n := 0; n < b.N; n++ {
		result, err = Evaluate(expr, args)
	}

	_ = err
	_ = result
}
