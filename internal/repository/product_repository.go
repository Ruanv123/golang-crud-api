package repository

import (
	"github.com/ruanv123/api-go-crud/internal/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) error
	FindAll(name string, page, limit int) ([]model.Product, int64, error)
	FindByID(id uint) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

// Create adiciona um novo produto ao banco de dados
func (r *productRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

// FindAll retorna todos os produtos
func (r *productRepository) FindAll(name string, page, limit int) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	query := r.db.Model(&model.Product{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	query.Count(&total)

	err := query.Offset((page - 1) * limit).Limit(limit).Find(&products).Error

	return products, total, err
}

// FindByID retorna um produto pelo seu ID
func (r *productRepository) FindByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

// Update atualiza as informações de um produto
func (r *productRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

// Delete remove um produto pelo seu ID
func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}
