package main

import "strings"

type StringsBuffer []string

// Swap is part of sort.Interface.
func (sb *StringsBuffer) Swap(i, j int) {
	(*sb)[i], (*sb)[j] = (*sb)[j], (*sb)[i]
}

// Len is part of sort.Interface.
func (sb *StringsBuffer) Len() int {
	return len(*sb)
}

func (sb *StringsBuffer) Less(i, j int) bool {
	return strings.Compare((*sb)[i], (*sb)[j]) == -1
}
