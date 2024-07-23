package golang

import (
	"math"
	"time"
)

// AddTime to add hours, minutes, seconds, nanoseconds to time. Not change original time, return new time.
// This function is different from the time.Time.Add method.
//
// 添加小时、分钟、秒、纳秒到时间，该函数不会改变原始时间，而是返回一个新的时间。此方法不同于 time.Time.Add 方法不会增加额外的时间。
func AddTime(t time.Time, hours, minutes, seconds, nanoseconds int) time.Time {
	y, m, d := t.Date()
	h, x, s := t.Clock()
	n := t.Nanosecond()
	return time.Date(y, m, d, h+hours, x+minutes, s+seconds, n+nanoseconds, t.Location())
}

const (
	timeHour  = 24
	timeClock = 60
	timeNano  = 1_000_000_000

	Day              = 24 * time.Hour
	DateTimeZone     = "2006-01-02 15:04:05 MST"           // 带时区的普通时间格式
	DateTimeNanoZone = "2006-01-02 15:04:05.999999999 MST" // 带时区和纳秒的普通时间格式
)

// Days to convert duration to days.
//
// 将持续时间转换为天数。
func Days(duration time.Duration) float64 {
	return duration.Hours() / timeHour
}

// AddDays to add days to time. Not change original time, return new time.
//
// 添加天数到时间，该函数不会改变原始时间，而是返回一个新的时间。会有时间误差。
func AddDays(t time.Time, days float64) time.Time {
	D := math.Floor(days) // day, integer part of time
	F := days - D         // clock, fractional part of time

	y, m, d := t.Date()
	h, x, s := t.Clock()
	n := t.Nanosecond()

	// 小数转为时分秒
	F *= timeHour
	hours := int(math.Floor(F))

	F -= float64(hours)
	F *= timeClock
	minutes := int(math.Floor(F))

	F -= float64(minutes)
	F *= timeClock
	seconds := int(F)

	F -= float64(seconds)
	F *= timeNano
	nano := int(F)

	return time.Date(y, m, d+int(D), h+hours, x+minutes, s+seconds, n+nano, t.Location())
}

func TimeCompare(a, b *time.Time) int {
	if a == nil {
		if b == nil {
			return 0
		}
		return 1
	}
	if b == nil {
		return -1
	}
	if a.Equal(*b) {
		return 0
	}
	if a.Before(*b) {
		return -1
	} else if a.After(*b) {
		return 1
	}
	return 0
}
