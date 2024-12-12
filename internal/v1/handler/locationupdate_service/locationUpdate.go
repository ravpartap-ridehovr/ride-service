package locationupdateservice

import (
	"github.com/bufbuild/protovalidate-go"
	redis "github.com/redis/go-redis/v9"
	configs "github.com/ridehovr/rides/configs"
	v1 "github.com/ridehovr/rides/gen/protos/api/rides/v1/ridesconnect"
	"github.com/ridehovr/rides/internal/v1/repository"
	"github.com/ridehovr/rides/internal/v1/util/geoindex"
	"go.uber.org/zap"
)

const (
	key_prefix_LOC = "loc_"
)

type LocationUpdateService struct {
	config *configs.Configurations
	//entStore *dynamoutils.EntStore
	// exchange *broker.Exchange
	logger    *zap.Logger
	redis     *redis.Client
	validator *protovalidate.Validator
	repo      repository.RidesInfoRepository
	geo       *geoindex.GeoIndex
	v1.LocationUpdateServiceHandler
}

func NewLocationUpdateService(
	config *configs.Configurations,
	logger *zap.Logger,
	rdb *redis.Client,
	validator *protovalidate.Validator,
	repo repository.RidesInfoRepository,
	geo *geoindex.GeoIndex) *LocationUpdateService {
	return &LocationUpdateService{
		config:    config,
		logger:    logger,
		validator: validator,
		redis:     rdb,
		repo:      repo,
		geo:       geo,
	}
}
