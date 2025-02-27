package cli

import (
	"fmt"

	"devilxstudios.com/repomorph/internal/migrator"

	"github.com/spf13/cobra"
)

// MigrationOptions defines the structure for storing migration parameters
// These options are set via command-line flags and passed to the migration process
type MigrationOptions struct {
	SourceProvider         string // Source provider (e.g., GitHub, GitLab, Bitbucket)
	TargetProvider         string // Target provider (e.g., GitHub, GitLab, Bitbucket)
	SourceURL              string // URL of the source repository
	TargetURL              string // URL of the target repository
	SourceToken            string // Authentication token for the source provider
	TargetToken            string // Authentication token for the target provider
	ConfigFile             string // Configuration file for bulk migrations
	KeepAllBranchesHistory bool   // Whether to keep all branch history during migration
	DesiredGitWorkflow     string // The desired Git workflow for the migrated repository
	DeleteSourceRepo       bool   // Whether to delete the source repository after migration
	CarryUsers             bool   // Whether to migrate users and permissions
	SyncTeams              bool   // Whether to sync teams between providers
	InviteMissingUsers     bool   // Whether to invite missing users to the new repository
	TransferIssues         bool   // Whether to transfer issues from the source repository
	TransferWiki           bool   // Whether to transfer the repository wiki
	TransferPullRequests   bool   // Whether to transfer open pull requests
	TransferCiCdPipelines  bool   // Whether to transfer CI/CD pipeline configurations
}

var migrationOptions MigrationOptions

// migrateCmd captures user input and calls `migrator.ExecuteMigration()`
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate a repository from one provider to another",
	Long:  "This command facilitates the migration of repositories, including their history, users, and configurations, from one Git provider to another.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔄 Starting repository migration...")

		// Select appropriate migrator based on the source provider
		m := migrator.SelectMigrator(migrator.MigratorConfig{
			SourceProvider: migrationOptions.SourceProvider,
			SourceURL:      migrationOptions.SourceURL,
			TargetURL:      migrationOptions.TargetURL,
			SourceToken:    migrationOptions.SourceToken,
			TargetToken:    migrationOptions.TargetToken,
		})
		if m == nil {
			fmt.Println("❌ Unsupported provider:", migrationOptions.SourceProvider)
			return
		}

		// Define the migration configuration with advanced options
		config := &migrator.MigrationConfig{
			KeepAllBranchesHistory: migrationOptions.KeepAllBranchesHistory,
			DesiredGitWorkflow:     migrationOptions.DesiredGitWorkflow,
			DeleteSourceRepo:       migrationOptions.DeleteSourceRepo,
			CarryUsers:             migrationOptions.CarryUsers,
		}

		// Execute the migration process
		fmt.Println("🚀 Executing migration...")
		migrator.ExecuteMigration(m, config)
		fmt.Println("✅ Migration completed successfully!")
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
