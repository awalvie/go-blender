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

func buildRun(cmd *cobra.Command, args []string) {
	fmt.Println("build called")
}
