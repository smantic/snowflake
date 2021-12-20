package snowflake

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {

}

func TestUnpack(t *testing.T) {

	epoch := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)

	// one milliseconds since 2021/1/1
	x := (1 << 22) | 1

	ti, mid, sid := Unpack(int64(x), epoch)

	if exp := epoch.Add(1 * time.Millisecond); ti != exp {
		t.Errorf("expected %v, but got: %v", exp, ti)
	}

	if mid != 0 {
		t.Errorf("expected 0 machine ID, got: %d\n", mid)
	}

	if sid != 1 {
		t.Errorf("expected %v, but got %v\n", 1, sid)
	}
}
