package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wStruct struct {
	num  int
	word string
}

func sortSlice(slice []wStruct) []wStruct {
	sort.SliceStable(slice, func(i, j int) bool { return slice[i].word < slice[j].word })
	sort.SliceStable(slice, func(i, j int) bool { return slice[i].num > slice[j].num })
	return slice
}

func Top10(inStr string) []string {
	if inStr == "" {
		return nil
	}
	wE := wStruct{}
	words := []wStruct{}
	var n int
	workSlice := strings.Fields(inStr)
	outSlice := []string{}
	wMap := make(map[string]int)
	for i := range workSlice {
		wMap[workSlice[i]] = 0
	}
	for i := range workSlice {
		if _, inMap := wMap[workSlice[i]]; inMap {
			wMap[workSlice[i]]++
		}
	}
	for key, val := range wMap {
		wE.word = key
		wE.num = val
		words = append(words, wE)
	}
	words = sortSlice(words)
	if len(words) < 10 {
		n = len(words)
	} else {
		n = 10
	}
	sorted := words[:n]
	sorted = sortSlice(sorted)
	for i := range sorted {
		outSlice = append(outSlice, sorted[i].word)
	}
	return outSlice
}
