package slice

import (
	"fmt"
)

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

func (a *AssistSlice) Map2String(src interface{}) (dest []string) {

	switch r := src.(type) {
	case map[int]string:
		tmpMap := src.(map[int]string)
		for _, v := range tmpMap {
			dest = append(dest, v)
		}
	case 	map[int8]string:
		tmpMap := src.(map[int8]string)
		for _, v := range tmpMap {
			dest = append(dest, v)
		}
	case 	map[int16]string:
		tmpMap := src.(map[int16]string)
		for _, v := range tmpMap {
			dest = append(dest, v)
		}
	case 	map[int32]string:
		tmpMap := src.(map[int32]string)
		for _, v := range tmpMap {
			dest = append(dest, v)
		}
	case 	map[int64]string:
		tmpMap := src.(map[int64]string)
		for _, v := range tmpMap {
			dest = append(dest, v)
		}
	case 	map[uint]string:
		tmpMap := src.(map[uint]string)
		for _, v := range tmpMap {
			dest = append(dest, v)
		}
	case 	map[uint8]string:
		tmpMap := src.(map[uint8]string)
		for _, v := range tmpMap {
			dest = append(dest, v)
		}
	case 	map[uint16]string:
		tmpMap := src.(map[uint16]string)
		for _, v := range tmpMap {
			dest = append(dest, v)
		}
	case 	map[uint32]string:
		tmpMap := src.(map[uint32]string)
		for _, v := range tmpMap {
			dest = append(dest, v)
		}
	case 	map[uint64]string:
		tmpMap := src.(map[uint64]string)
		for _, v := range tmpMap {
			dest = append(dest, v)
		}
	case 	map[string]string:
		tmpMap := src.(map[string]string)
		for _, v := range tmpMap {
			dest = append(dest, v)
		}
	case 	map[float32]string:
		tmpMap := src.(map[float32]string)
		for _, v := range tmpMap {
			dest = append(dest, v)
		}
	case 	map[float64]string:
		tmpMap := src.(map[float64]string)/**/
		for _, v := range tmpMap {
			dest = append(dest, v)
		}
	default:
		fmt.Println("do nothing src type: " , r)
	}
	return
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