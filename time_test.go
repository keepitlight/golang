package golang

import (
	"testing"
	"time"
)

func TestTimeRange(t *testing.T) {
	now := time.Now()
	end := now.Add(24 * time.Hour)
	r := TimeRange(&now, &end)
	if !r.In(now.Add(1 * time.Hour)) {
		t.Fatal("time range error")
	}
	if r.In(now.Add(-1 * time.Second)) {
		t.Fatal("time range error")
	}
	if r.In(now.Add(25 * time.Hour)) {
		t.Fatal("time range error")
	}
}
