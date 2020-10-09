package pointers

import "fmt"

// Number is a struct
type Number struct {
	Value int
}

// Print is printing info
func Print() {
	i := 5
	fmt.Println("i is: ", i)
	fmt.Println("address of i is: ", &i)
}

// ChangeValue changes the i of Number Value
func ChangeValue(num Number) {
	num.Value = 10
}

// ChangePointerValue changes the i of Number Pointer
func ChangePointerValue(num *Number) {
	num.Value = 10
}

// ChangeValue changes the i of Number Value
func (num *Number) ChangeValue() {
	num.Value = 10
}
