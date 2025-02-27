package migrator

import (
	"fmt"

	"devilxstudios.com/repomorph/internal/providers/github"
	"devilxstudios.com/repomorph/internal/providers/gitlab"
)

// SelectMigrator elige el migrador correcto según el proveedor de origen
type MigratorConfig struct {
	SourceProvider string
	SourceURL      string
	TargetURL      string
	SourceToken    string
	TargetToken    string
}

// SelectMigrator retorna una implementación de Migrator según el proveedor seleccionado
func SelectMigrator(config MigratorConfig) Migrator {
	switch config.SourceProvider {
	case "github":
		return &github.GitHubMigrator{
			SourceURL:   config.SourceURL,
			TargetURL:   config.TargetURL,
			SourceToken: config.SourceToken,
			TargetToken: config.TargetToken,
		}
	case "gitlab":
		return &gitlab.GitLabMigrator{
			SourceURL:   config.SourceURL,
			TargetURL:   config.TargetURL,
			SourceToken: config.SourceToken,
			TargetToken: config.TargetToken,
		}
	default:
		fmt.Println("❌ Unsupported provider:", config.SourceProvider)
		return nil
	}
}
