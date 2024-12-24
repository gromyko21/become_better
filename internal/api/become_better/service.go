package api

import (
	"context"

	config "become_better/config"
	gen "become_better/gen/become_better"
	"become_better/internal/api/services"
)

type MainService struct {
	gen.UnimplementedBecomeBetterServer
	config.App
	Ctx context.Context
	services.MainCategoriesInterface
}
