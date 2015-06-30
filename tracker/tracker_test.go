package tracker

import (
	"testing"
	"time"
)

func init() {
	Now = time.Unix(0, 0).UTC()
}

func TestDate(t *testing.T) {
	expecting := "1970-01-01"
	result := date(Now)
	if expecting != result {
		t.Errorf("expected:\n%s\ngot:\n%s\n", expecting, result)
	}
}
