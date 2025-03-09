package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, w := range msg {
		for _, b := range badWords {
			if b == w {
				return i
			}
		}
	}
	return -1
}
