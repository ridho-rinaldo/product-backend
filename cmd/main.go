package main

import (
	"fmt"
	"net/http"
	"os"

	echoPrometheus "github.com/globocom/echo-prometheus"
	config "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ridho-rinaldo/product-backend/config/postgresql"
	"github.com/rs/zerolog/log"
	"go.elastic.co/apm/module/apmechov4/v2"

	exampleHandler "github.com/ridho-rinaldo/product-backend/pkg/example/handler"
	exampleRepository "github.com/ridho-rinaldo/product-backend/pkg/example/repository"
	exampleUsecase "github.com/ridho-rinaldo/product-backend/pkg/example/usecase"

	productHandler "github.com/ridho-rinaldo/product-backend/pkg/product/handler"
	productRepository "github.com/ridho-rinaldo/product-backend/pkg/product/repository"
	productUsecase "github.com/ridho-rinaldo/product-backend/pkg/product/usecase"
)

func main() {
	// if os.Getenv("GO_ENV") == "local" {
	if err := config.Load(".env"); err != nil {
		fmt.Println(".env is not loaded properly")
		fmt.Println(err)
		os.Exit(2)
	}
	// }

	dbConn := postgresql.CreateConnection()

	r := echo.New()
	r.Debug = true
	r.Use(echoPrometheus.MetricsMiddleware())
	r.Use(apmechov4.Middleware())
	r.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	r.Use(middleware.Recover())
	r.Use(middleware.Logger())
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization", "traceparent", "tracestate"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	r.GET("/", func(context echo.Context) error {
		return context.HTML(http.StatusOK, "<strong>Golang Test "+os.Getenv("APP_VERSION")+"</strong>")
	})

	apiV1 := r.Group("/api/v1")

	exampleRepository := exampleRepository.NewExampleRepository(dbConn)
	exampleUsecase := exampleUsecase.NewExampleUsecase(exampleRepository)
	exampleHandler.NewHttpHandler(exampleUsecase).Mount(apiV1, dbConn)

	productRepository := productRepository.NewProductRepository(dbConn)
	productUsecase := productUsecase.NewProductUsecase(productRepository)
	productHandler.NewHttpHandler(productUsecase).Mount(apiV1, dbConn)

	err := r.Start(":" + os.Getenv("PORT"))
	if err != nil {
		log.Error().Msg(err.Error())
	}

}
