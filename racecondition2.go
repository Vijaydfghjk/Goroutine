// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"sync"
)

var (
	//mutex   sync.Mutex
	balance int
)

func deposite(value int, wg *sync.WaitGroup) {

	//mutex.Lock()
	fmt.Printf("Depositing %d to account with balance %d\n", value, balance)
	balance += value
	//mutex.Unlock()
	wg.Done()
}

func Withdraw(value int, wg *sync.WaitGroup) {
	//mutex.Lock()
	fmt.Printf("Withdrawing %d to account with balance %d\n", value, balance)
	balance -= value
	//mutex.Unlock()
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	balance = 1000

	wg.Add(2)

	go Withdraw(700, &wg)
	go deposite(500, &wg)

	wg.Wait()

	fmt.Println(balance)
}
