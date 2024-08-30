package migration

import (
	"github.com/mhshajib/oasis_boilerplate/pkg/migration"
	"github.com/spf13/cobra"
)

// MigrationCmd represents root migration command
var MigrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Migration create/drop collection and indices",
	Long:  `Migration create/drop collection and indices`,
}

var models = []migration.Model{
	// &migration.User{},
}
