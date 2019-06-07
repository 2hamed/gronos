package main

type Task struct {
	Command   string
	Schedules map[string]Schedules `yaml:"schedules"`
}

type Schedules struct {
	every     string   `yaml:"every"`
	weekdays  []string `yaml:"weekdays"`
	monthdays []int8   `yaml:"monthdays"`
	at        []int8   `yaml:"at"`

	except *Schedules `yaml:"except"`
}
