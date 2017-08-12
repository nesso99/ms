package ms

import "errors"

var s = 1000
var m = s * 60
var h = m * 60
var d = h * 24
var y = d*365 + d/4

func ms(val string) int {
	if len(val) > 0 {
		return parse(val)
	}
	err := errors.New("val is not a non-empty string or a valid number")
	panic(err)
}

func parse(val string) int {
	if len(val) > 100 {
		return -1
	}
	return 1
}
