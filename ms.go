package ms

import (
	"errors"
	"math"
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
 * Parse the given string `val` to ms.
 *
 * @param {string} val
 * @panic {Error} throw an error if val is not a non-empty string or a number
 * @return {int64}
 * @api public
 */
func ParseMs(val string) int64 {
	if len(val) > 0 {
		return process(val)
	}
	err := errors.New("val is not a non-empty string or a valid number")
	panic(err)
}

/**
 * Parse the given int `val` to string.
 *
 * @param {string} val
 * @param {bool} longFormat the string is in long format
 * @panic {Error} throw an error if val is not a non-empty string or a number
 * @return {string}
 * @api public
 */
func ParseString(val int64, longFormat bool) string {
	if val > 0 {
		if longFormat {
			return fmtLong(val)
		}
		return fmtShort(val)
	}
	err := errors.New("val is an unsigned integer")
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
	re := regexp.MustCompile(`(?i)^((?:\d+)?\.?\d+) *(milliseconds?|msecs?|ms|seconds?|secs?|s|minutes?|mins?|m|hours?|hrs?|h|days?|d|years?|yrs?|y)?$`)
	match := re.FindStringSubmatch(val)
	if match == nil {
		return 0
	}
	n, _ := strconv.ParseFloat(match[1], 64)
	if match[2] == "" {
		match[2] = "ms"
	}
	timeType := strings.ToLower(match[2])
	switch timeType {
	case "years", "year", "yrs", "yr", "y":
		return int64(n * float64(y))
	case "days", "day", "d":
		return int64(n * float64(d))
	case "hours", "hour", "hrs", "hr", "h":
		return int64(n * float64(h))
	case "minutes", "minute", "mins", "min", "m":
		return int64(n * float64(m))
	case "seconds", "second", "secs", "sec", "s":
		return int64(n * float64(s))
	case "milliseconds", "millisecond", "msecs", "msec", "ms":
		return int64(n)
	default:
		return 0
	}
}

/**
 * Short format for `ms`.
 *
 * @param {int64} ms
 * @return {string}
 * @api private
 */
func fmtShort(ms int64) string {
	var result int64
	if ms >= d {
		result = Round(float64(ms) / float64(d))
		return strconv.FormatInt(result, 10) + "d"
	}
	if ms >= h {
		result = Round(float64(ms) / float64(h))
		return strconv.FormatInt(result, 10) + "h"
	}
	if ms >= m {
		result = Round(float64(ms) / float64(m))
		return strconv.FormatInt(result, 10) + "m"
	}
	if ms >= s {
		result = Round(float64(ms) / float64(s))
		return strconv.FormatInt(result, 10) + "s"
	}
	return strconv.FormatInt(ms, 10) + "ms"
}

/**
 * Long format for `ms`.
 *
 * @param {int64} ms
 * @return {string}
 * @api private
 */

func fmtLong(ms int64) string {
	if ms >= d {
		return plural(ms, d, "day")
	}
	if ms >= h {
		return plural(ms, h, "hour")
	}
	if ms >= m {
		return plural(ms, m, "minute")
	}
	if ms >= s {
		return plural(ms, s, "second")
	}
	return strconv.FormatInt(ms, 10) + " ms"
}

/**
 * Pluralization helper.
 */

func plural(ms int64, n int64, name string) string {
	result := Round(float64(ms) / float64(n))
	if ms < n+n/2 {
		return strconv.FormatInt(result, 10) + " " + name
	}
	return strconv.FormatInt(result, 10) + " " + name + "s"
}

/*
* Round helper
 */
func Round(val float64) int64 {
	_, div := math.Modf(val)
	if div >= 0.5 {
		return int64(math.Ceil(val))
	}
	return int64(math.Floor(val))
}
