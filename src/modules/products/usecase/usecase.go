package products

import (
	dao "v1/src/infraestructure/db/dao"
	entities "v1/src/modules/products/domain/entities"
	repository "v1/src/modules/products/domain/repositories"
)

type ProductsUsecase struct {
	repo repository.ProductRepository
}

func NewProductsUsecase(repo *dao.MySQLProductDao) *ProductsUsecase {
	return &ProductsUsecase{
		repo: repo,
	}
}
func (puc *ProductsUsecase) FindAll() ([]entities.Product, error) {
	return puc.repo.FindAll()
}

func (puc *ProductsUsecase) FindById(id int) (entities.Product, error) {
	return puc.repo.FindById(id)
}
