package main

import (
	"fmt"
	"sort"
)

func main() {
	var things = []string{"maleta", "ropa", "reloj"}
	fmt.Println(things)

	things = append(things, "bolso")
	fmt.Println(things)

	fmt.Println(append(things[0:4]))

	heros := make([]string, 3, 3)
	heros[0] = "thor"
	heros[1] = "ironman"
	heros[2] = "spiderman"

	fmt.Println(heros)
	heros = append(heros, "deadpool")
	fmt.Println(heros)

	fmt.Println(cap(heros))

	myints := []int{5, 3, 7, 8, 9, 4, 2, 5}

	// isSorted, _ := sort.IntsAreSorted(myints)
	sort.Ints(myints)
	fmt.Println(myints)
}
