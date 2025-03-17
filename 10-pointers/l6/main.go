package main

import (
	"strings"
)

func removeProfanity(message *string) {
	if message == nil {
		return
	}

	profanity := []string{"fubb", "shiz", "witch"}
	for _, word := range profanity {
		*message = strings.ReplaceAll(*message, word, strings.Repeat("*", len(word)))
	}
}
