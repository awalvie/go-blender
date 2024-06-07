package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new site",
	Run:   initRun,
}

func initRun(cmd *cobra.Command, args []string) {
	fmt.Println("init called")
}
