package main

import "fmt"

func main() {

	score := make(map[string]int)
	score["ankit"] = 78
	fmt.Println(score)
	score["ankit"] = 58
	score["sumit"] = 65
	score["barry"] = 65
	score["sneha"] = 85

	fmt.Println(score)
	fmt.Println(score["sumit"])
	delete(score, "sneha")
	fmt.Println(score)

	for k, v := range score {
		fmt.Printf("Score of %v is %v \n", k, v)
	}
}
