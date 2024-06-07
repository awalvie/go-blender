package cmd

import (
	"log"
	"os"

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
	buildCmd.Flags().StringP("path", "p", ".", "Path to site source")
	buildCmd.Flags().StringP("output", "o", "build", "Path to site output")
	buildCmd.Flags().BoolP("create", "c", false, "Create output directory if it does not exist")

	// Add build command to root command
	rootCmd.AddCommand(buildCmd)
}

func buildRun(cmd *cobra.Command, args []string) {
	// Get values from flags
	path := cmd.Flag("path").Value.String()
	output := cmd.Flag("output").Value.String()
	create, err := cmd.Flags().GetBool("create")
	if err != nil {
		log.Fatalf("Error getting create flag: %v", err)
	}

	// Check if the path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("Path %s does not exist", path)
	}

	if create {
		// Create the output directory
		err := os.MkdirAll(output, 0755)
		if err != nil {
			log.Fatalf("Error creating output directory %s: %v", output, err)
		}
		log.Printf("Output directory %s created", output)
	} else {
		// Check if the output directory exists
		if _, err := os.Stat(output); os.IsNotExist(err) {
			log.Fatalf("Output directory %s does not exist", output)
		}
	}

	// Build the site:
	// 1. Read the markdown files in the index directory
	// 2. Render them using the templates in the templates directory
	// 3. Write the rendered HTML files to the output directory
	// 4. Copy the static files to the output directory
	log.Printf("Building site from %s to %s", path, output)
}
