package main

import (
	"fmt"
	// "github.com/danushkaf/go-basic-samples/day01/deferst"
	// "github.com/danushkaf/go-basic-samples/day01/hello"
	// "github.com/danushkaf/go-basic-samples/day01/oop"
	// "github.com/danushkaf/go-basic-samples/day01/pointers"
	// "github.com/danushkaf/go-basic-samples/day01/types"
)

func main() {
	fmt.Println("Hello world")

	// // fmt
	// var st2 string
	// st2 = "Hello"
	// st := "Danushka"
	// fmt.Printf("%s world %s !\nHave a nice day.\n", st2, st)
	//
	// log := fmt.Sprintf("Hello world %s !\nHave a nice day.", st)
	// fmt.Println(log)
	//
	// // Libs
	// fmt.Println(hello.Hello())
	//
	// // Types
	// types.Types()
	//
	// // Pointers Intro
	// pointers.Print()
	//
	// // Pointer Deep Dive
	// num := pointers.Number{
	// 	Value: 5,
	// }
	// fmt.Printf("Value of i is : %d \n", num.Value)
	// pointers.ChangeValue(num)
	// fmt.Printf("Value of i after passing value is : %d \n", num.Value)
	// pointers.ChangePointerValue(&num)
	// fmt.Printf("Value of i after passing pointer is : %d \n", num.Value)
	//
	// // Use Object Method
	// num = pointers.Number{
	// 	Value: 5,
	// }
	// fmt.Printf("Value of i is : %d \n", num.Value)
	// num.ChangeValue()
	// fmt.Printf("Value of i after passing value is : %d \n", num.Value)
	//
	// // Defer
	// deferst.DoDBOperations()
	//
	// // OOP
	//
	// // Interface and Inheritance and Polymorphism
	// r := oop.Rectangle{Length: 5, Width: 3}
	// fmt.Println("Rectangle details are: ", r)
	// fmt.Println("Area of Rectangle: ", r.Area())
	// q := oop.Square{Side: 5}
	// fmt.Println("Square details are: ", q)
	//
	// var s oop.ShaperArea
	// var sp oop.ShaperPerimeter
	//
	// s = r
	// sp = oop.ShaperPerimeter(r)
	// fmt.Println("Area of ShaperArea(Rectangle): ", s.Area())
	// fmt.Println("Perimeter of ShaperPerimeter(Rectangle): ", sp.Perimeter())
	// fmt.Println()
	//
	// s = q
	// sp = oop.ShaperPerimeter(q)
	// fmt.Println("Area of ShaperArea(Square): ", s.Area())
	// fmt.Println("Perimeter of ShaperPerimeter(Square): ", sp.Perimeter())
	// fmt.Println()
	//
	// shapeAreaArr := [...]oop.ShaperArea{r, q}
	// shapePerimeterArr := [...]oop.ShaperPerimeter{r, q}
	//
	// fmt.Println("Looping through shapes for area ...")
	// for n := range shapeAreaArr {
	// 	fmt.Println("Shape details: ", shapeAreaArr[n])
	// 	fmt.Println("Area of this shape is: ", shapeAreaArr[n].Area())
	// }
	//
	// fmt.Println("Looping through shapes for Perimeter ...")
	// for n := range shapePerimeterArr {
	// 	fmt.Println("Shape details: ", shapePerimeterArr[n])
	// 	fmt.Println("Perimeter of shape: ", shapePerimeterArr[n].Perimeter())
	// }
	//
	// // Multiple Inheritance
	// cp := new(oop.CameraPhone) //a new camera phone instance
	// fmt.Println("Our new CameraPhone exhibits multiple behaviors ...")
	// // fmt.Println("First switch it on : ", cp.SwitchOn())
	// fmt.Println("It can take a picture: ", cp.TakePicture())
	// fmt.Println("It can also make calls: ", cp.Call())
}
