package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var mystring string
	// fmt.Scanln(&mystring)
	// fmt.Println(mystring)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter your full name :")
	// myname, _ := reader.ReadString('\n')
	// fmt.Println(myname)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating :")
	myrating, _ := reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Println(mynumRating + 2)
}
