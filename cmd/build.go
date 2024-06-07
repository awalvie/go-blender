package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the site",
	Run:   buildRun,
}

func init() {
	// Setup build flags
	buildCmd.Flags().StringP("config", "c", "config.yaml", "Build config file")

	// Add build command to root command
	rootCmd.AddCommand(buildCmd)
}

func buildRun(cmd *cobra.Command, args []string) {
	fmt.Println("build called")
}
