package server

import (
	"context"
	"github.com/edcab/gokit-helloworld/pkg/helloworld"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Hello endpoint.Endpoint
}

func MakeGetEndpoint(srv helloworld.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		d, err := srv.SayHelloToMe(ctx, request)
		if err != nil {
			return getResponse{d, err.Error()}, nil
		}
		return getResponse{d, ""}, nil
	}
}
