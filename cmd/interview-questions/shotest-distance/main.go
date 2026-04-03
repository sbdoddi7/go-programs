package main

func shortestDistance(words []string, word1 string, word2 string) int {
	minDistance := len(words)
	lastWord1Index := -1
	lastWord2Index := -1

	for i, word := range words {
		if word == word1 {
			lastWord1Index = i
			if lastWord2Index != -1 {
				distance := abs(lastWord1Index - lastWord2Index)
				if distance < minDistance {
					minDistance = distance
				}
			}
		} else if word == word2 {
			lastWord2Index = i
			if lastWord1Index != -1 {
				distance := abs(lastWord2Index - lastWord1Index)
				if distance < minDistance {
					minDistance = distance
				}
			}
		}
	}

	return minDistance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	words := []string{"practice", "makes", "perfect", "coding", "makes"}
	word1 := "coding"
	word2 := "practice"
	println(shortestDistance(words, word1, word2)) // Output: 3
}
