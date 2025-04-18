package mockjs

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func Random_range(min, max int) int {
	rangeVal := max - min + 1
	if rangeVal <= 0 {
		return 0
	}

	// Generate a random number in the range [0, rangeVal)
	n, err := rand.Int(rand.Reader, big.NewInt(int64(rangeVal)))
	if err != nil {
		return 0
	}

	return min + int(n.Int64())
}

func RandomFloat() float64 {
	numerator, err := rand.Int(rand.Reader, big.NewInt(1<<53))
	if err != nil {
		return 0
	}
	denominator := float64(1 << 53)
	randomFloat := float64(numerator.Int64()) / denominator

	return randomFloat
}

func Random_Float_Range(min, max float64) float64 {
	if min >= max {
		return 0
	}
	// Generate a random number between 0 and 1
	numerator, err := rand.Int(rand.Reader, big.NewInt(1<<53))
	if err != nil {
		return 0
	}

	denominator := float64(1 << 53)
	randomFloat := float64(numerator.Int64()) / denominator

	// Scale to the desired range
	return min + randomFloat*(max-min)
}

func normalizeString(s string) string {
	// Convert to lowercase and remove special characters
	s = strings.ToLower(s)
	s = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r == ' ' {
			return r
		}
		return -1
	}, s)
	// Remove extra spaces
	s = strings.Join(strings.Fields(s), " ")
	return strings.ReplaceAll(s, ".", "")
}

func StringsSimilar(s1, s2 string) bool {
	s1 = normalizeString(s1)
	s2 = normalizeString(s2)

	// If one string contains the other, they're similar enough
	if strings.Contains(s1, s2) || strings.Contains(s2, s1) {
		return true
	}

	// Split into words and check if most words match
	words1 := strings.Fields(s1)
	words2 := strings.Fields(s2)

	// Count matching words
	matches := 0
	for _, w1 := range words1 {
		for _, w2 := range words2 {
			if w1 == w2 {
				matches++
				break
			}
		}
	}

	// If more than 70% of words match, consider them similar
	minWords := len(words1)
	if len(words2) < minWords {
		minWords = len(words2)
	}
	return float64(matches)/float64(minWords) >= 0.7
}
