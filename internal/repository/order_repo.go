package repository

import (
	"ecommerce-microservices/internal/database"
	"ecommerce-microservices/internal/domain"
)

// OrderRepository adalah interface untuk repository order
type OrderRepository interface {
	Create(order *domain.Order) error
	FindAll() ([]domain.Order, error)
	FindById(id uint) (domain.Order, error)
	Update(id uint, order *domain.Order) error
	Delete(id uint) error
}

// orderRepository adalah implementasi konkret dari OrderRepository
type orderRepository struct{}

// NewOrderRepository membuat instansi baru dari orderRepository
func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

// Create menyimpan order baru ke database
func (r *orderRepository) Create(order *domain.Order) error {
	return database.DB.Create(order).Error
}

// FindAll mengambil semua order dari database
func (r *orderRepository) FindAll() ([]domain.Order, error) {
	var orders []domain.Order
	err := database.DB.Find(&orders).Error
	return orders, err
}

// FindById mengambil order berdasarkan ID
func (r *orderRepository) FindById(id uint) (domain.Order, error) {
	var order domain.Order
	err := database.DB.First(&order, id).Error
	return order, err
}

// Update memperbarui informasi order di database
func (r *orderRepository) Update(id uint, order *domain.Order) error {
	return database.DB.Model(&domain.Order{}).Where("id = ?", id).Updates(order).Error
}

// Delete menghapus order berdasarkan ID
func (r *orderRepository) Delete(id uint) error {
	return database.DB.Delete(&domain.Order{}, id).Error
}
