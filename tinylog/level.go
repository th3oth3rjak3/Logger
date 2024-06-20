package tinylog

// Level is used to define the severity of a log message.
type Level byte

const (
	// LevelDebug represents the lowest severity level, mostly for debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents log information that is deemed valuable.
	LevelInfo
	// LevelError represents the highest logging level, only to be used for errors.
	LevelError
)

// DisplayName converts the logging level into a three (3) character name that represents the log level.
func (l Level) DisplayName() string {
	switch l {
	case LevelDebug:
		return "DBG"
	case LevelInfo:
		return "INF"
	case LevelError:
		return "ERR"
	default:
		return ""
	}
}
