package main

import "fmt"

func help() {
	fmt.Print(`
  Usage: gtt [options]

  Options:

  -v, --version       output version number
  -h, --help          output this information
  start               start day
  task <description>  add a task
  pause               pause day
  resume              resume day
  status              show day status


  Generating an invoice document:

  invoice --from=YYYY-MM-DD --to=YYYY-MM-DD --cost-per-hour=1.45

`)
}
