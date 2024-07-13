package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
	"tracking_service/configs"
	_ "tracking_service/docs"
	"tracking_service/internal/domain"
	"tracking_service/internal/handlers/order_handler"
	"tracking_service/internal/handlers/user_handler"
	"tracking_service/internal/http"
	"tracking_service/internal/repository"
	"tracking_service/internal/service"
	"tracking_service/pkg/db"
)

type Server struct {
	app    *echo.Echo
	config *configs.Config
	logger *logrus.Logger
	db     *gorm.DB
}

func NewServer(config *configs.Config, logger *logrus.Logger) *Server {
	return &Server{
		config: config,
		logger: logger,
	}
}

func (s *Server) Run() {
	database, err := db.Connect(s.config.DB)
	if err != nil {
		s.logger.Fatal("Could not connect to the database")
	}
	s.db = database

	err = s.db.AutoMigrate(&domain.Order{}, &domain.UserDTO{})
	if err != nil {
		s.logger.Fatal("Could not run migrations")
	}

	s.app = echo.New()

	s.app.Use(middleware.Logger())
	s.app.Use(middleware.Recover())

	s.initRoutes()

	s.logger.Fatal(s.app.Start(":" + s.config.ServerPort))
}

func (s *Server) initRoutes() {
	orderRepo := repository.NewOrderRepository(s.db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := order_handlers.NewOrderHandler(orderService)

	userRepo := repository.NewUserRepository(s.db)
	userService := service.NewUserService(userRepo, s.config.JWTSecret)
	userHandler := user_handler.NewUserHandler(userService)

	s.app.POST("/users/register", userHandler.RegisterUser)
	s.app.POST("/users/login", userHandler.LoginUser)
	s.app.POST("/users/logout", userHandler.LogoutUser)
	s.app.GET("/users/protected", userHandler.Protected, userService.JWTMiddleware)

	r := s.app.Group("/order_handler")
	r.Use(http.JWTMiddleware1(s.config.JWTSecret))
	r.POST("", orderHandler.CreateOrder)
	r.GET("/:id", orderHandler.GetOrder)
	r.PUT("/:id", orderHandler.UpdateOrder)

	s.app.GET("/swagger/*", echoSwagger.WrapHandler)
}
