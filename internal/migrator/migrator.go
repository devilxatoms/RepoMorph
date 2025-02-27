package migrator

import (
	"fmt"
	"log"
)

// Migrator define la interfaz para la migración de repositorios
// Cualquier proveedor (GitHub, GitLab, Bitbucket) debe implementarla
type Migrator interface {
	Authenticate() error
	CloneRepository() error
	CreateRepository() error
	PushRepository() error
	SetPermissions() error
	CarryUsers() error // Agregado para la migración de usuarios y permisos
}

// ExecuteMigration ejecuta la migración con opciones avanzadas si están definidas
func ExecuteMigration(m Migrator, config *MigrationConfig) error {
	log.Println("🚀 Starting migration...")

	if err := m.Authenticate(); err != nil {
		return fmt.Errorf("❌ Authentication failed: %w", err)
	}

	if err := m.CloneRepository(); err != nil {
		return fmt.Errorf("❌ Clone failed: %w", err)
	}

	if err := m.CreateRepository(); err != nil {
		return fmt.Errorf("❌ Repo creation failed: %w", err)
	}

	if err := m.PushRepository(); err != nil {
		return fmt.Errorf("❌ Push failed: %w", err)
	}

	if err := m.SetPermissions(); err != nil {
		return fmt.Errorf("❌ Permissions setup failed: %w", err)
	}

	if config != nil && config.CarryUsers {
		if err := m.CarryUsers(); err != nil {
			return fmt.Errorf("❌ User migration failed: %w", err)
		}
	}

	log.Println("✅ Migration completed successfully!")
	return nil
}

// MigrationConfig almacena las opciones avanzadas de migración
type MigrationConfig struct {
	CarryUsers             bool   `json:"carry_users" yaml:"carry_users"`
	DesiredGitWorkflow     string `json:"desired_git_workflow" yaml:"desired_git_workflow"`
	KeepAllBranchesHistory bool   `json:"keep_all_branches_history" yaml:"keep_all_branches_history"`
	DeleteSourceRepo       bool   `json:"delete_source_repo" yaml:"delete_source_repo"`
}
