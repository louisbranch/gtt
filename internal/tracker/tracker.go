package tracker

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

var Now = time.Now()

type Tracker struct {
	Name string
	Days map[string]Day
}

func New(filename string) (*Tracker, error) {
	t := &Tracker{Name: filename, Days: make(map[string]Day)}
	flags := os.O_RDONLY
	f, err := os.OpenFile(t.Name, flags, 0644)
	if os.IsNotExist(err) {
		return t, nil
	}
	defer f.Close()
	if err != nil {
		return t, err
	}
	dec := gob.NewDecoder(f)
	err = dec.Decode(t)
	if err != nil && err != io.EOF {
		return t, err
	}
	return t, nil
}

func (t *Tracker) Save() error {
	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	f, err := os.OpenFile(t.Name, flags, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	err = enc.Encode(t)
	return err
}

func (t *Tracker) ToJSON() error {
	f := os.Stdout
	defer f.Close()
	blob, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	_, err = f.Write(blob)
	return err
}

func (t *Tracker) NewDay() (Day, error) {
	key := date(Now)
	d, ok := t.Days[key]
	if ok {
		return d, fmt.Errorf("Day %s already exists", key)
	}
	return Day{Start: Now, Tasks: []Task{}, Pauses: []Pause{}}, nil
}

func (t *Tracker) Today() (Day, error) {
	key := date(Now)
	d, ok := t.Days[key]
	if !ok {
		return d, fmt.Errorf("Day %s hasn't been started yet", key)
	}
	return d, nil
}

func (t *Tracker) SaveDay(d Day) {
	key := date(d.Start)
	t.Days[key] = d
}

func date(t time.Time) string {
	return t.Format("2006-01-02")
}
