package repositories

import (
	entities "v1/src/modules/products/domain/entities"
)

type ProductRepository interface {
	FindAll() ([]entities.Product, error)
	FindById(int) (entities.Product, error)
}
