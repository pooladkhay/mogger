package mogger

type LogLevel string

const (
	Info  LogLevel = "Info"
	Debug LogLevel = "Debug"
	Warn  LogLevel = "Warn"
	Fatal LogLevel = "Fatal"
)
