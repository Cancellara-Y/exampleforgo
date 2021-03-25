package main

import (
	"fmt"
	"net/url"
)

func main() {
	urr := "https://www.baidu.com"
	r, err := url.Parse(urr)

	fmt.Println(err, r.Host, r.Path, r.RawPath, " ersdfr")
}
