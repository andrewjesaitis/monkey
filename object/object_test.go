package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello"}
	hello2 := &String{Value: "Hello"}
	bye1 := &String{Value: "Goodbye"}
	bye2 := &String{Value: "Goodbye"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if bye1.HashKey() != bye2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if hello1.HashKey() == bye1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	one := &Integer{Value: 1}
	another_one := &Integer{Value: 1}
	two := &Integer{Value: 2}
	another_two := &Integer{Value: 2}

	if one.HashKey() != another_one.HashKey() {
		t.Errorf("integers with same value have different hash keys")
	}
	if two.HashKey() != another_two.HashKey() {
		t.Errorf("integers with same value have different hash keys")
	}
	if one.HashKey() == two.HashKey() {
		t.Errorf("integers with different values have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	ya := &Boolean{Value: true}
	ya2 := &Boolean{Value: true}
	nah := &Boolean{Value: false}
	nah2 := &Boolean{Value: false}

	if ya.HashKey() != ya2.HashKey() {
		t.Errorf("booleans with same value have different hash keys")
	}
	if nah.HashKey() != nah2.HashKey() {
		t.Errorf("booleans with same value have different hash keys")
	}
	if ya.HashKey() == nah.HashKey() {
		t.Errorf("booleans with different values have same hash keys")
	}
}
