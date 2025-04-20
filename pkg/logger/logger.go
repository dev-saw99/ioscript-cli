package logger

import (
	"io"
	"log"
	"log/slog"
	"os"
	"time"
)

// LoggerConfig holds the configuration for the logger
// It includes the logging level, file mode, and log file name
type LoggerConfig struct {

	// Logging level (e.g., Debug, Info, Warn, Error)
	Level slog.Level

	// FileMode indicates whether to log to a file
	// If true, LogFileName must be provided
	FileMode bool

	// LogFileName is the name of the log file
	// This is only used if FileMode is true
	// If FileMode is false, this field is ignored
	// and can be empty
	LogFileName string
}

var logger *slog.Logger

// NewLogger initializes a new logger with the specified configuration
// params:
// - config: a pointer to LoggerConfig struct
//   - level: the logging level (e.g., Debug, Info, Warn, Error)
//   - fileMode: whether to log to a file
//   - logFileName: the name of the log file (if fileMode is true)
func NewLogger(config *LoggerConfig) {
	// Create a new logger with default option
	var output io.Writer = os.Stdout

	if config.FileMode && config.LogFileName != "" {
		// Open the log file
		file, err := os.OpenFile(config.LogFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("failed to open log file: %v", err)
		} else {
			output = file
		}
	}

	// Custom time formatter that only includes up to seconds
	timeFormatter := func(t time.Time) string {
		return t.Format("2006-01-02T15:04:05Z")
	}

	opts := slog.HandlerOptions{
		Level: config.Level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Format time to seconds only
			if a.Key == "time" {
				if t, ok := a.Value.Any().(time.Time); ok {
					return slog.String("time", timeFormatter(t))
				}
			}
			return a
		},
	}

	// Create the logger with our custom handler
	handler := slog.NewJSONHandler(output, &opts)
	logger = slog.New(handler)
}

func Info(msg string, args ...any) {
	if logger != nil {
		logger.Info(msg, args...)
	}
}

func Debug(msg string, args ...any) {
	if logger != nil {
		logger.Debug(msg, args...)
	}
}

func Error(msg string, args ...any) {
	if logger != nil {
		logger.Error(msg, args...)
	}
}

func Warn(msg string, args ...any) {
	if logger != nil {
		logger.Warn(msg, args...)
	}
}

func Panic(msg string, args ...any) {
	if logger != nil {
		logger.Error(msg, args...)
		panic(msg)
	}
}

func Fatal(msg string) {
	if logger != nil {
		log.Fatalln(msg)
	}
}
