package mockjs

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Window struct {
	StartTime int64 // lets make it time.now().milli - like 1000-500 ms for page load time?
	WindowFunctions
	JSON JSON
}

type JSON struct {
}

type Interval chan bool

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

func InitWindow() *Window {
	return &Window{
		StartTime: time.Now().UnixMilli(),
	}
}

func (w *Window) Btoa(s string) string {
	data := []byte(s)
	str := base64.StdEncoding.EncodeToString(data)
	return str
}

func (w *Window) Atob(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

func (w *Window) TypeOf(v interface{}) string {
	typeOf := fmt.Sprintf("%T", v)
	typeOf = strings.ReplaceAll(typeOf, "bool", "boolean")
	typeOf = strings.ReplaceAll(typeOf, "int", "number")
	return typeOf
}

func (w *Window) FromCharcode(v []any) string {
	var result string
	for _, codePoint := range v {
		result += string(rune(Math.ToInt64(codePoint)))
	}
	return result
}

func (w *Window) Sha(sh int, s string) string {
	var h hash.Hash
	switch sh {
	case 1:
		h = sha1.New()
	case 256:
		h = sha256.New()
	case 512:
		h = sha512.New()
	default:
		h = sha256.New()
	}
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (w *Window) Md5(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func (w *Window) EncodeURIComponent(s string) string {
	return url.QueryEscape(s)
}

func (w *Window) CharcodeAt(s string, n int) int {
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
func (w *Window) NewDate() int64 {
	return time.Now().UnixMilli()
}

// new date string timestamp
func (w *Window) NewDateString() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}

// get the current performance
func (w *Window) Performance() int64 {
	return time.Now().UnixMilli() - w.StartTime
}

func (w *Window) SetTimeout(fn func(), delay time.Duration) *time.Timer {
	return time.AfterFunc(delay, fn)
}

func (w *Window) SetInterval(fn func(), interval time.Duration) Interval {
	ticker := time.NewTicker(interval)
	stop := make(chan bool)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fn()
			case <-stop:
				return
			}
		}
	}()
	return stop
}

func (w *Window) ClearInterval(interval Interval) {
	interval <- true
}
