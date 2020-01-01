package evaluator

import (
	"testing"

	"github.com/andrewjesaitis/monkey/object"
)

func TestQuote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`quote(5)`, `5`},
		{`quote(5+3)`, `(5 + 3)`},
		{`quote(foobar)`, `foobar`},
		{`quote(foo + bar)`, `(foo + bar)`},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		quote, ok := evaluated.(*object.Quote)

		if !ok {
			t.Fatalf("expected *object.Quote. got=%T (%+v)",
				evaluated, evaluated)
		}

		if quote == nil {
			t.Fatalf("quote.Node is nil")
		}

		if quote.Node.String() != tt.expected {
			t.Errorf("not equal. got=%q, want=%q", quote.Node.String(), tt.expected)
		}
	}
}

func TestUnquote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`quote(unquote(5))`, `5`},
		{`quote(unquote(5+3))`, `8`},
		{`quote(5 + unquote(3+3))`, `(5 + 6)`},
		{`quote(unquote(4-2) + 5)`, `(2 + 5)`},
		{`let foobar = 8; quote(foobar)`, `foobar`},
		{`let foobar = 8; quote(unquote(foobar))`, `8`},
		{`quote(unquote(true))`, `true`},
		{`quote(unquote(true == false))`, `false`},
		{`quote(unquote("abc"))`, `abc`},
		{`quote(unquote("a" + "bc"))`, `abc`},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		quote, ok := evaluated.(*object.Quote)

		if !ok {
			t.Fatalf("expected *object.Quote. got=%T (%+v)",
				evaluated, evaluated)
		}

		if quote == nil {
			t.Fatalf("quote.Node is nil")
		}

		if quote.Node.String() != tt.expected {
			t.Errorf("not equal. got=%q, want=%q", quote.Node.String(), tt.expected)
		}
	}
}
