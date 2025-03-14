package infrastructure

import (
	"gorm.io/gorm"

	db "v1/src/infraestructure/db/adapter"
	entities "v1/src/modules/products/domain/entities"
)

type MySQLProductDao struct {
	db *gorm.DB
}

func NewMySQLProductDao(connection *db.DBConnection) *MySQLProductDao {
	return &MySQLProductDao{db: connection.DB}
}

func (dao *MySQLProductDao) FindAll() ([]entities.Product, error) {
	var product []entities.Product

	err := dao.db.Model(&entities.Product{}).Find(&product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (dao *MySQLProductDao) FindById(id int) (entities.Product, error) {
	var product entities.Product

	err := dao.db.Model(&entities.Product{}).Where("id = ?", id).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}
