package main

import (
	"fmt"
	"sort"
)

func main() {

	data := []string{
		"Abc", "zxc", "ghij",
	}
	sort.Strings(data)
	fmt.Println(data)

}
