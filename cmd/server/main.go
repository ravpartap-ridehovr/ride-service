// main package to start the service
// version: v0.0.1

package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	// "time"

	protovalidate "github.com/bufbuild/protovalidate-go"
	redis "github.com/redis/go-redis/v9"
	zap "go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	helper "github.com/ridehovr/rides/cmd/helper"
	configs "github.com/ridehovr/rides/configs"
	"github.com/ridehovr/rides/internal/v1/repository"
	"github.com/ridehovr/rides/internal/v1/util/geoindex"
	sr "github.com/ridehovr/rides/internal/v1/util/streamregisters"
)

// setting up service and starting server.
func main() {
	ctx := context.Background()
	// initialize configurations
	config, err := configs.New()
	if err != nil {
		log.Fatal("unable get configurations:", err)
	}

	// Initialize Logger
	conn := zap.NewProductionConfig()
	conn.OutputPaths = []string{"stdout"}
	logger, err := conn.Build()
	// check if any error
	if err != nil {
		log.Fatal("unable setup logger:", err)
	}
	defer logger.Sync() // flushes buffer, if any
	// update logger with service name
	logger = logger.With(zap.String("service", config.Service.Name))

	// setup db
	// db, entStore, err := helper.DynamoDBSetup(config)
	// if err != nil {
	// 	logger.Fatal("unable to setup dynamodb:", zap.NamedError("error", err))
	// }

	// // setup broker
	// exchange, broker, err := helper.BrokerSetup(config)
	// if err != nil {
	// 	logger.Fatal("unable to setup broker:", zap.NamedError("error", err))
	// }

	// setup redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Endpoint,
		Password: "",
		DB:       0,
	})

	err = redisClient.Ping(ctx).Err()
	if err != nil {
		logger.Fatal("unable to setup redis client:", zap.NamedError("error", err))
	}

	db, err := gorm.Open(postgres.Open(config.Postgres.Dns), &gorm.Config{})
	geo := geoindex.NewGeoIndex(redisClient, "onlineDrivers_Mohali") //geoindex.NewGeoIndex(redisClient, "onlineDrivers_Mohali")

	streamRegister := sr.NewStreamRegister()

	rideBook := repository.NewRidesInfoRepository(db)

	go listenForExpiryEvents(redisClient, ctx, rideBook, geo)

	// setup validator
	validator, err := protovalidate.New()
	if err != nil {
		logger.Fatal("failed to initialize validator:", zap.NamedError("error", err))
	}

	// gRPC Server Instance
	server := helper.NewGrpcServer(
		config,
		logger,
		redisClient,
		validator,
		rideBook,
		geo,
		streamRegister,
	)
	// start grpc server
	helper.StartServer(ctx, logger, config, server)
}

func listenForExpiryEvents(rdb *redis.Client, ctx context.Context, repo repository.RidesInfoRepository, geo *geoindex.GeoIndex) {
	pubsub := rdb.Subscribe(ctx, "__keyevent@0__:expired") // Adjust DB index if necessary
	defer pubsub.Close()

	fmt.Println("Listening for key expiry events...")

	for msg := range pubsub.Channel() {
		fmt.Printf("Key expired: %s\n", msg.Payload)
		handleSessionExpiry(ctx, msg.Payload, repo, geo)
	}
}

func handleSessionExpiry(ctx context.Context, sessionID string, repo repository.RidesInfoRepository, geo *geoindex.GeoIndex) {
	fmt.Println("========Expired Key=========")
	fmt.Println(sessionID)
	sessionIdArray := strings.Split(sessionID, "_")
	userId, _ := strconv.Atoi(sessionIdArray[1])
	log, _ := repo.FindDriverOnlineLog(uint(userId))
	log.SessionEndTime = timePtr(time.Now())
	repo.UpdateDriverOnlineLog(log)
	// remove driver ===
	geo.RemoveDriver(ctx, sessionIdArray[1])
}

func timePtr(t time.Time) *time.Time {
	return &t
}
