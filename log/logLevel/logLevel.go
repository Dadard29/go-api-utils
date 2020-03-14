package logLevel

const (
	DEBUG   = 0
	INFO    = 1
	WARNING = 2
	ERROR   = 3
	FATAL   = 4
)

func LevelFromBool(verbose bool) int {
	level := INFO
	if verbose {
		level = DEBUG
	}
	return level
}
