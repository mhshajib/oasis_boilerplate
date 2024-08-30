package seeder

import (
	"github.com/mhshajib/oasis_boilerplate/pkg/seeder"
	"github.com/spf13/cobra"
)

// SeederCmd represents seeder's seed command
var SeederCmd = &cobra.Command{
	Use:   "seeder",
	Short: "seeder create/drop fake data in database",
	Long:  `seeder create/drop fake data in database`,
}

// keep in order so that user_id can be used by merchant, merchant_id can be used by offer etc
var seeders = []seeder.Seeder{
	// &seeder.UserSeeder{},
}
