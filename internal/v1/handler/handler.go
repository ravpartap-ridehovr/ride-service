package handler

import (
	protovalidate "github.com/bufbuild/protovalidate-go"
	redis "github.com/redis/go-redis/v9"
	zap "go.uber.org/zap"

	// s3 "github.com/defensestation/goutils/presign/providers/s3"

	configs "github.com/ridehovr/rides/configs"
	driver "github.com/ridehovr/rides/internal/v1/handler/driver_service"
	locationupdate "github.com/ridehovr/rides/internal/v1/handler/locationupdate_service"
	ride "github.com/ridehovr/rides/internal/v1/handler/ride_service"
	trip "github.com/ridehovr/rides/internal/v1/handler/trip_service"
	"github.com/ridehovr/rides/internal/v1/repository"
	"github.com/ridehovr/rides/internal/v1/util/geoindex"
	sr "github.com/ridehovr/rides/internal/v1/util/streamregisters"
)

func NewServices(
	config *configs.Configurations,
	logger *zap.Logger,
	redis *redis.Client,
	validator *protovalidate.Validator,
	repo repository.RidesInfoRepository,
	geo *geoindex.GeoIndex,
	streamRegister *sr.StreamRegister,
	// s3 *s3.S3,
) (*driver.DriverService,
	*locationupdate.LocationUpdateService,
	*ride.RideService,
	*trip.TripService,
) {
	driverService := driver.NewDriverService(config, logger, redis, validator, repo, geo, streamRegister)
	locationupdateService := locationupdate.NewLocationUpdateService(config, logger, redis, validator, repo, geo)
	rideService := ride.NewRideService(config, logger, redis, validator, repo, geo, streamRegister)
	tripService := trip.NewTripService(config, logger, redis, validator, repo, geo)
	return driverService, locationupdateService, rideService, tripService
}
