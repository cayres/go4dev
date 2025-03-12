package main

func getCounts(messagedUsers []string, validUsers map[string]int) {
	userCounts := make(map[string]int)
	for _, userName := range messagedUsers {
		userCounts[userName]++
	}

	for userName, count := range userCounts {
		if _, ok := validUsers[userName]; ok {
			validUsers[userName] += count
		}
	}
}
