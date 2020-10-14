package unbuff

import (
	"fmt"
	"math/rand"
	"time"
)

func calculateValue(values chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	values <- value
	fmt.Println("Only Executes after another goroutine performs a receive on the channel")
}

// Sample invokes go routines
func Sample() {
	values := make(chan int)
	defer close(values)

	go calculateValue(values)
	go calculateValue(values)

	value := <-values
	fmt.Println(value)

	// value = <-values
	// fmt.Println(value)

	time.Sleep(1000 * time.Millisecond)
}
