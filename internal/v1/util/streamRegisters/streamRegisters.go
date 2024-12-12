package streamregisters

import (
	"sync"

	"connectrpc.com/connect"
	pb "github.com/ridehovr/rides/gen/protos/api/rides/v1"
)

// GeoIndex provides methods to manage geospatial data
type StreamRegister struct {
	sync.RWMutex
	ActiveDriverStreams map[string]*connect.BidiStream[pb.ActiveDriverRequest, pb.ActiveDriverResponse]
}

func NewStreamRegister() *StreamRegister {
	return &StreamRegister{
		ActiveDriverStreams: make(map[string]*connect.BidiStream[pb.ActiveDriverRequest, pb.ActiveDriverResponse]),
	}
}
