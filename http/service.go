package main

import (
	"context"
	api "frame/api"
	"net/http"
	"net/url"
)

func (g *Gate) List(ctx context.Context, request *http.Request, req url.Values) (interface{}, error) {
	return g.SServiceClient(ctx,
		func(ctx context.Context, ssClient api.SServiceClient) (i interface{}, e error) {
			return ssClient.List(ctx, &api.ListReq{Messages: req.Get("message")})
		})
}
