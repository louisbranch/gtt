package invoice

import (
	"testing"
	"time"

	"github.com/louisbranch/gtt/internal/tracker"
)

func TestInvoice_SumDurations(t *testing.T) {
	now := time.Now()
	i := &Invoice{}
	i.Days = tracker.Days{
		{Start: now, Tasks: []tracker.Task{
			{End: now.Add(8 * time.Hour)},
		}},
		{Start: now, Tasks: []tracker.Task{
			{End: now.Add(4 * time.Hour)},
		}},
	}
	i.sumDurations()
	expected := 12 * time.Hour
	if expected != i.Duration {
		t.Errorf("Expecting duration to eq %d, was %d", expected, i.Duration)
	}
}

func TestInvoice_TotalCost(t *testing.T) {
	i := &Invoice{
		cost:     12.23,
		Duration: 12*time.Hour + 45*time.Minute,
	}
	expected := 155.9325
	result := i.TotalCost()
	if expected != result {
		t.Errorf("Expecting total cost to eq %.4f, was %.4f", expected, result)
	}
}

func TestIncoide_DurationFormated(t *testing.T) {
	i := &Invoice{
		Duration: 12*time.Hour + 45*time.Minute,
	}
	expected := "12h45m"
	result := i.DurationFormated()
	if expected != result {
		t.Errorf("Expecting duration to eq %s, was %s", expected, result)
	}
}
