package invoice

import (
	"bytes"
	"errors"
	"flag"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/louisbranch/gtt/internal/tracker"
)

const dateFormat = "2006-01-02"

// Invoice holds the information necessary to create the billing info
type Invoice struct {
	cost     float64
	From     time.Time
	To       time.Time
	Days     []tracker.Day
	Duration time.Duration
}

// New generates an invoice for the time period
func New(t *tracker.Tracker) (*Invoice, error) {
	buf := bytes.NewBufferString("Invoice parsing failed\n")
	f := flag.NewFlagSet("invoice", flag.ContinueOnError)
	f.SetOutput(buf)

	fromFlag := f.String("from", "", "start day, format: YYYY-MM-DD")
	toFlag := f.String("to", "", "end day, format: YYYY-MM-DD")
	cost := f.Float64("cost-per-hour", 0, "cost per hour")
	if err := f.Parse(os.Args[2:]); err != nil {
		return nil, err
	}

	from, errFrom := time.Parse(dateFormat, *fromFlag)
	to, errTo := time.Parse(dateFormat, *toFlag)
	if errFrom != nil || errTo != nil {
		f.PrintDefaults()
		return nil, errors.New(buf.String())
	}

	i := &Invoice{
		From: from,
		To:   to,
		cost: *cost,
	}
	i.filterDays(t.Days)
	i.sumDurations()

	return i, nil
}

func (i *Invoice) filterDays(days map[string]tracker.Day) {
	var filter tracker.Days
	for _, day := range days {
		if day.Start.After(i.From) && day.Start.Before(i.To.Add(24*time.Hour)) {
			filter = append(filter, day)
		}
	}
	sort.Sort(filter)
	i.Days = filter
}

func (i *Invoice) sumDurations() {
	var dur time.Duration
	for _, d := range i.Days {
		dur += d.Duration()
	}
	i.Duration = dur
}

// TotalCost returns the number of hours multiplied by the cost per hour
func (i *Invoice) TotalCost() float64 {
	return i.Duration.Hours() * float64(i.cost)
}

// DurationFormated formats the total duration as hh:mm
func (i *Invoice) DurationFormated() string {
	h := strconv.Itoa(int(i.Duration.Hours()))
	m := strconv.Itoa(int(i.Duration.Minutes()) % 60)
	return h + "h" + m + "m"
}
