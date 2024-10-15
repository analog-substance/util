package string_utils

import (
	"sort"
	"strings"
)

func SortedLowerUnique(s []string) []string {
	lineMap := make(map[string]bool)
	for _, v := range s {
		lineMap[strings.TrimSpace(strings.ToLower(v))] = true
	}
	s = []string{}
	for k := range lineMap {
		s = append(s, k)
	}
	sort.Strings(s)

	return s
}
