package services

import (
	"cashier/models"
	"cashier/repositories"
)

type ProductsService struct {
	repo *repositories.ProductsRepository
}

func NewProductsService(repo *repositories.ProductsRepository) *ProductsService {
	return &ProductsService{repo: repo}
}

func (s *ProductsService) GetAll() ([]models.Products, error) {
	return s.repo.GetAll()
}

func (s *ProductsService) GetByID(id int) (*models.Products, error) {
	return s.repo.GetByID(id)
}

func (s *ProductsService) Create(data *models.Products) error {
	return s.repo.Create(data)
}

func (s *ProductsService) Update(product *models.Products) error {
	return s.repo.Update(product)
}

func (s *ProductsService) Delete(id int) error {
	return s.repo.Delete(id)
}
