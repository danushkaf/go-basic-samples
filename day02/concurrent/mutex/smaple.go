package mutex

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex   sync.Mutex
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}

// Sample invokes go routines
func Sample() {
	var wg sync.WaitGroup
	wg.Add(5)
	go withdraw(700, &wg)
	go deposit(500, &wg)
	go deposit(500, &wg)
	go deposit(500, &wg)
	go withdraw(700, &wg)
	wg.Wait()

	fmt.Printf("New Balance %d\n", balance)

	time.Sleep(1000 * time.Millisecond)
}

func init() {
	balance = 1000
	fmt.Println("Init function called in mutex")
}
