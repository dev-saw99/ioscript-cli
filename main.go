package main

import (
	"log/slog"

	"github.com/dev-saw99/ioscript-cli/cmd"
	"github.com/dev-saw99/ioscript-cli/pkg/logger"
)

func init() {
	// Initialize the logger
	loggerConfig := &logger.LoggerConfig{
		Level:       slog.LevelDebug,
		FileMode:    true,
		LogFileName: "ioscript.log",
	}
	logger.NewLogger(loggerConfig)
}

func main() {
	// Run the ioScript
	cmd.Run_ioScript()
}
