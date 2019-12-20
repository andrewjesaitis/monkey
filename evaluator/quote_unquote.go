package evaluator

import (
	"github.com/andrewjesaitis/monkey/ast"
	"github.com/andrewjesaitis/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
