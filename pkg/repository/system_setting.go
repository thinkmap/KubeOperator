package repository

import (
	"github.com/KubeOperator/KubeOperator/pkg/db"
	"github.com/KubeOperator/KubeOperator/pkg/model"
)

type SystemSettingRepository interface {
	Get(key string) (model.SystemSetting, error)
	List() ([]model.SystemSetting, error)
	Save(systemSetting *model.SystemSetting) error
	ListByTab(tabName string) ([]model.SystemSetting, error)
}

func NewSystemSettingRepository() SystemSettingRepository {
	return &systemSettingRepository{}
}

type systemSettingRepository struct {
}

func (s systemSettingRepository) Get(key string) (model.SystemSetting, error) {
	var systemSetting model.SystemSetting
	if err := db.DB.Where(model.SystemSetting{Key: key}).First(&systemSetting).Error; err != nil {
		return systemSetting, err
	}
	return systemSetting, nil
}

func (s systemSettingRepository) List() ([]model.SystemSetting, error) {
	var systemSettings []model.SystemSetting
	err := db.DB.Model(model.SystemSetting{}).Find(&systemSettings).Error
	return systemSettings, err
}

func (s systemSettingRepository) Save(systemSetting *model.SystemSetting) error {
	if db.DB.NewRecord(systemSetting) {
		return db.DB.Create(&systemSetting).Error
	} else {
		return db.DB.Model(&systemSetting).Updates(&systemSetting).Error
	}
}

func (s systemSettingRepository) ListByTab(tabName string) ([]model.SystemSetting, error) {
	var systemSettings []model.SystemSetting
	err := db.DB.Where(model.SystemSetting{Tab: tabName}).Find(&systemSettings).Error
	return systemSettings, err
}
