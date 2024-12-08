package api

import (
	"context"
	
	gen "become_better/gen/become_better"
)

func (s *HelloService) SayHello(ctx context.Context, req *gen.HelloRequest) (*gen.HelloResponse, error) {
    return &gen.HelloResponse{
        Message: "testsss",
    }, nil
}