// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/rides/v1/rideApp.proto

package ridesconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	rides "github.com/ridehovr/rides/gen/protos/api/rides/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// LocationUpdateServiceName is the fully-qualified name of the LocationUpdateService service.
	LocationUpdateServiceName = "RideBook.LocationUpdateService"
	// RideServiceName is the fully-qualified name of the RideService service.
	RideServiceName = "RideBook.RideService"
	// DriverServiceName is the fully-qualified name of the DriverService service.
	DriverServiceName = "RideBook.DriverService"
	// TripServiceName is the fully-qualified name of the TripService service.
	TripServiceName = "RideBook.TripService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// LocationUpdateServiceUpdateUserLocationProcedure is the fully-qualified name of the
	// LocationUpdateService's UpdateUserLocation RPC.
	LocationUpdateServiceUpdateUserLocationProcedure = "/RideBook.LocationUpdateService/UpdateUserLocation"
	// RideServiceFindRideProcedure is the fully-qualified name of the RideService's FindRide RPC.
	RideServiceFindRideProcedure = "/RideBook.RideService/FindRide"
	// RideServiceTrackRideProcedure is the fully-qualified name of the RideService's TrackRide RPC.
	RideServiceTrackRideProcedure = "/RideBook.RideService/TrackRide"
	// RideServiceTrackTripProcedure is the fully-qualified name of the RideService's TrackTrip RPC.
	RideServiceTrackTripProcedure = "/RideBook.RideService/TrackTrip"
	// DriverServiceActiveDriverProcedure is the fully-qualified name of the DriverService's
	// ActiveDriver RPC.
	DriverServiceActiveDriverProcedure = "/RideBook.DriverService/ActiveDriver"
	// TripServiceTripStartNUpdatesProcedure is the fully-qualified name of the TripService's
	// TripStartNUpdates RPC.
	TripServiceTripStartNUpdatesProcedure = "/RideBook.TripService/TripStartNUpdates"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	locationUpdateServiceServiceDescriptor                  = rides.File_api_rides_v1_rideApp_proto.Services().ByName("LocationUpdateService")
	locationUpdateServiceUpdateUserLocationMethodDescriptor = locationUpdateServiceServiceDescriptor.Methods().ByName("UpdateUserLocation")
	rideServiceServiceDescriptor                            = rides.File_api_rides_v1_rideApp_proto.Services().ByName("RideService")
	rideServiceFindRideMethodDescriptor                     = rideServiceServiceDescriptor.Methods().ByName("FindRide")
	rideServiceTrackRideMethodDescriptor                    = rideServiceServiceDescriptor.Methods().ByName("TrackRide")
	rideServiceTrackTripMethodDescriptor                    = rideServiceServiceDescriptor.Methods().ByName("TrackTrip")
	driverServiceServiceDescriptor                          = rides.File_api_rides_v1_rideApp_proto.Services().ByName("DriverService")
	driverServiceActiveDriverMethodDescriptor               = driverServiceServiceDescriptor.Methods().ByName("ActiveDriver")
	tripServiceServiceDescriptor                            = rides.File_api_rides_v1_rideApp_proto.Services().ByName("TripService")
	tripServiceTripStartNUpdatesMethodDescriptor            = tripServiceServiceDescriptor.Methods().ByName("TripStartNUpdates")
)

// LocationUpdateServiceClient is a client for the RideBook.LocationUpdateService service.
type LocationUpdateServiceClient interface {
	UpdateUserLocation(context.Context) *connect.ClientStreamForClient[rides.UpdateUserLocationRequest, rides.UpdateUserLocationResponse]
}

// NewLocationUpdateServiceClient constructs a client for the RideBook.LocationUpdateService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLocationUpdateServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) LocationUpdateServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &locationUpdateServiceClient{
		updateUserLocation: connect.NewClient[rides.UpdateUserLocationRequest, rides.UpdateUserLocationResponse](
			httpClient,
			baseURL+LocationUpdateServiceUpdateUserLocationProcedure,
			connect.WithSchema(locationUpdateServiceUpdateUserLocationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// locationUpdateServiceClient implements LocationUpdateServiceClient.
type locationUpdateServiceClient struct {
	updateUserLocation *connect.Client[rides.UpdateUserLocationRequest, rides.UpdateUserLocationResponse]
}

// UpdateUserLocation calls RideBook.LocationUpdateService.UpdateUserLocation.
func (c *locationUpdateServiceClient) UpdateUserLocation(ctx context.Context) *connect.ClientStreamForClient[rides.UpdateUserLocationRequest, rides.UpdateUserLocationResponse] {
	return c.updateUserLocation.CallClientStream(ctx)
}

// LocationUpdateServiceHandler is an implementation of the RideBook.LocationUpdateService service.
type LocationUpdateServiceHandler interface {
	UpdateUserLocation(context.Context, *connect.ClientStream[rides.UpdateUserLocationRequest]) (*connect.Response[rides.UpdateUserLocationResponse], error)
}

// NewLocationUpdateServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLocationUpdateServiceHandler(svc LocationUpdateServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	locationUpdateServiceUpdateUserLocationHandler := connect.NewClientStreamHandler(
		LocationUpdateServiceUpdateUserLocationProcedure,
		svc.UpdateUserLocation,
		connect.WithSchema(locationUpdateServiceUpdateUserLocationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/RideBook.LocationUpdateService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LocationUpdateServiceUpdateUserLocationProcedure:
			locationUpdateServiceUpdateUserLocationHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLocationUpdateServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedLocationUpdateServiceHandler struct{}

func (UnimplementedLocationUpdateServiceHandler) UpdateUserLocation(context.Context, *connect.ClientStream[rides.UpdateUserLocationRequest]) (*connect.Response[rides.UpdateUserLocationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("RideBook.LocationUpdateService.UpdateUserLocation is not implemented"))
}

// RideServiceClient is a client for the RideBook.RideService service.
type RideServiceClient interface {
	FindRide(context.Context, *connect.Request[rides.FindRideRequest]) (*connect.ServerStreamForClient[rides.FindRideResponse], error)
	TrackRide(context.Context, *connect.Request[rides.TrackRideRequest]) (*connect.ServerStreamForClient[rides.TrackRideResponse], error)
	TrackTrip(context.Context, *connect.Request[rides.TrackTripRequest]) (*connect.ServerStreamForClient[rides.TrackTripResponse], error)
}

// NewRideServiceClient constructs a client for the RideBook.RideService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRideServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) RideServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &rideServiceClient{
		findRide: connect.NewClient[rides.FindRideRequest, rides.FindRideResponse](
			httpClient,
			baseURL+RideServiceFindRideProcedure,
			connect.WithSchema(rideServiceFindRideMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		trackRide: connect.NewClient[rides.TrackRideRequest, rides.TrackRideResponse](
			httpClient,
			baseURL+RideServiceTrackRideProcedure,
			connect.WithSchema(rideServiceTrackRideMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		trackTrip: connect.NewClient[rides.TrackTripRequest, rides.TrackTripResponse](
			httpClient,
			baseURL+RideServiceTrackTripProcedure,
			connect.WithSchema(rideServiceTrackTripMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// rideServiceClient implements RideServiceClient.
type rideServiceClient struct {
	findRide  *connect.Client[rides.FindRideRequest, rides.FindRideResponse]
	trackRide *connect.Client[rides.TrackRideRequest, rides.TrackRideResponse]
	trackTrip *connect.Client[rides.TrackTripRequest, rides.TrackTripResponse]
}

// FindRide calls RideBook.RideService.FindRide.
func (c *rideServiceClient) FindRide(ctx context.Context, req *connect.Request[rides.FindRideRequest]) (*connect.ServerStreamForClient[rides.FindRideResponse], error) {
	return c.findRide.CallServerStream(ctx, req)
}

// TrackRide calls RideBook.RideService.TrackRide.
func (c *rideServiceClient) TrackRide(ctx context.Context, req *connect.Request[rides.TrackRideRequest]) (*connect.ServerStreamForClient[rides.TrackRideResponse], error) {
	return c.trackRide.CallServerStream(ctx, req)
}

// TrackTrip calls RideBook.RideService.TrackTrip.
func (c *rideServiceClient) TrackTrip(ctx context.Context, req *connect.Request[rides.TrackTripRequest]) (*connect.ServerStreamForClient[rides.TrackTripResponse], error) {
	return c.trackTrip.CallServerStream(ctx, req)
}

// RideServiceHandler is an implementation of the RideBook.RideService service.
type RideServiceHandler interface {
	FindRide(context.Context, *connect.Request[rides.FindRideRequest], *connect.ServerStream[rides.FindRideResponse]) error
	TrackRide(context.Context, *connect.Request[rides.TrackRideRequest], *connect.ServerStream[rides.TrackRideResponse]) error
	TrackTrip(context.Context, *connect.Request[rides.TrackTripRequest], *connect.ServerStream[rides.TrackTripResponse]) error
}

// NewRideServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRideServiceHandler(svc RideServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	rideServiceFindRideHandler := connect.NewServerStreamHandler(
		RideServiceFindRideProcedure,
		svc.FindRide,
		connect.WithSchema(rideServiceFindRideMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	rideServiceTrackRideHandler := connect.NewServerStreamHandler(
		RideServiceTrackRideProcedure,
		svc.TrackRide,
		connect.WithSchema(rideServiceTrackRideMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	rideServiceTrackTripHandler := connect.NewServerStreamHandler(
		RideServiceTrackTripProcedure,
		svc.TrackTrip,
		connect.WithSchema(rideServiceTrackTripMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/RideBook.RideService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RideServiceFindRideProcedure:
			rideServiceFindRideHandler.ServeHTTP(w, r)
		case RideServiceTrackRideProcedure:
			rideServiceTrackRideHandler.ServeHTTP(w, r)
		case RideServiceTrackTripProcedure:
			rideServiceTrackTripHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRideServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRideServiceHandler struct{}

func (UnimplementedRideServiceHandler) FindRide(context.Context, *connect.Request[rides.FindRideRequest], *connect.ServerStream[rides.FindRideResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("RideBook.RideService.FindRide is not implemented"))
}

func (UnimplementedRideServiceHandler) TrackRide(context.Context, *connect.Request[rides.TrackRideRequest], *connect.ServerStream[rides.TrackRideResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("RideBook.RideService.TrackRide is not implemented"))
}

func (UnimplementedRideServiceHandler) TrackTrip(context.Context, *connect.Request[rides.TrackTripRequest], *connect.ServerStream[rides.TrackTripResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("RideBook.RideService.TrackTrip is not implemented"))
}

// DriverServiceClient is a client for the RideBook.DriverService service.
type DriverServiceClient interface {
	ActiveDriver(context.Context) *connect.BidiStreamForClient[rides.ActiveDriverRequest, rides.ActiveDriverResponse]
}

// NewDriverServiceClient constructs a client for the RideBook.DriverService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDriverServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DriverServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &driverServiceClient{
		activeDriver: connect.NewClient[rides.ActiveDriverRequest, rides.ActiveDriverResponse](
			httpClient,
			baseURL+DriverServiceActiveDriverProcedure,
			connect.WithSchema(driverServiceActiveDriverMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// driverServiceClient implements DriverServiceClient.
type driverServiceClient struct {
	activeDriver *connect.Client[rides.ActiveDriverRequest, rides.ActiveDriverResponse]
}

// ActiveDriver calls RideBook.DriverService.ActiveDriver.
func (c *driverServiceClient) ActiveDriver(ctx context.Context) *connect.BidiStreamForClient[rides.ActiveDriverRequest, rides.ActiveDriverResponse] {
	return c.activeDriver.CallBidiStream(ctx)
}

// DriverServiceHandler is an implementation of the RideBook.DriverService service.
type DriverServiceHandler interface {
	ActiveDriver(context.Context, *connect.BidiStream[rides.ActiveDriverRequest, rides.ActiveDriverResponse]) error
}

// NewDriverServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDriverServiceHandler(svc DriverServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	driverServiceActiveDriverHandler := connect.NewBidiStreamHandler(
		DriverServiceActiveDriverProcedure,
		svc.ActiveDriver,
		connect.WithSchema(driverServiceActiveDriverMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/RideBook.DriverService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DriverServiceActiveDriverProcedure:
			driverServiceActiveDriverHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDriverServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDriverServiceHandler struct{}

func (UnimplementedDriverServiceHandler) ActiveDriver(context.Context, *connect.BidiStream[rides.ActiveDriverRequest, rides.ActiveDriverResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("RideBook.DriverService.ActiveDriver is not implemented"))
}

// TripServiceClient is a client for the RideBook.TripService service.
type TripServiceClient interface {
	TripStartNUpdates(context.Context) *connect.BidiStreamForClient[rides.TripStartNUpdatesRequest, rides.TripStartNUpdatesResponse]
}

// NewTripServiceClient constructs a client for the RideBook.TripService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTripServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TripServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &tripServiceClient{
		tripStartNUpdates: connect.NewClient[rides.TripStartNUpdatesRequest, rides.TripStartNUpdatesResponse](
			httpClient,
			baseURL+TripServiceTripStartNUpdatesProcedure,
			connect.WithSchema(tripServiceTripStartNUpdatesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// tripServiceClient implements TripServiceClient.
type tripServiceClient struct {
	tripStartNUpdates *connect.Client[rides.TripStartNUpdatesRequest, rides.TripStartNUpdatesResponse]
}

// TripStartNUpdates calls RideBook.TripService.TripStartNUpdates.
func (c *tripServiceClient) TripStartNUpdates(ctx context.Context) *connect.BidiStreamForClient[rides.TripStartNUpdatesRequest, rides.TripStartNUpdatesResponse] {
	return c.tripStartNUpdates.CallBidiStream(ctx)
}

// TripServiceHandler is an implementation of the RideBook.TripService service.
type TripServiceHandler interface {
	TripStartNUpdates(context.Context, *connect.BidiStream[rides.TripStartNUpdatesRequest, rides.TripStartNUpdatesResponse]) error
}

// NewTripServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTripServiceHandler(svc TripServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	tripServiceTripStartNUpdatesHandler := connect.NewBidiStreamHandler(
		TripServiceTripStartNUpdatesProcedure,
		svc.TripStartNUpdates,
		connect.WithSchema(tripServiceTripStartNUpdatesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/RideBook.TripService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TripServiceTripStartNUpdatesProcedure:
			tripServiceTripStartNUpdatesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTripServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTripServiceHandler struct{}

func (UnimplementedTripServiceHandler) TripStartNUpdates(context.Context, *connect.BidiStream[rides.TripStartNUpdatesRequest, rides.TripStartNUpdatesResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("RideBook.TripService.TripStartNUpdates is not implemented"))
}