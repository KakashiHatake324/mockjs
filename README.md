`go get github.com/KakashiHatake324/mockjs`

I started this to make it easier on myself when converting JS functions to golang projects. Feel free to use it, fork it, contribute and do anything you would like with it. I've found it helpful when converting Akamai functions to golang.

`MATH FUNCTIONS`
```go
Math.pow_Example := mockjs.Math.Pow(10, 40.50)
Math.floor_Example := mockjs.Math.Floor(5)
Math.random_Example := mockjs.Math.Random()
Math.round_Example := mockjs.Math.Round(100.5)
Math.max_Example := mockjs.Math.Max([]int{2,3,5,6,7})
```

`WINDOW FUNCTIONS`
```go
window := mockjs.Window.InitWindow()

Window.btoa_Example := window.Btoa("hello world")
Window.atob_Example := window.Btoa("aGVsbG8gd29ybGQ=")
Window.newDate_Example := window.NewDate() // int64 timestamp in unix milli
Window.newDateString_Example := window.NewDateString()
Window.encodeUriComponent_Example := window.EncodeURIComponent("https://github.com/")
Window.charcodeAt_Example := window.CharcodeAt("charstring",3)
Window.performanceNow_Example := window.Performance() // time since init window in ms
Window.parseInt_Example := mockjs.Math.ParseInt(3, 10)
```

`ARRAY FUNCTIONS`
```go
Array.map_Example := mockjs.Array.Map([]interface{"4",5,"2",6,7,8}, mockjs.Number)
Array.pop_Example := mockjs.Array.Pop([]interface{"4",5,"2",6,7,8})
Array.pop_Example := mockjs.Array.PopString([]interface{"4","2"})
Array.pop_Example := mockjs.Array.PopInt([]interface{5,6,7,8})
```

`BUILT IN FUNCTIONS`
```go
Random_Number_Between := mockjs.Math.NumberBetween(5, 12)
```