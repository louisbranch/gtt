package main

import "fmt"

func help() {
	fmt.Println(`
  Usage: gtt [options]

  Options:

  -v, --version				output version number
  -h, --help          output this information
  start               start day
  task <description>  add a task
  pause               pause day
  resume              resume day
  status              show day status`)
}
