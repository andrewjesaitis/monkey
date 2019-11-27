package evaluator

import (
	"testing"

	"github.com/andrewjesaitis/monkey/lexer"
	"github.com/andrewjesaitis/monkey/object"
	"github.com/andrewjesaitis/monkey/parser"
)

func TestEvalIntergerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"123", 123},
		{"-5", -5},
		{"-123", -123},
		{"1+2", 3},
		{"5-2", 3},
		{"2+2+4-3", 5},
		{"5*5", 25},
		{"5*2-3", 7},
		{"6/3+7", 9},
		{"9-2*3", 3},
		{"7+6/2", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"false", false},
		{"true", true},
		{"1<2", true},
		{"2<1", false},
		{"1<1", false},
		{"1>1", false},
		{"2>1", true},
		{"1>2", false},
		{"1==1", true},
		{"1!=1", false},
		{"1==2", false},
		{"2!=1", true},
		{"(1<2) == true", true},
		{"(2<1) == true", false},
		{"(1>2) == false", true},
		{"(2>1) == false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!false", true},
		{"!true", false},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 5 }", 5},
		{"if (1 < 2) { 10 } else { 5 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integrer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integrer))
		} else {
			testNullObject(t, evaluated)
		}
	}

}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%v)", obj, obj)
		return false
	}
	return true
}
