package simple

import (
	"fmt"
	"math/rand"
	"time"
)

func calculateValue(values chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	values <- value
}

// Sample invokes go routines
func Sample() {
	values := make(chan int)
	defer close(values)

	go calculateValue(values)

	// for index := 0; index < 100; index++ {
	// 	fmt.Println("Go Channel Printline {}", index)
	// }

	value := <-values

	// for index := 0; index < 10; index++ {
	// 	fmt.Println("Go Channel Printline After {}", index)
	// }

	fmt.Println(value)

	time.Sleep(1000 * time.Millisecond)
}
