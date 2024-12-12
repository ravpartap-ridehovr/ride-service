package rideservice

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"connectrpc.com/connect"
	session "github.com/ridehovr/goutils/session"
	pb "github.com/ridehovr/rides/gen/protos/api/rides/v1"
	models "github.com/ridehovr/rides/internal/v1/model"
)

func (s *RideService) FindRide(ctx context.Context, req *connect.Request[pb.FindRideRequest], stream *connect.ServerStream[pb.FindRideResponse]) error {

	pickupPoint := req.Msg.PickUp
	dropOff := req.Msg.DropOf

	log.Println(pickupPoint, dropOff)

	riderID_Int, _ := strconv.Atoi(req.Msg.RiderId)
	distance := 20.00
	rideId, rideCreateErr := s.repo.CreateRide(&models.RidesInfo{
		RiderID:      riderID_Int,
		DropOffLoc:   models.GPoint{Latitude: dropOff.Latitude, Longitude: dropOff.Longitude},
		PickupLoc:    models.GPoint{Latitude: pickupPoint.Latitude, Longitude: pickupPoint.Longitude},
		DriverID:     -1,
		InitPrice:    10.00,
		InitDistance: &distance,
		WaitTime:     0,
		RideStatusID: 1,
		OTP:          "1234",
	})

	if rideCreateErr != nil {
		log.Println("erron ouccred in create ride entry===>")
		log.Panicln(rideCreateErr)
	}

	rideSession := session.New(s.redis, "ride_"+strconv.Itoa(rideId))
	rideSession.SetMetadata("rideId", strconv.Itoa(rideId))
	rideSession.SetMetadata("driverId", "-1")
	rideSession.SetMetadata("rideStage", "init")
	rideSession.Set(ctx, time.Minute*10)

	// nearbyUsers, _ := s.geo.FindNearbyDrivers(ctx, pickupPoint.Latitude, pickupPoint.Latitude, 100)
	nearbyUsers, _ := s.geo.GetAllMembers(ctx)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	go func() {

		for _, driverId := range nearbyUsers {
			fmt.Println("checking the driver", driverId)
			driverIdParts := strings.Split(driverId, "_")
			driver_Session := session.New(s.redis, key_prefix_ActiveDriver+driverIdParts[1])
			driver_Session.Get(ctx)
			driver_Status, err := driver_Session.GetMetadata("currentStatus")
			if err != nil {
				log.Println(err)
			}

			if driver_Status != "online" {
				fmt.Println("Skipping the driver", driverId)
				continue
			}

			rideSession.SetMetadata("driverId", driverIdParts[1])
			rideSession.SetMetadata("driverResp", "-1")

			updateErr := rideSession.Update(ctx)
			if updateErr != nil {
				log.Println(updateErr)
			}

			ride_driverId, err := rideSession.GetMetadata("driverId")
			fmt.Println("update session print 1 value", ride_driverId)

			updateErr = rideSession.Set(ctx, time.Minute*10)
			if updateErr != nil {
				log.Println(updateErr)
			}

			ride_driverId, err = rideSession.GetMetadata("driverId")
			fmt.Println("set session print 2 value", ride_driverId)

			driverStreamKey, err := driver_Session.GetMetadata("stream")
			if err != nil {
				log.Println(err)
			}

			driverStreamObj := s.streamRegisters.ActiveDriverStreams[driverStreamKey]
			fmt.Println("driver stream obj", driverStreamKey, driverStreamObj)

			streamErr := driverStreamObj.Send(&pb.ActiveDriverResponse{
				Response: &pb.ActiveDriverResponse_NewRideRequest{
					NewRideRequest: &pb.NewRideRequest{
						RideId:   strconv.Itoa(rideId),
						Distance: "20km",
						Price:    "10.00",
					},
				},
			})

			if streamErr != nil {
				log.Println(streamErr)
			}

			time.Sleep(60 * time.Second)

			driverId_int, err := strconv.Atoi(driverIdParts[1])
			if err != nil {
				log.Println("Something went wrong, try later")
				log.Println(err)
			}
			rideSession.Get(ctx)
			rideSession.Update(ctx)
			driverResp, err := rideSession.GetMetadata("driverResp")
			if err != nil {
				log.Println("Something went wrong, try later")
				log.Println(err)
			}
			log.Printf(driverResp)
			if driverResp == "1" {
				ride, _ := s.repo.FindRideByID(uint(rideId))
				ride.DriverID = driverId_int
				s.repo.UpdateRide(ride)
				log.Printf("Driver Assigned to Ride.")
				break
			}

		}
	}()

	for {
		select {
		case <-ctx.Done():

			return ctx.Err()

		case t := <-ticker.C:
			fmt.Println(t)

		}
	}

	//DriverOnlineNearBy := make(chan []string)
	// // Collect results
	// for result := range DriverOnlineNearBy {
	// 	fmt.Println(result)
	// }

	// counter := 0
	// for {

	// 	rideDetails := &pb.RideDetails{
	// 		Otp:                  "123",
	// 		DriverID:             "123",
	// 		DriverLicensePlateNo: "123",
	// 		TripId:               "123",
	// 	}
	// 	response := &pb.FindRideResponse{
	// 		Request: &pb.FindRideResponse_RideDetails{
	// 			RideDetails: rideDetails,
	// 		},
	// 	}
	// 	stream.Send(response)
	// 	counter++
	// 	if counter > 3 {
	// 		break
	// 	}

	// }

	return nil

}

// func searchNearDriver(ctx context.Context, loc *pb.LOC, geo *geoindex.GeoIndex, rdb *redis.Client, radius float64) ([]string, error) {

// 	return onlineNearbyUsers, nil

// }
