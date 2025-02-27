package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var mode string

// Root command for the CLI
var rootCmd = &cobra.Command{
	Use:   "repomorph",
	Short: "A CLI tool to migrate repositories between different Git providers",
	Long: `RepoMorph is a command-line tool designed to help developers migrate repositories
between different Git hosting services such as GitHub, GitLab, and Bitbucket.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Running in %s mode...\n", mode)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to RepoMorph CLI! Use --help to see available commands.")
	},
}

// Inicializar la CLI
func init() {
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "cli", "Execution mode: cli, ui, or api")
}

// Execute runs the CLI application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
