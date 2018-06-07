package main

import (
	"fmt"
)

var name string

func main(){
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("The name entered is: " +name)
}




