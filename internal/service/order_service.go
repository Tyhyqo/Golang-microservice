package service

import (
	"tracking_service/internal/domain"
	"tracking_service/internal/repository"
)

type OrderService interface {
	CreateOrder(order *domain.Order) error
	GetOrderById(id uint) (*domain.Order, error)
	UpdateOrder(order *domain.Order) error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(order *domain.Order) error {
	return s.repo.Create(order)
}

func (s *orderService) GetOrderById(id uint) (*domain.Order, error) {
	return s.repo.GetByID(id)
}

func (s *orderService) UpdateOrder(order *domain.Order) error {
	return s.repo.Update(order)
}
