package main

import "fmt"

func AddTwo(n int) int {
	return n + 2
}

func main() {
	fmt.Println("le nombre final est : ", AddTwo(10))
}
