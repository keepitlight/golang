package ranges

import (
	"time"

	"github.com/keepitlight/golang"
)

// Since like time.Since to create a new time range from start to now, nil if start is nil or start is after now.
//
// 从 start 到现在创建一个时间范围，如果 start 为 nil 或 start 在现在之后，则返回 nil
func Since(start *time.Time) golang.Range[*time.Time] {
	if start == nil {
		return nil
	}
	now := time.Now()
	if start.After(now) {
		return nil
	}
	return Time(start, &now)
}

// Duration return duration of time range
//
// 返回时间范围的持续时长
func Duration(t golang.Range[*time.Time]) time.Duration {
	if t == nil {
		return 0
	}
	l, u := t.Bounds()
	if l == nil || u == nil {
		return 0
	}
	return u.Sub(*l)
}

type timeInterval struct {
	start, end, initial, current *time.Time
	duration                     time.Duration
}

type TimeOption func(*timeInterval)

func Interval(r golang.Range[*time.Time], options ...TimeOption) golang.Interval[*time.Time] {
	start, end := r.Bounds()
	return Time(start, end, options...)
}

// Time creates a new time range, swap if the argument start is after the argument end.
//
// 创建一个时间范围，如果 start 在 end 之后，则交换 start 和 end
func Time(start, end *time.Time, options ...TimeOption) golang.Interval[*time.Time] {
	if start != nil && end != nil {
		if start.After(*end) {
			start, end = end, start
		}
	}
	initial := start
	t := &timeInterval{
		start:    start,
		end:      end,
		initial:  initial,
		current:  initial,
		duration: time.Hour, // 小时
	}
	for _, option := range options {
		option(t)
	}
	return t
}

// WithInitialValue set initial value of time range, default is start
//
// 设置时间范围的初始值，默认为 start
func WithInitialValue(v *time.Time) TimeOption {
	return func(interval *timeInterval) {
		if interval == nil || v == nil {
			return
		}
		if interval.start != nil && interval.start.After(*v) {
			return
		}
		if interval.end != nil && interval.end.Before(*v) {
			return
		}
		interval.initial = v
	}
}

// WithInterval set interval of time range, default is 1 hour
//
// 设置时间范围的间隔时长，默认为 1 小时
func WithInterval(duration time.Duration) TimeOption {
	return func(interval *timeInterval) {
		if interval == nil {
			return
		}
		interval.duration = duration
	}
}

func (t *timeInterval) Bounds() (lower, upper *time.Time) {
	return t.start, t.end
}

func (t *timeInterval) Comparer() func(a, b *time.Time) int {
	return golang.TimeCompare
}

func (t *timeInterval) Next() (current *time.Time, endOfInterval bool) {
	current = t.current
	if current == nil {
		if t.initial != nil {
			current = t.initial
		} else if t.start != nil {
			current = t.start
		}
	}
	if current == nil {
		endOfInterval = true
	} else if t.duration == 0 {
		endOfInterval = true
	} else if t.duration > 0 {
		n := current.Add(t.duration)
		if t.end != nil && (t.end.Before(*current) || t.end.Before(n)) {
			endOfInterval = true
		} else {
			t.current = &n
		}
	} else {
		n := current.Add(t.duration)
		if t.start != nil && (t.start.After(*current) || t.start.After(n)) {
			endOfInterval = true
		} else {
			t.current = &n
		}
	}
	return
}

func (t *timeInterval) Init() {
	if t.initial != nil {
		t.current = t.initial
	} else if t.start != nil {
		t.current = t.start
	}
}
