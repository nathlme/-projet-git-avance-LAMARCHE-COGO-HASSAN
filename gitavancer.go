package main

import "fmt"

func FiltrerPairs(nombres []int) []int {
	var pairs []int
	for _, nombre := range nombres {
		if nombre%2 == 0 {
			pairs = append(pairs, nombre)
		}
	}
	return pairs
}

func main() {
	liste := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pairs := FiltrerPairs(liste)
	fmt.Println("Nombres pairs :", pairs)
}

func getTens() []int {
	var tens []int
	for i := 10; i <= 100; i += 10 {
		tens = append(tens, i)
	}
}

func main() {
	tens := getTens()
	fmt.Println("Les dizaines de 1 à 100 :", tens)
}
