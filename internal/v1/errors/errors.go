package iamerr

import (
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

var (
	ErrInternal            = status.Errorf(codes.Internal, "oops! something went wrong")
	ErrMarshalAndBuild     = status.Errorf(codes.InvalidArgument, "input seems to be wrong")
	ErrUsageUpdate         = status.Errorf(codes.Canceled, "provided usage has been reached")
	ErrIvalidGithubUrl     = status.Errorf(codes.Canceled, "package github url seems to be wrong")
	InvalidConfigureCreds  = status.Errorf(codes.InvalidArgument, "invalid credentials input for configure plugin")
	InvalidConfigureOption = status.Errorf(codes.InvalidArgument, "invalid options input for configure plugin")
	ErrVendorAlreadyExists = status.Errorf(codes.Canceled, "vendor already registered")
)