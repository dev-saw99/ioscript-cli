package commands

import (
	"fmt"

	"github.com/dev-saw99/ioscript-cli/pkg/logger"
	"github.com/spf13/cobra"
)

// Version information
const (
	VersionNumber = "0.1"
	VersionName   = "FoxWrap"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version of ioscript CLI",
	Long:  `Shows the version number of the currently installed ioscript cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ioscript:v%s (%s)\n", VersionNumber, VersionName)
		logger.Info("Version command executed", "version", VersionNumber, "name", VersionName)
	},
}
