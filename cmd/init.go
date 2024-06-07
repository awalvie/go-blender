package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new site",
	Run:   initRun,
}

func init() {
	// Setup init flags
	initCmd.Flags().StringP("path", "p", ".", "Path to initialize site")

	// Add init command to root command
	rootCmd.AddCommand(initCmd)
}

// initRun gets called when the init command in run in the cli
// It creates the directory structure for a new site in the specified path
func initRun(cmd *cobra.Command, args []string) {
	// Get path from the flag
	path := cmd.Flag("path").Value.String()

	// Check if the path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("Path %s does not exist", path)
	}

	// Create the directory structure
	var dirs = []string{
		"static",
		"templates",
		"index",
	}

	// Create the directories
	for _, dir := range dirs {
		err := os.MkdirAll(path+"/"+dir, 0755)
		if err != nil {
			log.Fatalf("Error creating directory %s: %v", dir, err)
		}
	}

	// Create the default config file
	configFile := path + "/config.yaml"
	configContent := `dirs:
  templates: ./templates
  static: ./static
  index: ./index`

	err := os.WriteFile(configFile, []byte(configContent), 0644)
	if err != nil {
		log.Fatalf("Error creating config file: %v", err)
	}

	log.Printf("Site initialized in %s", path)
}
