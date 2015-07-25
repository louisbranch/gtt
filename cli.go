package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/luizbranco/gtt/tracker"
)

const (
	file    = ".gtt"
	version = "0.0.1"
)

func start(t *tracker.Tracker) error {
	d, err := t.NewDay()
	if err != nil {
		return err
	}
	t.SaveDay(d)
	return nil
}

func task(t *tracker.Tracker) error {
	if len(os.Args) < 3 {
		return errors.New("Task description is required")
	}
	d, err := t.Today()
	if err != nil {
		return err
	}
	err = d.Task(os.Args[2])
	if err != nil {
		return err
	}
	t.SaveDay(d)
	return nil
}

func pause(t *tracker.Tracker) error {
	d, err := t.Today()
	if err != nil {
		return err
	}
	err = d.Pause()
	if err != nil {
		return err
	}
	t.SaveDay(d)
	return nil
}

func resume(t *tracker.Tracker) error {
	d, err := t.Today()
	if err != nil {
		return err
	}
	err = d.Resume()
	if err != nil {
		return err
	}
	t.SaveDay(d)
	return nil
}

func status(t *tracker.Tracker) error {
	d, err := t.Today()
	if err != nil {
		return err
	}
	status := d.Status()
	fmt.Printf("[STATUS] %s\n", status)
	return nil
}

func run(t *tracker.Tracker) error {
	switch os.Args[1] {
	case "start":
		return start(t)
	case "task":
		return task(t)
	case "pause":
		return pause(t)
	case "resume":
		return resume(t)
	case "status":
		return status(t)
	default:
		return errors.New("Unknown command")
	}
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}
	t, err := tracker.New(file)
	if err != nil {
		log.Fatal(err)
	}
	err = run(t)
	if err == nil {
		err = t.Save()
	}
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err.Error())
		os.Exit(1)
	}
}
