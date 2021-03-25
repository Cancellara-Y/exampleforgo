package main

import "fmt"

func main() {

	a := []string{"11", "22", "33"}

	b := a

	a = []string{}

	fmt.Println(b)
}
