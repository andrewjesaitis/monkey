package evaluator

import (
	"testing"

	"github.com/andrewjesaitis/monkey/lexer"
	"github.com/andrewjesaitis/monkey/object"
	"github.com/andrewjesaitis/monkey/parser"
)

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{"5 + true;", "type mismatch: INTEGER + BOOLEAN"},
		{"5 + true; 5;", "type mismatch: INTEGER + BOOLEAN"},
		{"-true", "unknown operator: -BOOLEAN"},
		{"true + false", "unknown operator: BOOLEAN + BOOLEAN"},
		{"5; true - true; 5;", "unknown operator: BOOLEAN - BOOLEAN"},
		{"if (10 > 1) { true * true; }", "unknown operator: BOOLEAN * BOOLEAN"},
		{`
         if (10 > 3) {
	         if (5 < 10) {
                 return true / false;
             }
         return 1;
         }
         `,
			"unknown operator: BOOLEAN / BOOLEAN"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)

		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}

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

func TestReturnStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9*9", 10},
		{"return 2 * 5 * 3; 9", 30},
		{"9; return 8; 2*2;", 8},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
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
