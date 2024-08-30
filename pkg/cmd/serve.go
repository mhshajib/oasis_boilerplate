package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/mhshajib/oasis_boilerplate/pkg/config"
	"github.com/mhshajib/oasis_boilerplate/pkg/conn"
	"github.com/mhshajib/oasis_boilerplate/pkg/log"
	appMiddleware "github.com/mhshajib/oasis_boilerplate/pkg/middleware"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve run available servers such as: HTTP/JSON or gRPC",
		Long:  `Serve run available servers such as: HTTP/JSON or gRPC`,
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Info("Connecting database")
			if err := conn.ConnectMongoDB(); err != nil {
				log.Fatal(err)
			}
			log.Info("Database connected successfully!")

			log.Info("Connecting cache server")
			if err := conn.ConnectDefaultCache(); err != nil {
				log.Fatal(err)
			}
			log.Info("Cache server connected successfully!")
		},
		Run: serve,
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	// Initialize stop channel for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	// build and run http server
	httpSrv := buildHTTP(cmd, args)
	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("HTTP server failed: ", err)
		}
	}()

	// Build and run gRPC server
	grpcCfg := config.GRpcApp()
	grpcSrv := buildGRPC(cmd, args, grpcCfg)
	grpcLis, err := net.Listen("tcp", fmt.Sprintf(":%s", strconv.Itoa(grpcCfg.GRpcPort))) // Define grpcPort as needed
	if err != nil {
		log.Fatal("Failed to listen for gRPC: ", err)
	}
	go func() {
		if err := grpcSrv.Serve(grpcLis); err != nil {
			log.Fatal("gRPC server failed: ", err)
		}
	}()

	<-stop
	log.Info("Shutting down servers...")

	// Shutdown HTTP server
	httpCtx, httpCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer httpCancel()
	if err := httpSrv.Shutdown(httpCtx); err != nil && err != http.ErrServerClosed {
		log.Fatal("Failed to shutdown HTTP server: ", err)
	}

	// Shutdown gRPC server
	grpcSrv.GracefulStop()
	log.Info("Server shutdown successful!")
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if len(path) < 1 || path[0] != '/' {
		panic("path must start with '/' in FileServer")
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

// buildHTTP register available handlers and return a http server
func buildHTTP(cmd *cobra.Command, args []string) *http.Server {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(appMiddleware.NewStructuredLogger(log.DefaultLogger()))
	// r.Use(appMiddleware.Auth)

	r.Handle("/asset/*", http.StripPrefix("/asset", http.FileServer(http.Dir("./asset"))))

	httpCfg := config.HttpApp()
	db := conn.MongoDB()
	fmt.Println(db)

	cacher := conn.DefaultCache()
	fmt.Println(cacher)

	//Repositories

	//Usecases

	//Delivery

	httpPort := fmt.Sprintf(":%d", httpCfg.HTTPPort)
	log.Info("HTTP Listening on port", httpPort)
	return &http.Server{
		Addr:              httpPort,
		Handler:           r,
		ReadHeaderTimeout: httpCfg.ReadTimeout,
		WriteTimeout:      httpCfg.WriteTimeout,
		IdleTimeout:       httpCfg.IdleTimeout,
	}
}

func buildGRPC(cmd *cobra.Command, args []string, grpcCfg config.GRpcApplciation) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.MaxConcurrentStreams(grpcCfg.MaxConcurrentStreams),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: grpcCfg.MaxConnectionIdle * time.Minute, // Similar to IdleTimeout
			Time:              grpcCfg.Time * time.Hour,                // Frequency of pings sent to clients
			Timeout:           grpcCfg.Timeout * time.Second,
		}),
		grpc.MaxRecvMsgSize(1024*1024*grpcCfg.MaxRecvMsgSize), // 4 MB
		grpc.MaxSendMsgSize(1024*1024*grpcCfg.MaxSendMsgSize), // 4 MB
	)

	// db := conn.MongoDB()

	// pb.RegisterStorageServiceServer(grpcServer, fileStorageServer)
	log.Info("gRpc Listening on port", fmt.Sprintf(":%s", strconv.Itoa(grpcCfg.GRpcPort)))
	return grpcServer
}
