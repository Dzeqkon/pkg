package slice

import "strings"

/*
[9 9 8 4 2 9 1 7 - a 5 4 b - 3 3 1 6 - c d f 3 - 8 7 d 9 f b 5 7] -> "99842917-a54b-3316-cdf3-87d9fb57"
*/
func ArrayToString(arrays []string) string {
	return strings.Join(arrays, "")
}

func Diff(base, exclude []string) (result []string) {
	excludeMap := make(map[string]bool)
	for _, s := range exclude {
		excludeMap[s] = true
	}
	for _, s := range base {
		if !excludeMap[s] {
			result = append(result, s)
		}
	}
	return result
}

func Unique(ss []string) (result []string) {
	smap := make(map[string]bool)
	for _, s := range ss {
		smap[s] = true
	}
	for s := range smap {
		result = append(result, s)
	}
	return result
}

func FindString(array []string, str string) int {
	for index, s := range array {
		if str == s {
			return index
		}
	}
	return -1
}

func StringIn(str string, array []string) bool {
	return FindString(array, str) > -1
}
