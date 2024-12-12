// helper functions for main

package helper

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/connect"
	// "connectrpc.com/grpchealth"
	"connectrpc.com/otelconnect"
	validate "connectrpc.com/validate"
	protovalidate "github.com/bufbuild/protovalidate-go"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	redis "github.com/redis/go-redis/v9"
	zap "go.uber.org/zap"

	configs "github.com/ridehovr/rides/configs"
	pb "github.com/ridehovr/rides/gen/protos/api/rides/v1/ridesconnect"
	handler "github.com/ridehovr/rides/internal/v1/handler"
	"github.com/ridehovr/rides/internal/v1/repository"
	"github.com/ridehovr/rides/internal/v1/util/geoindex"
	sr "github.com/ridehovr/rides/internal/v1/util/streamregisters"
)

func newCORS() *cors.Cors {
	// To let web developers play with the demo service from browsers, we need a
	// very permissive CORS setup.
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: func(_ /* origin */ string) bool {
			// Allow all origins, which effectively disables CORS.
			return true
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
		},
		// Let browsers cache CORS information for longer, which reduces the number
		// of preflight requests. Any changes to ExposedHeaders won't take effect
		// until the cached data expires. FF caps this value at 24h, and modern
		// Chrome caps it at 2h.
		MaxAge: int(2 * time.Hour / time.Second),
	})
}

// register New gRPC service with all the conigs.
func RegisterServices(
	config *configs.Configurations,
	logger *zap.Logger,
	redis *redis.Client,
	validator *protovalidate.Validator,
	mux *http.ServeMux,
	interceptors connect.Option,
	repo repository.RidesInfoRepository,
	geo *geoindex.GeoIndex,
	streamRegister *sr.StreamRegister,
) *http.ServeMux {
	// register new auth servuce
	newDriverService, newLocationUpdateService, newRidesService, newTripService := handler.NewServices(config, logger, redis, validator, repo, geo, streamRegister)

	mux.Handle(pb.NewDriverServiceHandler(newDriverService, interceptors))
	mux.Handle(pb.NewLocationUpdateServiceHandler(newLocationUpdateService, interceptors))
	mux.Handle(pb.NewRideServiceHandler(newRidesService, interceptors))
	mux.Handle(pb.NewTripServiceHandler(newTripService, interceptors))

	return mux
}

// func BrokerSetup(config *configs.Configurations) (*broker.Exchange, *broker.Broker, error) {
// 	// setup message broker
// 	newbroker := broker.NewBroker(config.Broker.Endpoint,
// 		&broker.EndpointOptions{
// 			Protocol: config.Broker.Protocol,
// 			Username: config.Broker.Secrets.Username,
// 			Password: config.Broker.Secrets.Password,
// 			Port:     config.Broker.Port,
// 		},
// 	)

// 	// build exchange
// 	exchange, err := newbroker.BuildExchange(config.Station.Name)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return exchange, newbroker, nil
// }

// Register New gRPC service with all the conigs.
func NewGrpcServer(
	config *configs.Configurations,
	logger *zap.Logger,
	redis *redis.Client,
	validator *protovalidate.Validator,
	repo repository.RidesInfoRepository,
	geo *geoindex.GeoIndex,
	streamRegister *sr.StreamRegister,

) *http.ServeMux {

	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		logger.Fatal("creating validateInterceptor failed:", zap.NamedError("error", err))
	}

	otelInterceptor, err := otelconnect.NewInterceptor()
	if err != nil {
		logger.Fatal("creating otelInterceptor failed:", zap.NamedError("error", err))
	}

	interceptors := connect.WithInterceptors(otelInterceptor, validateInterceptor)

	mux := http.NewServeMux()

	// register services
	RegisterServices(config, logger, redis, validator, mux, interceptors, repo, geo, streamRegister)
	return mux
}

func StartServer(ctx context.Context, logger *zap.Logger, config *configs.Configurations, mux *http.ServeMux) {
	// Let's allow our queues and http server to drain properly during shutdown.
	// We'll create a channel to listen for SIGINT (Ctrl+C) to signal
	// to our application to gracefully shutdown.
	addr := fmt.Sprintf("0.0.0.0:%s", config.Service.Port)

	srv := &http.Server{
		Addr: addr,
		Handler: h2c.NewHandler(
			newCORS().Handler(mux),
			&http2.Server{},
		),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}
	fmt.Println(addr)
	// middleware := authn.NewMiddleware(authenticate)
	//  handler := middleware.Wrap(mux)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	fmt.Println(addr)
	go func() {
		fmt.Println("Runngin Server")
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println("HTTP listen and serve:", err)
		}
	}()

	<-signals
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("HTTP shutdown:", zap.NamedError("error", err)) //nolint:gocritic
	}
}
