package ms

import (
	"testing"
)

// ParseMS(string)
// short string
func TestPreserveMS(t *testing.T) {
	x := ParseMs("100")
	if x != 100 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 100)
	}
}

func TestM2MS(t *testing.T) {
	x := ParseMs("1m")
	if x != 60000 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 60000)
	}
}

func TestH2MS(t *testing.T) {
	x := ParseMs("1h")
	if x != 3600000 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 3600000)
	}
}

func TestD2MS(t *testing.T) {
	x := ParseMs("2d")
	if x != 172800000 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 172800000)
	}
}

func TestS2MS(t *testing.T) {
	x := ParseMs("1s")
	if x != 1000 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 1000)
	}
}

func TestMS2MS(t *testing.T) {
	x := ParseMs("100ms")
	if x != 100 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 100)
	}
}

func TestDECIMA(t *testing.T) {
	x := ParseMs("1.5h")
	if x != 5400000 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 5400000)
	}
}

func TestMultipleSpaces(t *testing.T) {
	x := ParseMs("1   s")
	if x != 1000 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 1000)
	}
}

func TestCaseIntensitive(t *testing.T) {
	x := ParseMs("1.5H")
	if x != 5400000 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 5400000)
	}
}

func TestRound2MS(t *testing.T) {
	x := ParseMs("1.5ms")
	if x != 1 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 1)
	}
}

// long string
func TestMillisecond2MS(t *testing.T) {
	x := ParseMs("53 milliseconds")
	if x != 53 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 53)
	}
}

func TestMsec2MS(t *testing.T) {
	x := ParseMs("17 msecs")
	if x != 17 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 17)
	}
}

func TestSec2MS(t *testing.T) {
	x := ParseMs("1 sec")
	if x != 1000 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 1000)
	}
}

func TestMin2MS(t *testing.T) {
	x := ParseMs("1 min")
	if x != 60000 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 60000)
	}
}

func TestHr2MS(t *testing.T) {
	x := ParseMs("1 hr")
	if x != 3600000 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 3600000)
	}
}

func TestDay2MS(t *testing.T) {
	x := ParseMs("2 days")
	if x != 172800000 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 172800000)
	}
}

func TestDecimaLong(t *testing.T) {
	x := ParseMs("1.5 hours")
	if x != 5400000 {
		t.Errorf("MS was incorrect, got: %d, want: %d.", x, 5400000)
	}
}

// ParseString(int64)
func TestSupportMS(t *testing.T) {
	x := ParseString(500, true)
	if x != "500 ms" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x, "500 ms")
	}
}

func TestSupportSecond(t *testing.T) {
	x := ParseString(1000, true)
	if x != "1 second" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x, "1 second")
	}
	x1 := ParseString(1200, true)
	if x1 != "1 second" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x1, "1 second")
	}
	x2 := ParseString(10000, true)
	if x2 != "10 seconds" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x2, "10 seconds")
	}
}

func TestSupportMinute(t *testing.T) {
	x := ParseString(60*1000, true)
	if x != "1 minute" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x, "1 minute")
	}
	x1 := ParseString(60*1200, true)
	if x1 != "1 minute" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x1, "1 minute")
	}
	x2 := ParseString(60*10000, true)
	if x2 != "10 minutes" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x2, "10 minutes")
	}
}

func TestSupportHour(t *testing.T) {
	x := ParseString(60*60*1000, true)
	if x != "1 hour" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x, "1 hour")
	}
	x1 := ParseString(60*60*1200, true)
	if x1 != "1 hour" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x1, "1 hour")
	}
	x2 := ParseString(60*60*10000, true)
	if x2 != "10 hours" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x2, "10 hours")
	}
}

func TestSupportDay(t *testing.T) {
	x := ParseString(24*60*60*1000, true)
	if x != "1 day" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x, "1 day")
	}
	x1 := ParseString(24*60*60*1200, true)
	if x1 != "1 day" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x1, "1 day")
	}
	x2 := ParseString(24*60*60*10000, true)
	if x2 != "10 days" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x2, "10 days")
	}
}

func TestSupportRound(t *testing.T) {
	x := ParseString(234234234, true)
	if x != "3 days" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x, "3 days")
	}
}

func TestSupportShortMS(t *testing.T) {
	x := ParseString(500, false)
	if x != "500ms" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x, "500ms")
	}
}

func TestSupportS(t *testing.T) {
	x := ParseString(1000, false)
	if x != "1s" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x, "1s")
	}
	x2 := ParseString(10000, false)
	if x2 != "10s" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x2, "10s")
	}
}

func TestSupportM(t *testing.T) {
	x := ParseString(60*1000, false)
	if x != "1m" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x, "1m")
	}
	x2 := ParseString(60*10000, false)
	if x2 != "10m" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x2, "10m")
	}
}

func TestSupportH(t *testing.T) {
	x := ParseString(60*60*1000, false)
	if x != "1h" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x, "1h")
	}
	x2 := ParseString(60*60*10000, false)
	if x2 != "10h" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x2, "10h")
	}
}

func TestSupportD(t *testing.T) {
	x := ParseString(24*60*60*1000, false)
	if x != "1d" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x, "1d")
	}
	x2 := ParseString(24*60*60*10000, false)
	if x2 != "10d" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x2, "10d")
	}
}

func TestRoundShort(t *testing.T) {
	x := ParseString(234234234, false)
	if x != "3d" {
		t.Errorf("MS was incorrect, got: %s, want: %s.", x, "3d")
	}
}
