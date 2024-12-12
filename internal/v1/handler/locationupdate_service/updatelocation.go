package locationupdateservice

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"connectrpc.com/connect"
	session "github.com/ridehovr/goutils/session"
	pb "github.com/ridehovr/rides/gen/protos/api/rides/v1"
	model "github.com/ridehovr/rides/internal/v1/model"
)

func (s *LocationUpdateService) UpdateUserLocation(ctx context.Context, stream *connect.ClientStream[pb.UpdateUserLocationRequest]) (*connect.Response[pb.UpdateUserLocationResponse], error) {

	success := true
	isFirst := true

	for stream.Receive() {
		// Receive the next message from the stream

		msg := stream.Msg()
		if msg == nil {
			break // End of stream
		}

		LocUpdateKey := key_prefix_LOC + msg.UserId
		LocSession := session.New(s.redis, LocUpdateKey)
		err := LocSession.Get(ctx)
		if err != nil {

			fmt.Printf("session Doesnot Exist")
			LocSession.SetMetadata("lastUpdated", time.Now().String())
			LocSession.SetMetadata("usertype", msg.UserType)
			LocSession.SetMetadata("lat", strconv.FormatFloat(msg.UserLocation.Latitude, 'f', 2, 64))
			LocSession.SetMetadata("long", strconv.FormatFloat(msg.UserLocation.Longitude, 'f', 2, 64))
			LocSession.Set(ctx, time.Minute*2)

			if msg.UserType == "Driver" {
				DriverId, _ := strconv.Atoi(msg.UserId)
				dberr := s.repo.CreateDriverOnlineLog(&model.Driveronlinelogs{
					Driverid:         DriverId,
					SessionStartTime: timePtr(time.Now()),
					SessionStartLoc:  &model.GPoint{Latitude: msg.UserLocation.Latitude, Longitude: msg.UserLocation.Longitude},
					SessionDate:      timePtr(time.Now()),
				})
				if dberr != nil {
					fmt.Println("db Erron Occured")
					fmt.Println(dberr)
				}
			}

		} else {
			fmt.Println("Session Found")
			LocSession.SetMetadata("lastUpdated", time.Now().String())
			LocSession.SetMetadata("usertype", msg.UserType)
			LocSession.SetMetadata("lat", strconv.FormatFloat(msg.UserLocation.Latitude, 'f', 2, 64))
			LocSession.SetMetadata("long", strconv.FormatFloat(msg.UserLocation.Longitude, 'f', 2, 64))
			LocSession.Set(ctx, time.Minute*2)
		}
		if msg.UserType == "Driver" {
			s.geo.AddDriver(ctx, "geo_"+msg.UserId, msg.UserLocation.Latitude, msg.UserLocation.Longitude)
			fmt.Println("Location Update on Redis Geo")
			if isFirst == true {

			}
		}

		isFirst = false
	}

	if err := stream.Err(); err != nil {
		log.Printf("Stream error: %v", err)
		return nil, err
	}

	// Respond with success status
	return connect.NewResponse(&pb.UpdateUserLocationResponse{
		Success: success,
	}), nil
}
func timePtr(t time.Time) *time.Time {
	return &t
}
