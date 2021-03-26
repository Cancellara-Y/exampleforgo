package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func First() {

	s := "a%\xc5z"
	ss := strings.ToValidUTF8(s, "")
	fmt.Println(ss)
	fmt.Printf("%q\n", s)
	if !utf8.ValidString(s) {
		v := make([]rune, 0, len(s))
		for i, r := range s {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(s[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		s = string(v)
	}
	fmt.Printf("%q\n", s)
}

func Second() {
	str := "——языкほん中蚊#$%^&!@*()_~_+=-123ENG_-AA.-—-。ΦφϕkKK"
	tmp := ""
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			tmp += string(r)
			fmt.Printf("汉字%v\n", string(r))
		} else if unicode.IsLetter(r) {
			fmt.Printf("Letter：%v\n", string(r))
			tmp += string(r)
		} else if unicode.IsDigit(r) {
			fmt.Printf("数字%v\n", string(r))
			tmp += string(r)
		} else if unicode.IsSymbol(r) {
			tmp += string(r)
			fmt.Printf("特殊字符%v\n", string(r))
		} else if unicode.IsPunct(r) {
			if r == '-' || r == '_' || r == '.' {
				fmt.Printf("可接受标点符号: %v\n", string(r))
				tmp += string(r)
			}
			fmt.Printf("标点符号：%v\n", string(r))
		} else {
			fmt.Printf("认不出来%v\n", string(r))
		}
	}
	fmt.Println(tmp)
}

func CheckFileName(filename string) string {

	tmp := ""
	for _, r := range filename {
		if unicode.Is(unicode.Han, r) || unicode.IsLetter(r) || unicode.IsDigit(r) {
			tmp += string(r)
		}

		if r == '-' || r == '_' || r == '.' || r == '—' {
			tmp += string(r)
		}
	}
	return tmp
}

func main() {
	var i int = 0
	var t int = 0
	for {
		fmt.Printf("%c", i)
		i++
		//从零一直打印
		time.Sleep(time.Nanosecond)
		//如果打印的太快，有时会不出结果，所以要停顿一下
		//以下几行是每隔60个换一下行，方便观察结果
		t++
		if t%60 == 0 {
			fmt.Println()
		}
	}
}
