package main

import "fmt"

func RemoveRepeatedString(arr []string) (newArr []string) {
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func main() {
	strs := []string{"sdf", "sdf", "sdfsdf", "sf", "xx", "xx"}
	fmt.Println(strs)

	fmt.Println(RemoveRepeatedString(strs))
}
