package time

import (
	"sync"
	"time"
)

// Timer represents a timer that can be stopped and resumed. If you need a fixed duration timer,
// please use the standard library time.Timer. It is concurrency safe.
//
// 能够停止和唤醒的分段计时器，如果要使用定长的计时器，请使用标准库的 time.Timer。
// 并发安全
type Timer struct {
	start    time.Time
	duration time.Duration
	paused   bool
	l        sync.Mutex
}

// Start a timer
//
// 启动计时器
func Start() *Timer {
	return &Timer{
		l:     sync.Mutex{},
		start: time.Now(),
	}
}

// Resume a timer if it is stopped, if it is running, do nothing
//
// 唤醒停止的计时器，如果计时器未停止，则什么也不做
func (t *Timer) Resume() {
	t.l.Lock()
	defer t.l.Unlock()

	if t.paused {
		t.start = time.Now()
		t.paused = false
	}
	return
}

// Stop a timer and return the duration since start, if the timer is already stopped, only return the elapsed time
//
// 停止计时器并返回时长，如果计时器已停止，则仅返回已计时的时长
func (t *Timer) Stop() time.Duration {
	t.l.Lock()
	defer t.l.Unlock()

	if !t.paused {
		t.duration += time.Since(t.start)
		t.paused = true
	}

	return t.duration
}
