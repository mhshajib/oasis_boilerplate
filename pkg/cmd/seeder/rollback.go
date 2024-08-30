package seeder

import (
	"context"

	"github.com/mhshajib/oasis_boilerplate/pkg/conn"
	"github.com/mhshajib/oasis_boilerplate/pkg/log"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

// seedRollbackCmd represents rollback seeder command
var seedRollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "rollback remove fake data from database",
	Long:  `rollback remove fake data from database`,
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Connecting database")
		if err := conn.ConnectMongoDB(); err != nil {
			log.Fatal(err)
		}
		log.Info("Database connected successfully!")
	},
	Run: seedRollback,
}

func init() {
	SeederCmd.AddCommand(seedRollbackCmd)
}

func seedRollback(cmd *cobra.Command, args []string) {
	log.Info("Removing fake data")
	ctx := context.Background()
	db := conn.MongoDB()
	defer db.Client().Disconnect(ctx)

	for _, s := range seeders {
		log.Info("Truncating: ", s.Name())
		coll := db.Collection(s.Name())
		_, err := coll.DeleteMany(ctx, bson.M{})
		if err != nil {
			log.ErrorWithFields("Failed to remove fake data", log.Fields{
				"event":      "seed",
				"collection": s.Name(),
				"error":      err.Error(),
			})
			return
		}
		log.Info("Truncate success: ", s.Name())
	}

	log.Info("Fake data removed successfully!")
}
