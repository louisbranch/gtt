# gtt [![Build Status](https://drone.io/github.com/luizbranco/gtt/status.png)](https://drone.io/github.com/luizbranco/gtt/latest)

**TL/DR** If you code and need to create regular invoices of your working hours,
you might want to give gtt a try.

## What is gtt

**g**it **t**ime **t**racker is a simple program that allows you to track how
much time you spend coding per day and, at the end of a period, generate an
invoice. Even though it has git on its name, gtt can work with any version
control system that supports commit hooks or by directly issuing commands from
the terminal.

<!--more-->

## How it works

Assuming you have downloaded [gtt](http://github.com/luizbranco/gtt/releases)
and you have added it to your PATH, this is an example of a workflow:

```shell

# creates a .gtt file in the current directory
gtt init

# adds a hook to .git/hooks/commit-msg
gtt hook git

# starts the current day
gtt start

# ...

# tracks the first line of the commit message and when it was made
git commit -m "fix bug #42"

# ...

# pauses the tracking time
gtt pause

# ... grabbing lunch :)

# resumes the time
gtt resume

# tracks a task manually (that is what the commit hook uses under the hood)
gtt task "attending a very productive meeting"

# displays how much time was spent between the start of the day and the last task, excluding pauses
gtt status
> [STATUS] 2h35m

```

### Generating an invoice

gtt can create a simple invoice between two periods as a html page:

```shell

# outputs an html invoice to stdout, you can the write to a file
gtt invoice --from=2015-08-01 --to=2015-08-31 --cost-per-hour=1.00 > invoice.html

```
