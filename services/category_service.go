package services

import (
	"go-kasir-api/models"
	"go-kasir-api/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (service *CategoryService) GetAll() ([]models.Category, error) {
	return service.repo.GetAll()
}

func (service *CategoryService) Create(data *models.Category) error {
	return service.repo.Create(data)
}

func (service *CategoryService) GetByID(id int) (*models.Category, error) {
	return service.repo.GetByID(id)
}

func (service *CategoryService) Update(data *models.Category) error {
	return service.repo.Update(data)
}

func (service *CategoryService) Delete(id int) error {
	return service.repo.Delete(id)
}
