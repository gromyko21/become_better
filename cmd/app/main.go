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

	config "become_better/config"
	gen "become_better/gen/become_better"
	api "become_better/internal/api/become_better"
	swagerDocs "become_better/internal/api/docs"
)


func main() {
	localConfig := config.New()
	ctx := context.Background()

	db, err := config.NewPG(ctx, localConfig.ConnString)
	if err != nil {
		logrus.Fatal(err)
	}

	app := config.App{
		Postgres: db,
	}

	go runRest(localConfig)
	runGRPc(localConfig, app)
}

func runGRPc(localConfig *config.Config, app config.App) {

	lis, err := net.Listen("tcp", ":"+localConfig.CommonConfig.GRPcPort)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	logrus.Info(fmt.Sprintf("Started listening tcp port %s for GRPc", localConfig.CommonConfig.GRPcPort))

	grpcServer := grpc.NewServer()
	gen.RegisterBecomeBetterServer(grpcServer, &api.HelloService{App:app})
	reflection.Register(grpcServer)
	logrus.Info("Service has been started")

	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

func runRest(config *config.Config) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gen.RegisterBecomeBetterHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%s", config.CommonConfig.Host, config.CommonConfig.GRPcPort), opts)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET", "/", swagerDocs.SwaggerFile)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET", "/swagger.json", swagerDocs.SwaggerPage)
	if err != nil {
		panic(err)
	}

	logrus.Info(fmt.Sprintf("Server listening at %s", fmt.Sprintf("%s:%s", config.CommonConfig.Host, config.CommonConfig.HTTPport)))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.CommonConfig.HTTPport), mux); err != nil {
		panic(err)
	}
}
