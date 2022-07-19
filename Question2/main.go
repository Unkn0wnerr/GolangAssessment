package main

import "fmt"

func iterate(c chan int, slc []int) {
	for _, val := range slc {
		c <- val
	}
	close(c)
}

func main() {
	c := make(chan int)
	slc := []int{10, 20, 35, 100, 200, 502}
	go iterate(c, slc)

	for val := range c {
		fmt.Println(val)
	}

}
