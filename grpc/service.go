package main

import (
	"context"
	api "frame/api"
)

type SService struct {
	port string
}

func newSService(port string) api.SServiceServer {
	return &SService{
		port: port,
	}
}

func (s *SService) List(ctx context.Context, req *api.ListReq) (*api.ListResp, error) {
	if req.Messages == "1" {
		return &api.ListResp{
			Name: "hailuo",
			Port: s.port,
		}, nil
	} else {
		return &api.ListResp{
			Name: "hhh",
			Port: s.port,
		}, nil
	}
}
