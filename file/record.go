package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	fi, err := os.Open("/")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
	}
}
