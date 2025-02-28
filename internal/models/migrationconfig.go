package models

type MigrationConfig struct {
	SourceProvider         string `json:"source_provider" yaml:"source_provider"`
	TargetProvider         string `json:"target_provider" yaml:"target_provider"`
	SourceURL              string `json:"source_url" yaml:"source_url"`
	TargetRepoName         string `json:"target_repo_name" yaml:"target_repo_name"`
	TargetURL              string `json:"target_url" yaml:"target_url"`
	SourceToken            string `json:"source_token" yaml:"source_token"`
	TargetToken            string `json:"target_token" yaml:"target_token"`
	KeepAllBranchesHistory bool   `json:"keep_all_branches_history" yaml:"keep_all_branches_history"`
	DesiredGitWorkflow     string `json:"desired_git_workflow" yaml:"desired_git_workflow"`
	DeleteSourceRepo       bool   `json:"delete_source_repo" yaml:"delete_source_repo"`
	CarryUsers             bool   `json:"carry_users" yaml:"carry_users"`
	SyncTeams              bool   `json:"sync_teams" yaml:"sync_teams"`
	InviteMissingUsers     bool   `json:"invite_missing_users" yaml:"invite_missing_users"`
	TransferIssues         bool   `json:"transfer_issues" yaml:"transfer_issues"`
	TransferWiki           bool   `json:"transfer_wiki" yaml:"transfer_wiki"`
	TransferPullRequests   bool   `json:"transfer_pull_requests" yaml:"transfer_pull_requests"`
	TransferCiCdPipelines  bool   `json:"transfer_ci_cd_pipelines" yaml:"transfer_ci_cd_pipelines"`
}

// BulkMigrationConfig represents multiple migrations
type BulkMigrationConfig struct {
	Repositories []MigrationConfig `json:"repositories" yaml:"repositories"`
}
