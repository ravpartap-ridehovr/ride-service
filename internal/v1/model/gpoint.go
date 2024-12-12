package models

import (
	"database/sql/driver"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"log"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkb"
)

type GPoint struct {
	Latitude  float64
	Longitude float64
}

// Implement the Scanner interface
func (g *GPoint) Scan(value interface{}) error {
	// Check for the expected []byte type

	str, ok := value.(string)
	if !ok {
		return errors.New("failed to scan GPoint: unsupported type")
	}

	ewkbBytes, err := hex.DecodeString(str)
	if err != nil {
		log.Fatalf("failed to decode hex string: %v", err)
	}

	// Unmarshal the EWKB
	geomPoint, err := ewkb.Unmarshal(ewkbBytes)
	if err != nil {
		log.Fatalf("failed to unmarshal EWKB: %v", err)
	}

	// Assert the geometry type to a Point
	point, ok := geomPoint.(*geom.Point)
	if !ok {
		log.Fatal("geometry is not a point")
	}

	// Get coordinates
	coords := point.Coords()

	g.Longitude = coords[0]
	g.Latitude = coords[1]

	return nil
}

// Implement the Valuer interface to store the GPoint back to the database
func (g GPoint) Value() (driver.Value, error) {
	// Convert GPoint back to WKB for storage in the database
	// Here, we will construct a WKB representation manually

	// Create a new point with SRID 4326 (WGS 84 coordinate system)
	point := geom.NewPoint(geom.XY).MustSetCoords([]float64{g.Longitude, g.Latitude})
	point.SetSRID(4326) // Set SRID to 4326 for geographic coordinates

	// Marshal the point to EWKB
	ewkbBytes, err := ewkb.Marshal(point, binary.BigEndian)
	if err != nil {
		log.Fatalf("failed to marshal point to EWKB: %v", err)
	}

	// Convert EWKB bytes to a hex string
	ewkbHexStr := hex.EncodeToString(ewkbBytes)

	//fmt.Print(ewkbHexStr, "===== string")
	return ewkbHexStr, nil
}
