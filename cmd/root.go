package cmd

import (
	"os"

	"github.com/dhftah/srv/internal/server"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "srv",
	Short: "Serve the directory",
	Long: `Serve the directory. For example, 

	# Serve the current directory locally, on port 8000
	$ srv
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		srv, err := server.New()

		if err != nil {
			return err
		}

		if err := srv.ListenAndServe(); err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
