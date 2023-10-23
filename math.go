package mockjs

import (
	"fmt"
	"math"
	"strconv"
)

var Math MathStruct

type MathStruct struct {
	MathFunctions
}

type MathFunctions interface {
	NumberBetween(float64) float64
	Round(float64) float64
	Floor(float64) float64
	Random() float64
	Sqrt(float64)
	ToString(interface{}) string
	ToInt(interface{}) int
	ToFloat64(interface{})
	Max([]interface{}) int
	Ceil(float64)
	Pow(float64, float64)
}

// Returns a supplied numeric expression rounded to the nearest integer.
func (*MathStruct) NumberBetween(min, max float64) float64 {
	return Random_Float_Range(min, max)
}

// Returns a pseudorandom number between 0 and 1.
func (*MathStruct) Random() float64 {
	return Random_Float_Range(0, 1)
}

// Returns the greatest integer less than or equal to its numeric argument.
func (*MathStruct) Floor(max interface{}) float64 {
	switch c := max.(type) {
	case float64:
		return float64(Random_range(0, int(c)))
	case int:
		return float64(Random_range(0, c))
	case int32:
		return float64(Random_range(0, int(c)))
	case int64:
		return float64(Random_range(0, int(c)))
	default:
		return 0
	}
}

// Returns a supplied numeric expression rounded to the nearest integer.
func (*MathStruct) Round(round float64) float64 {
	return math.Round(round)
}

// Returns the square root of a number.
func (*MathStruct) Sqrt(sqrt float64) float64 {
	return math.Sqrt(sqrt)
}

// Returns the smallest integer greater than or equal to its numeric argument.
func (*MathStruct) Ceil(c float64) float64 {
	return math.Ceil(c)
}

// Returns the larger of a set of supplied numeric expressions.
func (*MathStruct) Max(c interface{}) int {
	var max int
	switch f := c.(type) {
	case []int:
		for _, o := range f {
			if max < Math.ToInt(o) {
				max = Math.ToInt(o)
			}
		}
	case []int32:
		for _, o := range f {
			if max < Math.ToInt(o) {
				max = Math.ToInt(o)
			}
		}
	case []int64:
		for _, o := range f {
			if max < Math.ToInt(o) {
				max = Math.ToInt(o)
			}
		}
	case []float32:
		for _, o := range f {
			if max < Math.ToInt(o) {
				max = Math.ToInt(o)
			}
		}
	case []float64:
		for _, o := range f {
			if max < Math.ToInt(o) {
				max = Math.ToInt(o)
			}
		}
	}
	return max
}

// Returns the value of a base expression taken to a specified power.
func (*MathStruct) Pow(o, p float64) float64 {
	return math.Pow(o, p)
}

// Returns a string representation of an object.
func (*MathStruct) ToString(v interface{}) string {
	switch c := v.(type) {
	case int:
		return strconv.Itoa(c)
	case int32:
		return strconv.Itoa(int(c))
	case int64:
		return strconv.Itoa(int(c))
	case float32:
		return fmt.Sprintf("%f", c)
	case float64:
		return strconv.FormatFloat(c, 'f', -1, 64)
	case string:
		return c
	default:
		return ""
	}
}

// Returns a float64 representation of an object.
func (*MathStruct) ToFloat64(v interface{}) float64 {
	switch c := v.(type) {
	case int:
		return float64(c)
	case int32:
		return float64(c)
	case int64:
		return float64(c)
	case float32:
		return float64(c)
	case float64:
		return c
	case string:
		f, _ := strconv.ParseFloat(c, 64)
		return f
	default:
		return 0
	}
}

// Returns an int representation of an object.
func (*MathStruct) ToInt(v interface{}) int {
	switch c := v.(type) {
	case int:
		return int(c)
	case int32:
		return int(c)
	case int64:
		return int(c)
	case float32:
		return int(c)
	case float64:
		return int(c)
	case string:
		r, _ := strconv.Atoi(c)
		return r
	default:
		return 0
	}
}

func (*MathStruct) ParseInt(v, c interface{}) int64 {
	solved, _ := strconv.ParseInt(Math.ToString(Math.ToInt(v)), Math.ToInt(c), 64)
	return solved
}
