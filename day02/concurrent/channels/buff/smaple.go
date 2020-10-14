package buff

import (
	"fmt"
	"math/rand"
	"time"
)

func calculateValue(values chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	time.Sleep(1000 * time.Millisecond)
	values <- value
	fmt.Println("Only Executes after another goroutine performs a receive on the channel")
}

// Sample invokes go routines
func Sample() {
	values := make(chan int, 3)
	defer close(values)

	go calculateValue(values)
	go calculateValue(values)
	go calculateValue(values)

	value := <-values
	fmt.Println(value)

	time.Sleep(1000 * time.Millisecond)
}
