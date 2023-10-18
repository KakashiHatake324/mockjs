package mockjs

var (
	Array  ArrayStruct
	Number MathTypes = "number"
	String MathTypes = "string"
)

type ArrayStruct struct {
}

type MathTypes string

type ArrayFunctions interface {
	Map(interface{}, MathTypes) interface{}
	Pop(interface{}) interface{}
	PopInt([]int) []int
	PopString([]string) []string
}

// Equivilant to js's array pop function
func (*ArrayStruct) Map(array interface{}, turnTo MathTypes) interface{} {
	if array == nil {
		return nil
	}
	switch turnTo {
	case Number:
		switch d := array.(type) {
		case []interface{}:
			var convert []int
			for _, v := range d {
				convert = append(convert, Math.ToInt(v))
			}
			return convert
		case []string:
			var convert []int
			for _, v := range d {
				convert = append(convert, Math.ToInt(v))
			}
			return convert
		case []int:
			return d
		case []int32:
			var convert []int
			for _, v := range d {
				convert = append(convert, Math.ToInt(v))
			}
			return convert
		case []int64:
			var convert []int
			for _, v := range d {
				convert = append(convert, Math.ToInt(v))
			}
			return convert
		case []float32:
			var convert []int
			for _, v := range d {
				convert = append(convert, Math.ToInt(v))
			}
			return convert
		case []float64:
			var convert []int
			for _, v := range d {
				convert = append(convert, Math.ToInt(v))
			}
			return convert
		default:
			return nil
		}
	case String:
		switch d := array.(type) {
		case []interface{}:
			var convert []string
			for _, v := range d {
				convert = append(convert, Math.ToString(v))
			}
			return convert
		case []string:
			return d
		case []int:
			var convert []string
			for _, v := range d {
				convert = append(convert, Math.ToString(v))
			}
			return convert
		case []int32:
			var convert []string
			for _, v := range d {
				convert = append(convert, Math.ToString(v))
			}
			return convert
		case []int64:
			var convert []string
			for _, v := range d {
				convert = append(convert, Math.ToString(v))
			}
			return convert
		case []float32:
			var convert []string
			for _, v := range d {
				convert = append(convert, Math.ToString(v))
			}
			return convert
		case []float64:
			var convert []string
			for _, v := range d {
				convert = append(convert, Math.ToString(v))
			}
			return convert
		default:
			return nil
		}
	default:
		return nil
	}
}

// Equivilant to js's array pop function
func (*ArrayStruct) Pop(array interface{}) interface{} {
	if array == nil {
		return nil
	}
	switch d := array.(type) {
	case []interface{}:
		return (d)[:len(d)-1]
	case []string:
		return (d)[:len(d)-1]
	case []int:
		return (d)[:len(d)-1]
	case []float64:
		return (d)[:len(d)-1]
	default:
		return nil
	}
}

func (*ArrayStruct) PopInt(array []int) []int {
	if array == nil {
		return nil
	}
	array = (array)[:len(array)-1]
	return array
}

func (*ArrayStruct) Slice(start, end int, array []interface{}) []interface{} {
	if array == nil {
		return nil
	}
	var newArray []interface{}
	for i := 0; i < len(array); i++ {
		if i >= start && i < end {
			newArray = append(newArray, array[i])
		}
	}
	return newArray
}

func (*ArrayStruct) PopString(array []string) []string {
	if array == nil {
		return nil
	}
	array = (array)[:len(array)-1]
	return array
}
