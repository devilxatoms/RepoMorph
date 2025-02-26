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

// Definir el comando `migrate`
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate a repository from one provider to another",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Starting migration...")
		fmt.Printf("🔹 Source: %s (%s)\n", sourceURL, sourceProvider)
		fmt.Printf("🔹 Target: %s (%s)\n", targetURL, targetProvider)
		fmt.Println("✅ Migration completed!")
	},
}

var (
	sourceProvider string
	targetProvider string
	sourceURL      string
	targetURL      string
	token          string
)

// Inicializar los comandos y flags
func init() {
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "cli", "Execution mode: cli, ui, or api")
	rootCmd.AddCommand(migrateCmd)

	migrateCmd.Flags().StringVarP(&sourceProvider, "source", "s", "", "Source provider (github/gitlab/bitbucket)")
	migrateCmd.Flags().StringVarP(&targetProvider, "target", "t", "", "Target provider (github/gitlab/bitbucket)")
	migrateCmd.Flags().StringVarP(&sourceURL, "source-url", "", "", "Source repository URL")
	migrateCmd.Flags().StringVarP(&targetURL, "target-url", "", "", "Target repository URL")
	migrateCmd.Flags().StringVarP(&token, "token", "", "", "Authentication token")
}

// Execute runs the CLI application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
