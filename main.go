package main

import "fmt"

func AddTwo(n int) int {
	return n + 2
}

func Square(n int) int {
	return n * n
}

func Negative(n int) int {
	return -n
}

func main() {
	fmt.Println("le nombre final est : ", AddTwo(10))
	fmt.Println("le nombre final est : ", Square(10))
	fmt.Println("le nombre final est : ", Negative(10))
}
