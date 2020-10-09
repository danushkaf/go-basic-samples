package pointers

import (
	"testing"
)

// resultNum is the wanted result
func resultNum() Number {
	return Number{
		Value: 10,
	}
}

// inputNum is the input object
func inputNum() Number {
	return Number{
		Value: 5,
	}
}

func TestPointers(t *testing.T) {
	want := resultNum()
	input := inputNum()
	ChangePointerValue(&input)
	if got := input; got.Value != want.Value {
		t.Errorf("ChangePointerValue() = %d, want = %d", got.Value, want.Value)
	}
}

func TestValue(t *testing.T) {
	want := inputNum()
	input := inputNum()
	ChangeValue(input)
	if got := input; got.Value != want.Value {
		t.Errorf("ChangeValue() = %d, want = %d", got.Value, want.Value)
	}
}

func TestObjectValue(t *testing.T) {
	want := resultNum()
	input := inputNum()
	input.ChangeValue()
	if got := input; got.Value != want.Value {
		t.Errorf("ChangeValue() = %d, want = %d", got.Value, want.Value)
	}
}
