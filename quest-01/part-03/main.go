// Everybody codes 2024: Quest 1: The Battle for the Farmlands
// Part 3
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
	empty := 0
	count := 1
	for _, letter := range input[0] {
		if count == 1 {
			empty = 0
		}
		switch string(letter) {
		case "A":
			values = append(values, 0)
		case "B":
			values = append(values, 1)
		case "C":
			values = append(values, 3)
		case "D":
			values = append(values, 5)
		case "x":
			empty++
		default:
			panic("Unknown value hit!")
		}
		if count == 3 {
			switch empty {
			case 0:
				values = append(values, 6)
			case 1:
				values = append(values, 2)
			case 2:
			}
			count = 1
		} else {
			count++
		}
	}
	return values
}
