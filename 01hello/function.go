package main

import "fmt"

func main() {
	result := multiply(2, 6)
	fmt.Println(result)

	resultSum, mylen, myname := adder(2, 6, 5, 8, 9)
	fmt.Println(resultSum, mylen, myname)
}

func multiply(v1, v2 int) int {
	return v1 * v2
}

func adder(values ...int) (int, int, string) {
	sum := 0
	for i := range values {
		sum = sum + values[i]
	}
	length := len(values)
	name := "just for fun"
	return sum, length, name
}
