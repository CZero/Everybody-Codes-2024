// Everybody codes 2024: Quest 1: The Battle for the Farmlands - Part 1
package main

import (
	"fmt"

	"github.com/CZero/libaoc"
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("notes.txt")
	if err != nil {
		panic(err)
	}
	values := getValues(input)
	fmt.Printf("Potions needed: %d\n", libaoc.SumSlice(values))
}

func getValues(input []string) (values []int) {
	for _, letter := range input[0] {
		switch string(letter) {
		case "A":
			values = append(values, 0)
		case "B":
			values = append(values, 1)
		case "C":
			values = append(values, 3)
		default:
			panic("Unknown value hit!")
		}
	}
	return values
}
