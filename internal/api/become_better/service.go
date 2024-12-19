package api

import (
	gen "become_better/gen/become_better"
	config "become_better/config"
)

type HelloService struct {
	gen.UnimplementedBecomeBetterServer
	config.App
}
