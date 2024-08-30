package seeder

import (
	"context"

	"github.com/mhshajib/oasis_boilerplate/pkg/conn"
	"github.com/mhshajib/oasis_boilerplate/pkg/log"
	"github.com/spf13/cobra"
)

// seedCmd represents root seeder command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed create fake data to database",
	Long:  `seed create fake data to database`,
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Connecting database")
		if err := conn.ConnectMongoDB(); err != nil {
			log.Fatal(err)
		}
		log.Info("Database connected successfully!")
	},
	Run: seed,
}

func init() {
	SeederCmd.AddCommand(seedCmd)
}

func seed(cmd *cobra.Command, args []string) {
	log.Info("Starting seeder")
	ctx := context.Background()
	db := conn.MongoDB()
	defer db.Client().Disconnect(ctx)

	for _, s := range seeders {
		log.Info("Seeding: ", s.Name())
		err := s.Seed(ctx, db)
		if err != nil {
			log.ErrorWithFields("Failed to create fake data", log.Fields{
				"event":      "seed",
				"collection": s.Name(),
				"error":      err.Error(),
			})
			return
		}
		log.Info("Seeding success: ", s.Name())
	}

	log.Info("Seeding successful!")
}
