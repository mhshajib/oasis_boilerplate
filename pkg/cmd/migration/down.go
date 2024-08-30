package migration

import (
	"context"
	"time"

	"github.com/mhshajib/oasis_boilerplate/pkg/conn"
	"github.com/mhshajib/oasis_boilerplate/pkg/log"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// downCmd represents root migration command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Down drop collection and indices",
	Long:  `Down drop collection and indices`,
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Connecting database")
		if err := conn.ConnectMongoDB(); err != nil {
			log.Fatal(err)
		}
		log.Info("Database connected successfully!")
	},
	Run: down,
}

func init() {
	MigrationCmd.AddCommand(downCmd)
}

func down(cmd *cobra.Command, args []string) {
	log.Info("Dropping collections and indices")
	ctx := context.Background()
	db := conn.MongoDB()
	defer db.Client().Disconnect(ctx)

	for _, m := range models {
		coll := db.Collection(m.Name())
		opts := options.DropIndexes().SetMaxTime(10 * time.Second)
		_, err := coll.Indexes().DropAll(ctx, opts)
		if err != nil || coll.Drop(ctx) != nil {
			log.ErrorWithFields("Failed to drop index", log.Fields{
				"event":      "migration_down",
				"collection": m.Name(),
				"error":      err.Error(),
			})
			return
		}

	}

	log.Info("Dropping collections and indices successful!")
}
