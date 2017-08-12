package ms

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var s int64 = 1000
var m int64 = s * 60
var h int64 = m * 60
var d int64 = h * 24
var y int64 = d*365 + d/4

/**
 * Parse or format the given `val`.
 *
 * Options:
 *
 *  - `long` verbose formatting [false]
 *
 * @param {string} val
 * @panic {Error} throw an error if val is not a non-empty string or a number
 * @return {int64}
 * @api public
 */
func Parse(val string) int64 {
	if len(val) > 0 {
		return process(val)
	}
	err := errors.New("val is not a non-empty string or a valid number")
	panic(err)
}

/**
 * Process the given `str` and return milliseconds.
 *
 * @param {String} str
 * @return {int64}
 * @api private
 */
func process(val string) int64 {
	if len(val) > 100 {
		return 0
	}
	re := regexp.MustCompile(`^((?:\d+)?\.?\d+) *(milliseconds?|msecs?|ms|seconds?|secs?|s|minutes?|mins?|m|hours?|hrs?|h|days?|d|years?|yrs?|y)?$`)
	match := re.FindStringSubmatch(val)
	if match == nil {
		return 0
	}
	n, _ := strconv.ParseInt(match[1], 10, 64)
	if match[2] == "" {
		match[2] = "ms"
	}
	timeType := strings.ToLower(match[2])

	switch timeType {
	case "years", "year", "yrs", "yr", "y":
		return n * y
	case "days", "day", "d":
		return n * d
	case "hours", "hour", "hrs", "hr", "h":
		return n * h
	case "minutes", "minute", "mins", "min", "m":
		return n * m
	case "seconds", "second", "secs", "sec", "s":
		return n * s
	case "milliseconds", "millisecond", "msecs", "msec", "ms":
		return n
	default:
		return 0
	}
}
