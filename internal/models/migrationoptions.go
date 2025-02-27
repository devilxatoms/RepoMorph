package models

// MigrationOptions represents user input, either from CLI flags or a configuration file.
type MigrationOptions struct {
	ConfigFile             string // Path to a configuration file (JSON, YAML, Excel) for bulk migration
	SourceProvider         string // Source provider (e.g., GitHub, GitLab, Bitbucket)
	TargetProvider         string // Target provider (e.g., GitHub, GitLab, Bitbucket)
	SourceURL              string // URL of the source repository
	TargetURL              string // URL of the target repository
	SourceToken            string // Authentication token for the source provider
	TargetToken            string // Authentication token for the target provider
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
