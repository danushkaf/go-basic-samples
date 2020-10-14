package mutex

import (
	"fmt"
	"sync"
)

// Deadlock is to simutalte deadlock
func Deadlock() {
	var b sync.Mutex
	fmt.Println("Locking for first time")
	b.Lock()
	fmt.Println("Trying to lock again")
	b.Lock()
	fmt.Println("This never executes as we are in deadlock")
}
