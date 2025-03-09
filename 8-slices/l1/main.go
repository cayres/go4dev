package main

import (
	"fmt"
)

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	var costs [3]int
	prevCost := 0

	for i, m := range messages {
		fmt.Println(i)
		cost := len(m) + prevCost
		costs[i] = cost
		prevCost = cost
	}

	return messages, costs
}
