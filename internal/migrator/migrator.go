package migrator

import "fmt"

// Migrator define la interfaz para la migración de repositorios
// Cualquier proveedor (GitHub, GitLab, Bitbucket) debe implementarla

type Migrator interface {
	Authenticate() error
	CloneRepository() error
	CreateRepository() error
	PushRepository() error
	SetPermissions() error
}

// ExecuteMigration ejecuta la migración con la implementación del proveedor seleccionado
func ExecuteMigration(m Migrator) {
	fmt.Println("\n🚀 Starting migration...")

	if err := m.Authenticate(); err != nil {
		fmt.Println("❌ Authentication failed:", err)
		return
	}

	if err := m.CloneRepository(); err != nil {
		fmt.Println("❌ Clone failed:", err)
		return
	}

	if err := m.CreateRepository(); err != nil {
		fmt.Println("❌ Repo creation failed:", err)
		return
	}

	if err := m.PushRepository(); err != nil {
		fmt.Println("❌ Push failed:", err)
		return
	}

	if err := m.SetPermissions(); err != nil {
		fmt.Println("❌ Permissions setup failed:", err)
		return
	}

	fmt.Println("✅ Migration completed successfully!")
}
