package cli

import (
	"fmt"
	"devilxstudios.com/repomorph/internal/migrator"
	"devilxstudios.com/repomorph/internal/providers/github"
	"devilxstudios.com/repomorph/internal/providers/gitlab"

	"github.com/spf13/cobra"
)

var (
	sourceProvider string
	targetProvider string
	sourceURL      string
	targetURL      string
	token          string
)

// migrateCmd captura el input y llama a `migrator.ExecuteMigration()`
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate a repository from one provider to another",
	Run: func(cmd *cobra.Command, args []string) {
		var m migrator.Migrator

		switch sourceProvider {
		case "github":
			m = &github.GitHubMigrator{SourceURL: sourceURL, TargetURL: targetURL, Token: token}
		case "gitlab":
			m = &gitlab.GitLabMigrator{SourceURL: sourceURL, TargetURL: targetURL, Token: token}
		default:
			fmt.Println("❌ Unsupported provider:", sourceProvider)
			return
		}

		// Ejecutar migración con la interfaz Migrator
		migrator.ExecuteMigration(m)
	},
}

// Inicializar los flags del comando `migrate`
func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().StringVarP(&sourceProvider, "source", "s", "", "Source provider (github/gitlab/bitbucket)")
	migrateCmd.Flags().StringVarP(&targetProvider, "target", "t", "", "Target provider (github/gitlab/bitbucket)")
	migrateCmd.Flags().StringVarP(&sourceURL, "source-url", "", "", "Source repository URL")
	migrateCmd.Flags().StringVarP(&targetURL, "target-url", "", "", "Target repository URL")
	migrateCmd.Flags().StringVarP(&token, "token", "", "", "Authentication token")
}
