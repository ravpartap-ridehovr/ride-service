package geoindex

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

// GeoIndex provides methods to manage geospatial data
type GeoIndex struct {
	Client *redis.Client
	Key    string
}

// NewGeoIndex initializes the GeoIndex with a Redis client and key
func NewGeoIndex(client *redis.Client, key string) *GeoIndex {
	return &GeoIndex{Client: client, Key: key}
}

// AddDriver adds or updates a driver's location
func (g *GeoIndex) AddDriver(ctx context.Context, driverID string, latitude, longitude float64) error {
	location := &redis.GeoLocation{
		Name:      driverID,
		Longitude: longitude,
		Latitude:  latitude,
	}
	mem, err := g.Client.GeoAdd(ctx, g.Key, location).Result()

	//members, err := g.Client.ZRange(ctx, g.Key, 0, -1).Result()

	fmt.Println("All addded members", mem, location, err)
	return err
}

// FindNearbyDrivers finds drivers within the given radius (in kilometers)
func (g *GeoIndex) FindNearbyDrivers(ctx context.Context, latitude, longitude, radius float64) ([]string, error) {
	results, err := g.Client.GeoRadius(ctx, g.Key, longitude, latitude, &redis.GeoRadiusQuery{
		Radius:    radius,
		Unit:      "km",
		WithDist:  true,
		WithCoord: false,
	}).Result()
	if err != nil {
		return nil, err
	}

	var driverIDs []string
	for _, location := range results {
		driverIDs = append(driverIDs, location.Name)
	}
	return driverIDs, nil
}

// RemoveDriver removes a driver from the geospatial index
func (g *GeoIndex) RemoveDriver(ctx context.Context, driverID string) error {
	_, err := g.Client.ZRem(ctx, g.Key, driverID).Result()
	return err
}

func (g *GeoIndex) GetAllMembers(ctx context.Context) ([]string, error) {
	members, err := g.Client.ZRange(ctx, g.Key, 0, -1).Result()
	if err != nil {
		log.Fatalf("Failed to get geo members: %v", err)
		return nil, err
	}
	return members, nil
}
