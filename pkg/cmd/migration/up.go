package migration

import (
	"context"
	"time"

	"github.com/mhshajib/oasis_boilerplate/pkg/conn"
	"github.com/mhshajib/oasis_boilerplate/pkg/log"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// upCmd represents root migration command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Up create collection and indices",
	Long:  `Up create collection and indices`,
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Connecting database")
		if err := conn.ConnectMongoDB(); err != nil {
			log.Fatal(err)
		}
		log.Info("Database connected successfully!")
	},
	Run: up,
}

func init() {
	MigrationCmd.AddCommand(upCmd)
}

func up(cmd *cobra.Command, args []string) {
	log.Info("Creating collections and indices")
	ctx := context.Background()
	db := conn.MongoDB()
	defer db.Client().Disconnect(ctx)

	for _, m := range models {
		coll := db.Collection(m.Name())
		opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
		_, err := coll.Indexes().CreateMany(ctx, m.Indices(), opts)
		if err != nil {
			log.ErrorWithFields("Failed to create index", log.Fields{
				"event":      "migration_down",
				"collection": m.Name(),
				"error":      err.Error(),
			})
			return
		}
	}

	log.Info("Creating collections and indices successful!")
}
