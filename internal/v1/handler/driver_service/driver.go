package driverservice

import (
	"github.com/bufbuild/protovalidate-go"
	redis "github.com/redis/go-redis/v9"
	configs "github.com/ridehovr/rides/configs"
	v1 "github.com/ridehovr/rides/gen/protos/api/rides/v1/ridesconnect"
	"github.com/ridehovr/rides/internal/v1/repository"
	"github.com/ridehovr/rides/internal/v1/util/geoindex"
	"github.com/ridehovr/rides/internal/v1/util/streamregisters"
	"go.uber.org/zap"
)

const (
	key_prefix_ActiveDriver   = "ActiveDriver_"
	key_prefix_streamRegister = "SR_"
	Online_Driver_Index       = "Online_Driver_Index"
)

type DriverService struct {
	config *configs.Configurations
	//entStore *dynamoutils.EntStore
	// exchange *broker.Exchange
	logger          *zap.Logger
	redis           *redis.Client
	validator       *protovalidate.Validator
	repo            repository.RidesInfoRepository
	geo             *geoindex.GeoIndex
	streamRegisters *streamregisters.StreamRegister
	v1.UnimplementedDriverServiceHandler
}

func NewDriverService(
	config *configs.Configurations,
	logger *zap.Logger,
	rdb *redis.Client,
	validator *protovalidate.Validator,
	repo repository.RidesInfoRepository,
	geo *geoindex.GeoIndex,
	streamRegisters *streamregisters.StreamRegister,
) *DriverService {
	return &DriverService{
		config:          config,
		logger:          logger,
		validator:       validator,
		redis:           rdb,
		repo:            repo,
		geo:             geo,
		streamRegisters: streamRegisters,
	}
}
