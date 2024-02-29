package main

import "fmt"

func Unbuffer(arr []int) <-chan int {

	a := make(chan int)

	go func() {
		for _, v := range arr {
			a <- v

		}
		close(a)

	}()
	return a
}

func main() {

	var c []int = make([]int, 3)

	c[0] = 10
	c[1] = 20
	c[2] = 30

	b := Unbuffer(c)

	k := make(chan int)

	go func() {
		for n := range b {

			k <- n

		}
		close(k)

	}()

	for i := range k {

		fmt.Println(i)

	}

}

/*

Process 

 
1. Basically unbuffered channel can contain maximum one value 

2. Buffer channel will block when buffer is full.

3. when writing un buffer channel we should be use go roututine for synchronization 





*/

