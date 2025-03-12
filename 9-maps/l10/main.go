package main

func getNameCounts(names []string) map[rune]map[string]int {
	result := make(map[rune]map[string]int)
	for _, name := range names {
		initial := []rune(name)[0]
		if _, ok := result[initial]; !ok {
			result[initial] = make(map[string]int)
		}
		result[initial][name]++
	}
	return result
}
