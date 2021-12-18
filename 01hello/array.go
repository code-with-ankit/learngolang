package main

import "fmt"

func main() {
	var numbers [5]string
	numbers[0] = "uno"
	numbers[1] = "dos"
	numbers[2] = "tres"

	fmt.Println(numbers)

	var colors = [3]string{"red", "yellow", "green"}
	fmt.Println(colors)
}
