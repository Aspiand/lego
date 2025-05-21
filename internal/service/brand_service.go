package service

import (
	"context"

	"github.com/Aspiand/lego/internal/models"
	"gorm.io/gorm"
)

type BrandService struct {
	db *gorm.DB
}

func NewBrandService(db *gorm.DB) *BrandService {
	return &BrandService{
		db: db,
	}
}

func (s *BrandService) Create(ctx context.Context, brand *models.Brand) error {
	if err := s.db.WithContext(ctx).Create(&brand).Error; err != nil {
		return err
	}

	return nil
}

func (s *BrandService) Update(ctx context.Context, id int, brand models.Brand) error {
	if err := s.db.WithContext(ctx).
		Model(&models.Brand{}).
		Where("id = ?", id).
		Updates(brand).Error; err != nil {
		return err
	}

	return nil
}

func (s *BrandService) Delete(ctx context.Context, id int) error {
	if err := s.db.WithContext(ctx).Delete(&models.Brand{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (s *BrandService) List(ctx context.Context, name string, page int, pageSize int) ([]models.Brand, error) {
	var brands []models.Brand

	if err := s.db.WithContext(ctx).
		Offset((page-1)*pageSize).
		Limit(pageSize).
		Where("name LIKE ?", "%"+name+"%"). // TODO: Handler later
		Find(&brands).Error; err != nil {
		return nil, err
	}

	return brands, nil
}

func (s *BrandService) GetByID(ctx context.Context, id int) (*models.Brand, error) {
	var brand models.Brand

	if err := s.db.WithContext(ctx).First(&brand, id).Error; err != nil {
		return nil, err
	}

	return &brand, nil
}
