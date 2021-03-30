package slice

import "fmt"

type AssistSlice struct {

}

func (a *AssistSlice) RemoveRepeatedString(src []string)( dest []string ){
	uniqMap := make(map[string]struct{})

	for _, v := range src {
		uniqMap[v] = struct{}{}
	}

	for k, _ := range uniqMap {
		dest = append(dest, k)
	}
	return
}


func (a *AssistSlice) Strings2Map(src []string) map[string]struct{} {
	uniqMap := make(map[string]struct{})

	for _, v := range src {
		uniqMap[v] = struct{}{}
	}
	return uniqMap
}

func (a *AssistSlice)GetSubStrByLen(str string, len int) string {
	var n, i int
	for i = range str {
		fmt.Println("i: ", i)
		if n == len {
			break
		}
		n++
	}
	s := str[:i]
	fmt.Println(s)
	return s
}