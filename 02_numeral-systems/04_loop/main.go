package main

import "fmt"

//Write a loop to print different formats, tab separated.

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)

	}
}
