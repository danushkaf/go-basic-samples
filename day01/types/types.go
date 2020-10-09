package types

import "fmt"

// Types method prints default values of premetive types
func Types() {
	var i int
	fmt.Println("default int is: ", i)
	var s string
	fmt.Println("default string is: ", s)
	var f float64
	fmt.Println("default float64 is: ", f)
	var arInt [3]int
	fmt.Println("default int array is: ", arInt)
	var c complex64
	fmt.Println("default complex64 is: ", c)
}
