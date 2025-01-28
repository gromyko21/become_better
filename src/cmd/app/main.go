package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	config "become_better/src/config"
	database "become_better/src/db"
	gen "become_better/src/gen/become_better"
	api "become_better/src/internal/api/become_better"
	swagerDocs "become_better/src/internal/api/docs"
	"become_better/src/internal/models"
	"become_better/src/internal/services"
)

func main() {
	localConfig := config.New()
	ctx := context.Background()

	db, err := database.NewPG(ctx, localConfig.ConnString)
	if err != nil {
		logrus.Fatal(err)
	}

	app := config.App{
		Postgres: db,
	}

	go runRest(localConfig)
	runGRPc(ctx, localConfig, app)
}

func runGRPc(ctx context.Context, localConfig *config.Config, app config.App) {

	lis, err := net.Listen("tcp", ":"+localConfig.CommonConfig.GRPcPort)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	logrus.Info(fmt.Sprintf("Started listening tcp port %s for GRPc", localConfig.CommonConfig.GRPcPort))

	grpcServer := grpc.NewServer()
	gen.RegisterBecomeBetterServer(grpcServer, &api.MainService{
		App: app,
		Ctx: ctx,
		MainCategoriesInterface: &services.CategoriesServiceImpl{
			CategoriesModelInterface: &models.CategoriesModelImpl{},
		},
		ProgressInterface: &services.ProgressService{
			ProgressModelInterface: &models.CategoriesModelImpl{},
			CategoriesModelInterface: &models.CategoriesModelImpl{},
		},
	})
	reflection.Register(grpcServer)
	logrus.Info("Service has been started")

	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

func runRest(config *config.Config) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	
	if err := gen.RegisterBecomeBetterHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%s", config.CommonConfig.Host, config.CommonConfig.GRPcPort), opts); err != nil {
		logrus.Fatalf("Ошибка при регистрации обработчика: %v", err)
	}

	if err := mux.HandlePath("GET", "/", swagerDocs.SwaggerFile); err != nil {
		logrus.Fatalf("Ошибка при обработке пути /: %v", err)
	}

	if err := mux.HandlePath("GET", "/swagger.json", swagerDocs.SwaggerPage); err != nil {
		logrus.Fatalf("Ошибка при обработке пути /swagger.json: %v", err)
	}

	addr := fmt.Sprintf(":%s", config.CommonConfig.HTTPport)
	logrus.Infof("REST сервер запущен на http://%s%s", config.CommonConfig.Host, addr)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			logrus.Errorf("Ошибка при остановке сервера: %v", err)
		}
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
