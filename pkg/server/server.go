package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
	"tracking_service/configs"
	_ "tracking_service/docs"
	"tracking_service/internal/delivery/http"
	"tracking_service/internal/domain"
	"tracking_service/internal/repository"
	"tracking_service/internal/service"
	"tracking_service/pkg/db"
)

type Server struct {
	config *configs.Config
	log    *logrus.Logger
	db     *gorm.DB
}

func NewServer(config *configs.Config, log *logrus.Logger) *Server {
	return &Server{
		config: config,
		log:    log,
	}
}

func (s *Server) Run() {
	database, err := db.Connect(s.config.DB)
	if err != nil {
		s.log.Fatal("Could not connect to the database")
	}
	s.db = database

	err = s.db.AutoMigrate(&domain.Order{}, &domain.User{})
	if err != nil {
		s.log.Fatal("Could not run migrations")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s.initRoutes(e)

	s.log.Fatal(e.Start(":" + s.config.ServerPort))
}

func (s *Server) initRoutes(e *echo.Echo) {
	orderRepo := repository.NewOrderRepository(s.db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := http.NewOrderHandler(orderService)

	userRepo := repository.NewUserRepository(s.db)
	userService := service.NewUserService(userRepo, s.config.JWTSecret)
	userHandler := http.NewUserHandler(userService)

	e.POST("/users/register", userHandler.RegisterUser)
	e.POST("/users/login", userHandler.LoginUser)

	r := e.Group("/orders")
	r.Use(http.JWTMiddleware(s.config.JWTSecret))
	r.POST("", orderHandler.CreateOrder)
	r.GET("/:id", orderHandler.GetOrder)
	r.PUT("/:id", orderHandler.UpdateOrder)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
