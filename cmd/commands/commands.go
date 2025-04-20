package commands

import (
	"os"

	"github.com/dev-saw99/ioscript-cli/pkg/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ioscript",
	Short: "ioscript - Interactive Go programming learning tool",
	Long: `ioscript is a CLI tool that provides a structured, interactive learning experience 
for users learning Go programming. It acts as both a tutorial and code execution environment,
helping users track progress through lessons and validate their code.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is provided, show help
		cmd.Help()
	},
}

func init() {
	registerCmds()
}

func registerCmds() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(startCmd)
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error("Command execution failed", "error", err)
		os.Exit(1)
	}
}
