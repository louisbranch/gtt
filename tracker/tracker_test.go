package tracker

import (
	"reflect"
	"testing"
	"time"
)

func init() {
	Now = time.Unix(0, 0).UTC()
}

func TestTrackerNewDay(t *testing.T) {
	tck := &Tracker{Days: make(map[string]Day)}
	d, err := tck.NewDay()
	if err != nil {
		t.Errorf("expected new day to not fail")
	}
	if d.Start.IsZero() {
		t.Errorf("expected new day to have started")
	}
	if d.Tasks == nil {
		t.Errorf("expected new day to have tasks")
	}
	if d.Pauses == nil {
		t.Errorf("expected new day to have pauses")
	}
}

func TestTrackerNewDay_SameDay(t *testing.T) {
	tck := &Tracker{Days: make(map[string]Day)}
	d, _ := tck.NewDay()
	tck.SaveDay(d)
	_, err := tck.NewDay()
	if err == nil {
		t.Errorf("expected same day to fail")
	}
}

func TestTrackerToday(t *testing.T) {
	tck := &Tracker{Days: make(map[string]Day)}
	d, _ := tck.NewDay()
	tck.SaveDay(d)
	expecting := d
	result, _ := tck.Today()
	if !reflect.DeepEqual(expecting, result) {
		t.Errorf("expected:\n%v\ngot:\n%v\n", expecting, result)
	}
}

func TestTrackerToday_NotCreated(t *testing.T) {
	tck := &Tracker{Days: make(map[string]Day)}
	_, err := tck.Today()
	if err == nil {
		t.Errorf("expected today to fail")
	}
}

func TestTrackerSaveDay(t *testing.T) {
	tck := &Tracker{Days: make(map[string]Day)}
	d, _ := tck.NewDay()
	tck.SaveDay(d)
	key := date(d.Start)
	expecting := d
	result := tck.Days[key]
	if !reflect.DeepEqual(expecting, result) {
		t.Errorf("expected:\n%v\ngot:\n%v\n", expecting, result)
	}
}

func TestDate(t *testing.T) {
	expecting := "1970-01-01"
	result := date(Now)
	if expecting != result {
		t.Errorf("expected:\n%s\ngot:\n%s\n", expecting, result)
	}
}
