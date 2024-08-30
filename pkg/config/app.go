package config

import (
	"time"

	"github.com/spf13/viper"
)

// HttpApplciation represents http applciation config
type HttpApplciation struct {
	HTTPPort int  `json:"http_port"`
	Verbose  bool `json:"debug"`

	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
	IdleTimeout  time.Duration `json:"idle_timeout"`

	PaginationLimit int `json:"pagination_limit"`
}

// GRpcApplciation represents grpc applciation config
type GRpcApplciation struct {
	GRpcPort int `json:"grpc_port"`

	MaxConcurrentStreams uint32        `json:"max_concurrent_Streams"`
	MaxConnectionIdle    time.Duration `json:"max_connection_idle"`
	Time                 time.Duration `json:"time"`
	Timeout              time.Duration `json:"timeout"`

	MaxRecvMsgSize int `json:"max_recv_msg_size"`
	MaxSendMsgSize int `json:"max_send_msg_size"`
}

var http_app HttpApplciation
var grpc_app GRpcApplciation

// HttpApp contains http app configurations
func HttpApp() HttpApplciation {
	return http_app
}

// GRpcApp contains grpc app configurations
func GRpcApp() GRpcApplciation {
	return grpc_app
}

func loadApp() {
	http_app = HttpApplciation{
		HTTPPort:        viper.GetInt("http_app.http_port"),
		Verbose:         viper.GetBool("http_app.verbose"),
		ReadTimeout:     viper.GetDuration("http_app.read_timeout") * time.Second,
		WriteTimeout:    viper.GetDuration("http_app.write_timeout") * time.Second,
		IdleTimeout:     viper.GetDuration("http_app.idle_timeout") * time.Second,
		PaginationLimit: viper.GetInt("http_app.pagination_limit"),
	}

	grpc_app = GRpcApplciation{
		GRpcPort:             viper.GetInt("grpc_app.grpc_port"),
		MaxConcurrentStreams: viper.GetUint32("grpc_app.max_concurrent_Streams"),

		MaxConnectionIdle: viper.GetDuration("grpc_app.max_connection_idle") * time.Second,
		Time:              viper.GetDuration("grpc_app.time") * time.Second,
		Timeout:           viper.GetDuration("grpc_app.timeout") * time.Second,

		MaxRecvMsgSize: viper.GetInt("grpc_app.max_recv_msg_size"),
		MaxSendMsgSize: viper.GetInt("grpc_app.max_send_msg_size"),
	}
}
