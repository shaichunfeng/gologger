package formatter

import (
	"github.com/shaichunfeng/gologger/levels"
	"github.com/shaichunfeng/gologger/types"
)

// Formatter type format raw logging data into something useful
type Formatter interface {
	// Format formats the log event data into bytes
	Format(event *LogEvent) ([]byte, error)
}

// LogEvent is the respresentation of a single event to be logged.
type LogEvent struct {
	Message      string
	Level        levels.Level
	Metadata     map[string]string
	KeyValuePair []types.KeyValuePair
}
