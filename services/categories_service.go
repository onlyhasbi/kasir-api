package services

import (
	"cashier/models"
	"cashier/repositories"
)

type CategoryService struct {
	repo *repositories.CategoriesRepository
}

func NewCategoriesService(repo *repositories.CategoriesRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]models.Categories, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetByID(id int) (*models.Categories, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Create(data *models.Categories) error {
	return s.repo.Create(data)
}

func (s *CategoryService) Update(category *models.Categories) error {
	return s.repo.Update(category)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
