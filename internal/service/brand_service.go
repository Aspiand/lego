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
	if err := s.db.WithContext(ctx).Model(&models.Brand{}).Where("id = ?", id).Updates(brand).Error; err != nil {
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

func (s *BrandService) List(ctx context.Context, page int, pageSize int) ([]models.Brand, int64, error) {
	var brands []models.Brand
	var total int64

	if err := s.db.WithContext(ctx).Offset((page - 1) * pageSize).Limit(pageSize).Find(&brands).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return brands, total, nil
}

func (s *BrandService) GetByID(ctx context.Context, id uint) (*models.Brand, error) {
	var brand models.Brand

	if err := s.db.WithContext(ctx).First(&brand, id).Error; err != nil {
		return nil, err
	}

	return &brand, nil
}

func (s *BrandService) GetByName(ctx context.Context, name string) (*models.Brand, error) {
	var brand models.Brand

	if err := s.db.WithContext(ctx).Where("name = ?", name).First(&brand).Error; err != nil {
		return nil, err
	}

	return &brand, nil
}

func (s *BrandService) Search(ctx context.Context, name string, page, pageSize int) ([]models.Brand, int64, error) {
	var brands []models.Brand
	var total int64

	err := s.db.WithContext(ctx).Model(&models.Brand{}).Where("name LIKE ?", "%"+name+"%").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = s.db.WithContext(ctx).Where("name LIKE ?", "%"+name+"%").Offset((page - 1) * pageSize).Limit(pageSize).Find(&brands).Error
	if err != nil {
		return nil, 0, err
	}

	return brands, total, nil
}
