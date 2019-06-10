package main

type Tasks map[string]Task

// Task is command which run by Schdule
type Task struct {
	Command   string     `yaml:"command"`
	Schedules []Schedule `yaml:"schedules"`
}

// Schedule is struct containing when the task should be run
type Schedule struct {
	Name      string   `yaml:"name"`
	Every     string   `yaml:"every"`
	Weekdays  []string `yaml:"weekdays"`
	Monthdays []int    `yaml:"monthdays"`
	At        []string `yaml:"at"`

	Except *Schedule `yaml:"except"`
}
