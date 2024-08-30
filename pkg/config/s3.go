package config

import (
	"time"

	"github.com/mhshajib/oasis_boilerplate/pkg/storage"
	"github.com/mhshajib/oasis_boilerplate/pkg/storage/s3"
	"github.com/spf13/viper"
)

var storageManager *storage.Manager
var SelectedStorageProvider string

func StorageManager() *storage.Manager {
	return storageManager
}

func loadStorageManager() {
	storageManager = storage.NewManager()
	SelectedStorageProvider = viper.GetString("storage.provider")

	s3Provider := s3.S3Provider{
		KeyId:                      viper.GetString("storage.s3.keyId"),
		KeySecret:                  viper.GetString("storage.s3.keySecret"),
		Region:                     viper.GetString("storage.s3.region"),
		Bucket:                     viper.GetString("storage.s3.bucket"),
		Timeout:                    viper.GetDuration("storage.s3.timeout") * time.Second,
		PresignedUrlExpirationMins: viper.GetDuration("storage.s3.presignedUrlExpirationMins") * time.Minute,
	}
	storageManager.RegisterProvider(SelectedStorageProvider, s3Provider)
}
