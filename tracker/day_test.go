package tracker

import (
	"reflect"
	"testing"
	"time"
)

func TestDayPaused_NoPauses(t *testing.T) {
	d := Day{Pauses: []Pause{}}
	expecting := false
	result := d.paused()
	if expecting != result {
		t.Errorf("expected:\n%t\ngot:\n%t\n", expecting, result)
	}
}

func TestDayPaused_PauseNotResumed(t *testing.T) {
	d := Day{Pauses: []Pause{}}
	d.Pause()
	expecting := true
	result := d.paused()
	if expecting != result {
		t.Errorf("expected:\n%t\ngot:\n%t\n", expecting, result)
	}
}

func TestDayPaused_PauseAndResumed(t *testing.T) {
	d := Day{Pauses: []Pause{}}
	d.Pause()
	d.Resume()
	expecting := false
	result := d.paused()
	if expecting != result {
		t.Errorf("expected:\n%t\ngot:\n%t\n", expecting, result)
	}
}

func TestDayTask_Unpaused(t *testing.T) {
	d := Day{Tasks: []Task{}}
	expecting := []Task{Task{End: Now, Description: "test"}}
	err := d.Task("test")
	if err != nil {
		t.Error(err)
	}
	result := d.Tasks
	if !reflect.DeepEqual(expecting, result) {
		t.Errorf("expected:\n%v\ngot:\n%v\n", expecting, result)
	}
}

func TestDayTask_Paused(t *testing.T) {
	d := Day{Tasks: []Task{}}
	d.Pause()
	err := d.Task("test")
	if err == nil {
		t.Error("Expected day task on paused day to fail")
	}
}

func TestDayPause_Unpaused(t *testing.T) {
	d := Day{Pauses: []Pause{}}
	d.Pause()
	expecting := []Pause{Pause{Start: Now}}
	result := d.Pauses
	if !reflect.DeepEqual(expecting, result) {
		t.Errorf("expected:\n%v\ngot:\n%v\n", expecting, result)
	}
}

func TestDayPause_Paused(t *testing.T) {
	d := Day{Pauses: []Pause{}}
	d.Pause()
	err := d.Pause()
	if err == nil {
		t.Error("Expected day pause on paused day to fail")
	}
}

func TestDayPause_Resumed(t *testing.T) {
	d := Day{Pauses: []Pause{}}
	d.Pause()
	d.Resume()
	err := d.Pause()
	if err != nil {
		t.Error("Expected day resumed to not fail on pause")
	}
}

func TestDayResume_Paused(t *testing.T) {
	d := Day{Pauses: []Pause{}}
	d.Pause()
	d.Resume()
	expecting := []Pause{Pause{Start: Now, End: Now}}
	result := d.Pauses
	if !reflect.DeepEqual(expecting, result) {
		t.Errorf("expected:\n%v\ngot:\n%v\n", expecting, result)
	}
}

func TestDayResume_Unpaused(t *testing.T) {
	d := Day{Pauses: []Pause{}}
	err := d.Resume()
	if err == nil {
		t.Error("Expected day resume on unpaused day to fail")
	}
}

func TestDayResume_Resumed(t *testing.T) {
	d := Day{Pauses: []Pause{}}
	d.Pause()
	d.Resume()
	err := d.Resume()
	if err == nil {
		t.Error("Expected day resume on unpaused day to fail")
	}
}

func TestDayStats_DayEmpty(t *testing.T) {
	d := Day{}
	result := d.Status()
	expecting := "0h0m"
	if expecting != result {
		t.Errorf("expected:\n%s\ngot:\n%s\n", expecting, result)
	}
}

func TestDayStats_DayComplete(t *testing.T) {
	d := Day{
		Start: Now,
		Tasks: []Task{
			Task{End: Now.Add(2 * time.Hour)},
			Task{End: Now.Add(5 * time.Hour)},
		},
		Pauses: []Pause{
			Pause{Start: Now.Add(1 * time.Hour), End: Now.Add(1*time.Hour + 15*time.Minute)},
			Pause{Start: Now.Add(3 * time.Hour), End: Now.Add(3*time.Hour + 25*time.Minute)},
		},
	}
	result := d.Status()
	expecting := "4h20m"
	if expecting != result {
		t.Errorf("expected:\n%s\ngot:\n%s\n", expecting, result)
	}
}

func TestDayStats_DayNoTasks(t *testing.T) {
	d := Day{Start: Now, Tasks: []Task{}}
	result := d.Status()
	expecting := "0h0m"
	if expecting != result {
		t.Errorf("expected:\n%s\ngot:\n%s\n", expecting, result)
	}
}

func TestDayStats_DayPaused(t *testing.T) {
	d := Day{
		Start: Now,
		Tasks: []Task{
			Task{End: Now.Add(2 * time.Hour)},
			Task{End: Now.Add(5 * time.Hour)},
		},
		Pauses: []Pause{
			Pause{Start: Now.Add(1 * time.Hour), End: Now.Add(1*time.Hour + 15*time.Minute)},
			Pause{Start: Now.Add(6 * time.Hour)},
		},
	}
	result := d.Status()
	expecting := "4h45m"
	if expecting != result {
		t.Errorf("expected:\n%s\ngot:\n%s\n", expecting, result)
	}
}
