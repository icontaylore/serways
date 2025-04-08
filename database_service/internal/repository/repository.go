package repository

import (
	"db_service/models"
	"errors"
	"gorm.io/gorm"
	"time"
)

type InterfaceRepository interface {
	Create(massa string) error
	List() ([]models.DatabaseModels, error)
	Delete(id int) (bool, error)
	Done(id int) (bool, error)
}

type RealizatorRepository struct {
	db *gorm.DB
}

func NewRealizeRepo(db *gorm.DB) InterfaceRepository {
	return &RealizatorRepository{db: db}
}

func (r *RealizatorRepository) Create(massa string) error {
	newModel := &models.DatabaseModels{
		Title:      massa,
		Created_at: time.Now(),
	}

	result := r.db.Create(newModel)
	if result.Error != nil {
		return errors.New("ошибка в создании записи в таблицу")
	}

	return nil
}

func (r *RealizatorRepository) List() ([]models.DatabaseModels, error) {
	var arr []models.DatabaseModels
	if err := r.db.Find(&arr).Error; err != nil {
		return nil, err
	}

	return arr, nil
}

func (r *RealizatorRepository) Delete(id int) (bool, error) {
	var res *models.DatabaseModels

	if err := r.db.First(&res, id).Error; err != nil {
		return false, err
	}

	if err := r.db.Delete(res).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *RealizatorRepository) Done(id int) (bool, error) {
	if err := r.db.Where("id = ?", id).Update("done", true).Error; err != nil {
		return false, err
	}

	return true, nil
}
