package common

import "time"

type Command struct {
	Command string
	Timeout time.Duration
}
