package config

import (
	"github.com/mhshajib/oasis_boilerplate/pkg/sms"
	"github.com/mhshajib/oasis_boilerplate/pkg/sms/bulkSmsBd"
	"github.com/spf13/viper"
)

var smsManager *sms.Manager
var SelectedProvider string
var SelectedSender string

func SmsManager() *sms.Manager {
	return smsManager
}

func loadSmsManager() {
	smsManager = sms.NewManager()
	SelectedProvider = viper.GetString("sms.provider")
	bulkSmsBdProvider := bulkSmsBd.BulkSmsBdProvider{
		APIKey:   viper.GetString("sms.bulk_sms_bd.api_key"),
		SenderID: viper.GetString("sms.bulk_sms_bd.sender_id"),
	}
	SelectedSender = viper.GetString("sms.bulk_sms_bd.sender_id")
	smsManager.RegisterProvider(SelectedProvider, bulkSmsBdProvider)
}
