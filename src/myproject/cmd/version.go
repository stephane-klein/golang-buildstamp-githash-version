package cmd

import (
	"fmt"
	"myproject/version"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"Version: %s\nUTC Build Time: %s\nGit Commit Hash: %s\n",
			version.Version,
			version.BuildStamp,
			version.GitHash,
		)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
