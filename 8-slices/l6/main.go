package main

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i, m := range messages {
		costs[i] = float64(len(m)) * 0.01
	}

	return costs
}
