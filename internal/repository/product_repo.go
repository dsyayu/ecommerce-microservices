package repository

import (
	"ecommerce-microservices/internal/database"
	"ecommerce-microservices/internal/domain"
)

type ProductRepository interface {
	Create(product *domain.Product) error
	FindAll() ([]domain.Product, error)
	FindById(id uint) (domain.Product, error)
	Update(id uint, product *domain.Product) error
	Delete(id uint) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Create(product *domain.Product) error {
	return database.DB.Create(product).Error
}
func (r *productRepository) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	err := database.DB.Find(&products).Error
	return products, err
}
func (r *productRepository) FindById(id uint) (domain.Product, error) {
	var product domain.Product
	err := database.DB.First(&product, id).Error
	return product, err
}
func (r *productRepository) Update(id uint, product *domain.Product) error {
	return database.DB.Model(&domain.Product{}).Where("id=?", id).Updates(product).Error

}
func (r *productRepository) Delete(id uint) error {
	return database.DB.Delete(&domain.Product{}, id).Error
}
