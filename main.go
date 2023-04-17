package main

import (
	"os"
	"strconv"

	"github.com/gvriofernando/test_saham_rakyat/config/postgres"
	"github.com/gvriofernando/test_saham_rakyat/config/redis"
	"github.com/labstack/echo/v4"

	userHttp "github.com/gvriofernando/test_saham_rakyat/service/user/controller/http"
	userRepo "github.com/gvriofernando/test_saham_rakyat/service/user/repository"
	userUseCase "github.com/gvriofernando/test_saham_rakyat/service/user/usecase"

	orderItemHttp "github.com/gvriofernando/test_saham_rakyat/service/order_item/controller/http"
	orderItemRepo "github.com/gvriofernando/test_saham_rakyat/service/order_item/repository"
	orderItemUseCase "github.com/gvriofernando/test_saham_rakyat/service/order_item/usecase"
)

func main() {
	e := echo.New()

	// postgres setting
	pgHost := os.Getenv("PG_HOST")
	pgUser := os.Getenv("PG_USER")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgPort := os.Getenv("PG_PORT")
	pgDbName := os.Getenv("PG_DBNAME")

	// redis setting
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDatabase := os.Getenv("REDIS_DATABASE")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	//Initiate Redis Client
	redisdb, err := strconv.Atoi(redisDatabase)
	if err != nil {
		e.Logger.Fatalf("Failed parsing redis database value of %v", redisDatabase)
	}

	redisService, err := redis.NewClient(redis.Config{
		Address:  redisAddress,
		Password: redisPassword,
		Database: redisdb,
	})

	if err != nil {
		e.Logger.Fatalf("Failed initializing redis: %v", err)
	}

	//Initiate Postgres
	pgdb := postgres.Init(postgres.ConfigDB{
		User:     pgUser,
		Password: pgPassword,
		Host:     pgHost,
		Port:     pgPort,
		Dbname:   pgDbName,
	})

	if pgdb.Error != nil {
		e.Logger.Fatalf("Failed initializing database: %v", pgdb.Error)
	}

	//Initialize User Domain
	userRepoInstance := userRepo.NewUserRepository(userRepo.UserConfig{
		Postgres: pgdb,
		Redis:    redisService,
	})
	userUseCaseInstance := userUseCase.NewUserUseCase(userRepoInstance)
	userHttp.NewUserController(e, userUseCaseInstance)

	//Initialize Order Item Domain
	orderItemRepoInstance := orderItemRepo.NewOrderItemRepository(orderItemRepo.OrderItemConfig{
		Postgres: pgdb,
		Redis:    redisService,
	})
	orderItemUseCaseInstance := orderItemUseCase.NewOrderItemUseCase(orderItemRepoInstance)
	orderItemHttp.NewOrderItemController(e, orderItemUseCaseInstance)

	e.Logger.Fatal(e.Start(":" + port))
}
