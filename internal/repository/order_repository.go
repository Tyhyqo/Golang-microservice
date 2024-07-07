package repository

import (
	"tracking_service/internal/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *domain.Order) error
	GetByID(id uint) (*domain.Order, error)
	Update(order *domain.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) GetByID(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.First(&order, id).Error
	return &order, err
}

func (r *orderRepository) Update(order *domain.Order) error {
	return r.db.Save(order).Error
}
