package slice

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