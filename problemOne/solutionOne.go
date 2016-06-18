package main

import "fmt"

func main() {

	belowTen := [...]int{3, 5, 6, 9}
	sum := 0
	for _, belowTen := range belowTen {
		sum += belowTen
	}
	fmt.Println("sum: ", sum)

}
