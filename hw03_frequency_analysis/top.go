package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wStruct struct {
	num  int
	word string
}

//func for structure sorting
func sortSlice(slice []wStruct) []wStruct {
	sort.SliceStable(slice, func(i, j int) bool {
		if slice[i].num != slice[j].num {
			return slice[i].num > slice[j].num
		}
		return slice[i].word < slice[j].word
	})
	return slice
}

func Top10(inStr string) []string {
	//return empty slice if string is empty
	if inStr == "" {
		return nil
	}
	//get work slice
	workSlice := strings.Fields(inStr)
	//init and fill work map
	wMap := make(map[string]int)
	for i := range workSlice {
		wMap[workSlice[i]]++
	}
	//init and fill structure to sort
	words := make([]wStruct, len(wMap))
	for key, val := range wMap {
		wE := wStruct{}
		wE.word = key
		wE.num = val
		words = append(words, wE)
	}
	//sort structure
	words = sortSlice(words)
	n := 10
	if len(words) < 10 {
		n = len(words)
	}
	//cut sorted structure
	words = words[:n]
	//get slice to return
	outSlice := []string{}
	for i := range words {
		outSlice = append(outSlice, words[i].word)
	}
	return outSlice
}
