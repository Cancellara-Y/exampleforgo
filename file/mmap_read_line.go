package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("/Users/yjf/tmp/gen.sh")
	if err != nil {
		fmt.Println(err)
		return
	}


	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		handler(line)

		if err != nil {
			if err == io.EOF {
				fmt.Println(err)
				return
			}
			fmt.Println(err)
			return
		}
	}
}

func handler (line string) {
	fmt.Println(line)
}