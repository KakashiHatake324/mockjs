package mockjs

import (
	"crypto/rand"
	"math/big"
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
