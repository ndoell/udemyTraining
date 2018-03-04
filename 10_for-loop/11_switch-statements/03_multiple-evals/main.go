package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Daniel", "Jenny":
		fmt.Println("Wassup Daniel, or, err, Jenny")
	case "Medhi", "Marcus":
		fmt.Println("Wassup Medhi, or, err, Marcus")
	case "Jullian", "Sushant":
		fmt.Println("Wassup Jullian, or, err, Sushant")
	default:
		fmt.Println("Have you no friends")
	}
}

