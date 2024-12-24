package api

import (
	"context"

	config "become_better/src/config"
	gen "become_better/src/gen/become_better"
	"become_better/src/internal/api/services"
)

type MainService struct {
	gen.UnimplementedBecomeBetterServer
	config.App
	Ctx context.Context
	services.MainCategoriesInterface
}
