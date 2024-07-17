package golang

import (
	"math"
	"time"
)

// TimeRange create a new time range, nil if start or end is nil, swap if start is after end.
//
// 创建一个时间范围，如果 start 或 end 为 nil，则返回 nil，如果 start 在 end 之后，则交换 start 和 end
func TimeRange(start, end *time.Time) *Range[time.Time] {
	if start == nil || end == nil {
		return nil
	}
	if start.After(*end) {
		start, end = end, start
	}
	return &Range[time.Time]{
		Lower: *start,
		Upper: *end,
		comparer: func(a, b time.Time) int {
			if a.Equal(b) {
				return 0
			}
			if a.Before(b) {
				return -1
			}
			return 1
		},
	}
}

// Duration return duration of time range
//
// 返回时间范围的持续时长
func Duration(t *Range[time.Time]) time.Duration {
	if t == nil {
		return 0
	}
	return t.Upper.Sub(t.Lower)
}

// Since like time.Since to create a new time range from start to now, nil if start is nil or start is after now.
//
// 从 start 到现在创建一个时间范围，如果 start 为 nil 或 start 在现在之后，则返回 nil
func Since(start *time.Time) *Range[time.Time] {
	if start == nil {
		return nil
	}
	now := time.Now()
	if start.After(now) {
		return nil
	}
	return TimeRange(start, &now)
}

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

	Day = 24 * time.Hour
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
