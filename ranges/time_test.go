package ranges

import (
	"testing"
	"time"
)

func TestTimeRange(t *testing.T) {
	now := time.Now()
	end := now.Add(24 * time.Hour)
	r := Time(&now, &end)
	e := now.Add(1 * time.Hour)
	if !In(r, &e) {
		t.Fatal("time range error")
	}
	e = now.Add(-1 * time.Second)
	if In(r, &e) {
		t.Fatal("time range error")
	}
	e = now.Add(25 * time.Hour)
	if In(r, &e) {
		t.Fatal("time range error")
	}
}
