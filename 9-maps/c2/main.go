package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	friends := make(map[string]struct{})
	for _, friend := range friendships[username] {
		friends[friend] = struct{}{}
	}

	suggestedFriends := make(map[string]struct{})
	for _, friend := range friendships[username] {
		for _, suggestedFriend := range friendships[friend] {
			if suggestedFriend == username {
				continue
			}

			if _, ok := friends[suggestedFriend]; ok {
				continue
			}
			suggestedFriends[suggestedFriend] = struct{}{}
		}
	}

	var result []string
	for friend := range suggestedFriends {
		result = append(result, friend)
	}
	return result
}
