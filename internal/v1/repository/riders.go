package repository

import (
	"errors"

	models "github.com/ridehovr/rides/internal/v1/model"
	"gorm.io/gorm"
)

// RidesInfoRepository defines methods for database operations
type RidesInfoRepository interface {
	CreateRide(ride *models.RidesInfo) (int, error)
	UpdateRide(ride *models.RidesInfo) error
	FindRideByID(id uint) (*models.RidesInfo, error)
	DeleteRideByID(id uint) error
	CreateDriverOnlineLog(ride *models.Driveronlinelogs) error
	UpdateDriverOnlineLog(ride *models.Driveronlinelogs) error
	FindDriverOnlineLog(id uint) (*models.Driveronlinelogs, error)
}

type ridesInfoRepository struct {
	db *gorm.DB
}

// NewRidesInfoRepository creates a new instance of the repository
func NewRidesInfoRepository(db *gorm.DB) RidesInfoRepository {
	return &ridesInfoRepository{db: db}
}

// CreateRide inserts a new ride into the database
func (r *ridesInfoRepository) CreateRide(ride *models.RidesInfo) (int, error) {
	result := r.db.Create(ride)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(ride.RideID), nil
}

// UpdateRide updates an existing ride in the database
func (r *ridesInfoRepository) UpdateRide(ride *models.RidesInfo) error {
	result := r.db.Save(ride)
	return result.Error
}

// FindRideByID retrieves a ride by its ID
func (r *ridesInfoRepository) FindRideByID(id uint) (*models.RidesInfo, error) {
	var ride models.RidesInfo
	result := r.db.First(&ride, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &ride, result.Error
}

// DeleteRideByID removes a ride by its ID
func (r *ridesInfoRepository) DeleteRideByID(id uint) error {
	result := r.db.Delete(&models.RidesInfo{}, id)
	return result.Error
}

// CreateRide inserts a new ride into the database
func (r *ridesInfoRepository) CreateDriverOnlineLog(ride *models.Driveronlinelogs) error {
	result := r.db.Omit("created_at", "updated_at", "deleted_at", "id").Create(ride)
	return result.Error
}

// UpdateRide updates an existing ride in the database
func (r *ridesInfoRepository) UpdateDriverOnlineLog(ride *models.Driveronlinelogs) error {
	result := r.db.Save(ride)
	return result.Error
}

// FindRideByID retrieves a ride by its ID
func (r *ridesInfoRepository) FindDriverOnlineLog(id uint) (*models.Driveronlinelogs, error) {
	var ride models.Driveronlinelogs
	result := r.db.Where("Driverid=?", id).Order("Driverid DESC").First(&ride)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &ride, result.Error
}
