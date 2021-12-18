package main

import "fmt"

func main() {
	p := 5
	typeP := &p
	p = p * 5
	fmt.Println(p)
	fmt.Println(*typeP)

}
