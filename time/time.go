package time

import (
	"fmt"
	"log/slog"
	"time"
)

// Measure to measure time
//
// 测试函数执行时间，示例：
//
//	func main() {
//	  fmt.Printf("elapsed: %s", time.Measure(func(){
//	    time.Sleep(time.Second)
//	  }))
//	  // output:
//	  // elapsed: 1.00032619s
//	}
func Measure(f func()) time.Duration {
	start := time.Now()
	f()

	return time.Since(start)
}

// ElapsedCallback represents a callback to measure elapsed time when done
//
// 测量执行时间的回调
type ElapsedCallback interface {
	// Done happens when elapsed time is done
	//
	// 将在测量完成时被调用
	Done(duration time.Duration)
}

type ElapsedFunc func(duration time.Duration)

func (f ElapsedFunc) Done(duration time.Duration) {
	f(duration)
}

// Elapsed to measure elapsed time
//
// 测量使用时间，例：
//
//	import (
//	  "time"
//
//	  t "github.com/keepitlight/golang/time"
//	)
//
//	func main() {
//	  defer t.Elapsed(t.Print)()
//	  time.Sleep(time.Second)
//	  // output:
//	  // elapsed: 1.00032619s
//	}
func Elapsed(f ElapsedCallback) (done func()) {
	start := time.Now()
	return func() {
		d := time.Since(start)
		f.Done(d)
	}
}

var (
	// Print implements ElapsedCallback to print elapsed time
	Print = ElapsedFunc(func(duration time.Duration) {
		fmt.Printf("elapsed: %s", duration)
	})
	// Logger implements ElapsedCallback to log elapsed time
	Logger = &Log{"elapsed", "duration"}
)

type Log struct {
	Msg string
	Key string
}

func (l *Log) Set(msg, key string) *Log {
	l.Msg = msg
	l.Key = key
	return l
}

func (l *Log) Done(duration time.Duration) {
	slog.Info(l.Msg, l.Key, duration)
	// log output:
	// 2025/12/30 23:00:38 INFO elapsed duration=1.001027916s
}
