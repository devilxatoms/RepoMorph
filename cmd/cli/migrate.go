package cli

import (
	"fmt"

	"devilxstudios.com/repomorph/internal/config"
	"devilxstudios.com/repomorph/internal/models"
	"github.com/spf13/cobra"
)

var migrationOptions models.MigrationOptions

// migrateCmd captures user input and calls `migrator.ExecuteMigration()`
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate a repository from one provider to another",
	Long:  "This command facilitates the migration of repositories, including their history, users, and configurations, from one Git provider to another.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔄 Starting repository migration...")

		fmt.Println("migrationOptions", &migrationOptions)

		// Load configuration from file or CLI flags
		migrationConfig, err := config.LoadConfig(&migrationOptions)
		if err != nil {
			fmt.Println("❌ Error loading configuration file:", err)
			return
		}

		fmt.Println("migrationOptions", &migrationOptions)

		fmt.Println("validating migrationConfig", migrationConfig)

		// // Select appropriate migrator based on the source provider
		// m := migrator.SelectMigrator(migrationConfig)
		// if m == nil {
		// 	fmt.Println("❌ Unsupported provider:", migrationConfig.SourceProvider)
		// 	return
		// }

		// // Execute the migration process
		// fmt.Println("🚀 Executing migration...")
		// migrator.ExecuteMigration(m, migrationConfig)
		// fmt.Println("✅ Migration completed successfully!")
	},
}

// Initialize command-line flags for the `migrate` command
// These flags allow users to configure their migration directly from the command line
func init() {
	fmt.Println("🔧 Initializing migrate command...")
	rootCmd.AddCommand(migrateCmd)

	migrateCmd.Flags().StringVarP(&migrationOptions.SourceProvider, "source", "s", "", "Source provider (github/gitlab/bitbucket)")
	migrateCmd.Flags().StringVarP(&migrationOptions.TargetProvider, "target", "t", "", "Target provider (github/gitlab/bitbucket)")
	migrateCmd.Flags().StringVarP(&migrationOptions.SourceURL, "source-url", "", "", "Source repository URL")
	migrateCmd.Flags().StringVarP(&migrationOptions.TargetURL, "target-url", "", "", "Target repository URL")
	migrateCmd.Flags().StringVarP(&migrationOptions.SourceToken, "source-token", "", "", "Source provider token")
	migrateCmd.Flags().StringVarP(&migrationOptions.TargetToken, "target-token", "", "", "Target provider token")
	migrateCmd.Flags().StringVarP(&migrationOptions.ConfigFile, "config", "c", "", "Configuration file for bulk migrations")
	migrateCmd.Flags().BoolVarP(&migrationOptions.KeepAllBranchesHistory, "keep-branches", "k", false, "Keep all branches history")
	migrateCmd.Flags().StringVarP(&migrationOptions.DesiredGitWorkflow, "workflow", "w", "", "Desired Git workflow")
	migrateCmd.Flags().BoolVarP(&migrationOptions.DeleteSourceRepo, "delete-source", "d", false, "Delete source repository after migration")
	migrateCmd.Flags().BoolVarP(&migrationOptions.CarryUsers, "carry-users", "u", false, "Carry users and permissions")
	migrateCmd.Flags().BoolVarP(&migrationOptions.SyncTeams, "sync-teams", "", false, "Sync teams between providers")
	migrateCmd.Flags().BoolVarP(&migrationOptions.InviteMissingUsers, "invite-missing-users", "", false, "Invite missing users to the target repository")
	migrateCmd.Flags().BoolVarP(&migrationOptions.TransferIssues, "transfer-issues", "", false, "Transfer repository issues")
	migrateCmd.Flags().BoolVarP(&migrationOptions.TransferWiki, "transfer-wiki", "", false, "Transfer repository wiki")
	migrateCmd.Flags().BoolVarP(&migrationOptions.TransferPullRequests, "transfer-prs", "", false, "Transfer open pull requests")
	migrateCmd.Flags().BoolVarP(&migrationOptions.TransferCiCdPipelines, "transfer-cicd", "", false, "Transfer CI/CD pipeline configurations")
}
