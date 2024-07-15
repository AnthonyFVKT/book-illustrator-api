package main

import (
	"github.com/AnthonyFVKT/book-illustrator-api/internal/config"
	"github.com/AnthonyFVKT/book-illustrator-api/internal/handler"
	"github.com/AnthonyFVKT/book-illustrator-api/internal/rpc"
	"github.com/AnthonyFVKT/book-illustrator-api/internal/service"
	"github.com/AnthonyFVKT/book-illustrator-api/internal/validator"
	pb "github.com/AnthonyFVKT/book-illustrator-srv/proto/illustrator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	illustrateGroup = "/illustrate"
	profileGroup    = "/profile"
)

func main() {
	/*ctx, cancel := context.WithCancel(context.Background())
	defer cancel()*/

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("unable to load config: %v", err)
	}

	illustratorConn, err := grpc.NewClient(cfg.IllustratorServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("unable to connect to illustrator service: %v", err)
	}
	illustratorClient := pb.NewIllustratorServiceClient(illustratorConn)

	illustratorRpc := rpc.NewIllustrator(illustratorClient)
	/*pgxConn, err := pgxpool.New(ctx, cfg.PostgresEndpoint)
	if err != nil {
		log.Fatalf("unable to connect to postgres: %v", err)
	}*/

	//profileReppository := repository.NewProfile(pgxConn)

	illustratorService := service.NewIllustrator(illustratorRpc)
	//profileService := service.NewProfile(profileReppository)

	illustratorHandler := handler.NewIllustrator(illustratorService)
	//profileHandler := handler.NewProfile(profileService)

	e := echo.New()
	e.Validator, err = validator.NewStructValidator()
	if err != nil {
		log.Fatalf("unable to create validator: %v", err)
	}
	e.Use(middleware.CORS())
	illustrate := e.Group(illustrateGroup)
	illustrate.POST("", illustratorHandler.Illustrate)

	//profile := e.Group(profileGroup)
	//profile.POST("", profileHandler.Create)

	e.Logger.Fatal(e.Start(cfg.ServerAddress))
}
