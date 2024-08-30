package conn

import (
	"context"
	"fmt"
	"time"

	"github.com/mhshajib/oasis_boilerplate/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoClient *mongo.Client

// ConnectMongoDB return a mongodb connection instance
func ConnectMongoDB() error {
	dbCfg := config.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	credential := options.Credential{
		Username: dbCfg.Username,
		Password: dbCfg.Password,
	}
	clientOpts := options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d", dbCfg.Host, dbCfg.Port)).
		SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return fmt.Errorf("conn: failed to connect database: %v", err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("conn: failed to connect database: %v", err)
	}
	mgoClient = client
	return nil
}

// MongoDB return mongodb databse
func MongoDB() *mongo.Database {
	return mgoClient.Database(config.DB().Name)
}
