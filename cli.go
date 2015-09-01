package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/luizbranco/gtt/internal/invoice"
	"github.com/luizbranco/gtt/internal/tracker"
)

const (
	file    = ".gtt"
	version = "1.0.0"
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
	if err := d.Task(os.Args[2]); err != nil {
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
	if err := d.Pause(); err != nil {
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
	if err := d.Resume(); err != nil {
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

func printInvoice(t *tracker.Tracker) error {
	i, err := invoice.New(t)
	if err != nil {
		return err
	}
	i.ToHTML()
	return nil
}

func toJson(t *tracker.Tracker) error {
	err := t.ToJSON()
	return err
}

func run(t *tracker.Tracker) error {
	switch os.Args[1] {
	case "-v", "--version":
		fmt.Printf("%s\n", version)
		return nil
	case "-h", "--help":
		help()
		return nil
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
	case "invoice":
		return printInvoice(t)
	case "json":
		return toJson(t)
	default:
		return errors.New("Unknown command")
	}
}

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	t, err := tracker.New(file)
	if err != nil {
		log.Fatal(err)
	}

	if err := run(t); err != nil {
		log.Fatal(err)
	}

	if err := t.Save(); err != nil {
		fmt.Printf("[ERROR] %s\n", err.Error())
		os.Exit(1)
	}
}
