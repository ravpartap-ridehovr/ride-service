package driverservice

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"connectrpc.com/connect"
	session "github.com/ridehovr/goutils/session"
	pb "github.com/ridehovr/rides/gen/protos/api/rides/v1"
)

func (s *DriverService) ActiveDriver(ctx context.Context, stream *connect.BidiStream[pb.ActiveDriverRequest, pb.ActiveDriverResponse]) error {
	firstMsg := true
	fmt.Println("Strem Activated.", firstMsg)
	for {
		msg, err := stream.Receive()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if msg.GetDriverOnlineRequest() != nil && firstMsg == true {
			s.streamRegisters.Lock()
			defer s.streamRegisters.Unlock()
			log.Printf("Driver Status set to online")
			DriverId := msg.GetDriverOnlineRequest().DriverId
			ActiveDriveSessionKey := key_prefix_ActiveDriver + DriverId
			ActiveDriveSession := session.New(s.redis, ActiveDriveSessionKey)
			err := ActiveDriveSession.Get(ctx)
			if err != nil {
				log.Println(err)
			}
			log.Printf("session Doesnot Exist")
			ActiveDriveSession.SetMetadata("currentStatus", "online")

			s.streamRegisters.ActiveDriverStreams[key_prefix_streamRegister+DriverId] = stream

			log.Println(s.streamRegisters)
			log.Println("Added to online session store")
			ActiveDriveSession.SetMetadata("stream", key_prefix_streamRegister+DriverId)
			ActiveDriveSession.Set(ctx, time.Hour*2)

			if err := s.redis.SAdd(ctx, Online_Driver_Index, DriverId).Err(); err != nil {
				return err
			}
			DrOnlineResp := &pb.DriverOnlineResponse{}
			ActiveDriverResp := &pb.ActiveDriverResponse{
				Response: &pb.ActiveDriverResponse_DriverOnlineResponse{
					DriverOnlineResponse: DrOnlineResp,
				},
			}
			stream.Send(ActiveDriverResp)

			log.Println("Stream Added to Registry")
			firstMsg = false
		}

		if msg.GetNewRideResponse() != nil {
			log.Println("Driver Response Received")
			reqMsgs := msg.GetNewRideResponse()
			rideId := reqMsgs.TripId
			rideSession := session.New(s.redis, "ride_"+rideId)
			err := rideSession.Get(ctx)
			if err != nil {
				fmt.Println(err)
			}

			ride_driverId, err := rideSession.GetMetadata("driverId")
			log.Println(ride_driverId, reqMsgs.DriverId)
			if ride_driverId == reqMsgs.DriverId {

				rideSession.SetMetadata("driverResp", reqMsgs.Status)
				rideSession.Set(ctx, time.Minute*10)
				rideSession.Update(ctx)
				testDriverResp, _ := rideSession.GetMetadata("driverResp")
				fmt.Println("========= after update===========")
				fmt.Println(testDriverResp)

			}

		}

		if msg.GetReachedAtPickupRequest() != nil {
			log.Println("Reached at Driver Location")
		}
	}
	return nil
}
