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

func (*Windows) TypeOf(v interface{}) string {
	typeOf := fmt.Sprintf("%T", v)
	typeOf = strings.ReplaceAll(typeOf, "bool", "boolean")
	typeOf = strings.ReplaceAll(typeOf, "int", "number")
	return typeOf
}

func (*Windows) FromCharcode(v []any) string {
	var result string
	for _, codePoint := range v {
		result += string(rune(Math.ToInt64(codePoint)))
	}
	return result
}

func (*Windows) Sha(sh int, s string) string {
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

func (*Windows) Md5(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
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
