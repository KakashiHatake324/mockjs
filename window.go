package mockjs

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

var Window Windows

type Windows struct {
	StartTime int64 // lets make it time.now().milli - like 1000-500 ms for page load time?
	WindowFunctions
	JSON JSON
}

type JSON struct {
}

func (*JSON) Stringify(s interface{}) string {
	d, _ := json.Marshal(s)
	return string(d)
}

func (*JSON) Parse(s string) interface{} {
	var f interface{}
	return json.Unmarshal([]byte(s), &f)
}

type WindowFunctions interface {
	Btoa(string) string
	Atob(string) string
	NewDate() int64
	NewDateString() string
	Date() string
	EncodeURIComponent(string) string
	CharcodeAt(string, int) int
	Performance() int64
}

func InitWindow() *Windows {
	return &Windows{
		StartTime: time.Now().UnixMilli(),
	}
}

func (*Windows) Btoa(s string) string {
	data := []byte(s)
	str := base64.StdEncoding.EncodeToString(data)
	return str
}

func (*Windows) Atob(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

func (*Windows) EncodeURIComponent(s string) string {
	return url.QueryEscape(s)
}

func (*Windows) CharcodeAt(s string, n int) int {
	// Convert the string to a rune slice
	runes := []rune(s)
	// Check if the index is within bounds
	if n >= 0 && n < len(runes) {
		// Get the Unicode code point of the character at the specified index
		return int(runes[n])
	} else {
		return 0
	}
}

// new date timestamp
func (*Windows) NewDate() int64 {
	return time.Now().UnixMilli()
}

// new date string timestamp
func (*Windows) NewDateString() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}

// get the current performance
func (w *Windows) Performance() int64 {
	return time.Now().UnixMilli() - w.StartTime
}
