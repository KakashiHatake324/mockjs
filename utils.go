package mockjs

import (
	"math/rand"
	"time"
)

func Random_range(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func RandomFloat() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}

func Random_Float_Range(minimum, maximum float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return (rand.Float64() * (maximum - minimum)) + minimum
}
