package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"devilxstudios.com/repomorph/internal/models"
	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v3"
)

// LoadConfig reads a configuration file (if provided) and returns a BulkMigrationConfig.
func LoadConfig(options *models.MigrationOptions) (*models.BulkMigrationConfig, error) {
	// Si no hay archivo de configuración, usamos los flags del CLI
	if options.ConfigFile == "" {
		fmt.Println("🚧 No config file provided, using CLI flags only.")
		return &models.BulkMigrationConfig{
			Repositories: []models.MigrationConfig{
				NewMigrationConfig(options),
			},
		}, nil
	}

	data, err := os.ReadFile(options.ConfigFile)
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to read config file: %w", err)
	}

	ext := filepath.Ext(options.ConfigFile)
	var bulkConfig models.BulkMigrationConfig

	switch ext {
	case ".json":
		if err := json.Unmarshal(data, &bulkConfig); err != nil {
			return nil, fmt.Errorf("❌ Failed to decode JSON: %w", err)
		}
	case ".yaml", ".yml":
		fmt.Println("🔍 YAML file detected")
		if err := yaml.Unmarshal(data, &bulkConfig); err != nil {
			return nil, fmt.Errorf("❌ Failed to decode YAML: %w", err)
		}
	case ".xlsx":
		f, err := excelize.OpenFile(options.ConfigFile)
		if err != nil {
			return nil, fmt.Errorf("❌ Failed to open Excel file: %w", err)
		}
		defer f.Close()

		rows, err := f.GetRows("Sheet1")
		if err != nil {
			return nil, fmt.Errorf("❌ Failed to read Excel sheet: %w", err)
		}

		// Iterar sobre todas las filas (excepto el header) y agregarlas al bulkConfig
		for i := 1; i < len(rows); i++ {
			migration := models.MigrationConfig{
				SourceProvider:         rows[i][0],
				TargetProvider:         rows[i][1],
				SourceURL:              rows[i][2],
				TargetRepoName:         rows[i][3],
				TargetURL:              rows[i][4],
				SourceToken:            rows[i][5],
				TargetToken:            rows[i][6],
				KeepAllBranchesHistory: parseBool(rows[i][7]),
				DesiredGitWorkflow:     rows[i][8],
				DeleteSourceRepo:       parseBool(rows[i][9]),
				CarryUsers:             parseBool(rows[i][10]),
				SyncTeams:              parseBool(rows[i][11]),
				InviteMissingUsers:     parseBool(rows[i][12]),
				TransferIssues:         parseBool(rows[i][13]),
				TransferWiki:           parseBool(rows[i][14]),
				TransferPullRequests:   parseBool(rows[i][15]),
				TransferCiCdPipelines:  parseBool(rows[i][16]),
			}
			bulkConfig.Repositories = append(bulkConfig.Repositories, migration)
		}
	default:
		return nil, fmt.Errorf("❌ Unsupported configuration file format")
	}

	fmt.Println("✅ Config file loaded successfully.")
	return &bulkConfig, nil
}

// NewMigrationConfig converts a single MigrationOptions into a MigrationConfig
func NewMigrationConfig(options *models.MigrationOptions) models.MigrationConfig {
	return models.MigrationConfig{
		SourceProvider:         options.SourceProvider,
		TargetProvider:         options.TargetProvider,
		SourceURL:              options.SourceURL,
		TargetRepoName:         options.TargetRepoName,
		TargetURL:              options.TargetURL,
		SourceToken:            options.SourceToken,
		TargetToken:            options.TargetToken,
		KeepAllBranchesHistory: options.KeepAllBranchesHistory,
		DesiredGitWorkflow:     options.DesiredGitWorkflow,
		DeleteSourceRepo:       options.DeleteSourceRepo,
		CarryUsers:             options.CarryUsers,
		SyncTeams:              options.SyncTeams,
		InviteMissingUsers:     options.InviteMissingUsers,
		TransferIssues:         options.TransferIssues,
		TransferWiki:           options.TransferWiki,
		TransferPullRequests:   options.TransferPullRequests,
		TransferCiCdPipelines:  options.TransferCiCdPipelines,
	}
}

// parseBool converts Excel string values like "true", "1", "yes" to bool
func parseBool(value string) bool {
	v, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return v
}
