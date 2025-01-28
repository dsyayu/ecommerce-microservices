package service

import (
	"ecommerce-microservices/internal/domain"
	"ecommerce-microservices/internal/repository"
)

type ProductService interface {
	Create(product *domain.Product) error
	GetAll() ([]domain.Product, error)
	GetById(id uint) (domain.Product, error)
	Update(id uint, product *domain.Product) error
	Delete(id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) Create(product *domain.Product) error {
	return s.repo.Create(product)
}
func (s *productService) GetAll() ([]domain.Product, error) {
	return s.repo.FindAll()
}
func (s *productService) GetById(id uint) (domain.Product, error) {
	return s.repo.FindById(id)
}
func (s *productService) Update(id uint, product *domain.Product) error {
	return s.repo.Update(id, product)
}
func (s *productService) Delete(id uint) error {
	return s.repo.Delete(id)
}
