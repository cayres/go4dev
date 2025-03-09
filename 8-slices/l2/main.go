package main

import (
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == "pro" {
		return messages[:], nil
	}

	if plan == "free" {
		return messages[:2], nil
	}

	return nil, fmt.Errorf("unsupported plan")
}
