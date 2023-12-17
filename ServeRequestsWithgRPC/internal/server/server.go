package server

import (
	api "github.com/arima0714/proglog/api/v1"
)

type Config struct {
	CommitLog CommitLog
}

var _ api.LogServer = (*grpcServer)(nil)

type grpcServer struct {
	api.UnimplementedLogServer
	*Config
}

type CommitLog interface {
	Append(*api.Record) (uint64, error)
	Read(uint64) (*api.Record, error)
}

func newgrpcServer(config *Config) (srv *grpcServer, err error) {
	srv = &grpcServer{Config: config}
	return srv, nil
}
