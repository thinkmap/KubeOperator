package repository

import (
	"github.com/KubeOperator/KubeOperator/pkg/constant"
	"github.com/KubeOperator/KubeOperator/pkg/db"
	"github.com/KubeOperator/KubeOperator/pkg/model"
	"github.com/jinzhu/gorm"
)

type ProjectResourceRepository interface {
	PageByProjectIdAndType(num, size int, projectId string, resourceType string) (int, []model.ProjectResource, error)
	Batch(operation string, items []model.ProjectResource) error
	Create(resource model.ProjectResource) error
	ListByResourceIdAndType(resourceId string, resourceType string) ([]model.ProjectResource, error)
	DeleteByResourceIdAnyResourceType(resourceId string, resourceType string) error
	ListByProjectNameAndType(projectName string, resourceType string) ([]model.ProjectResource, error)
}

func NewProjectResourceRepository() ProjectResourceRepository {
	return &projectResourceRepository{}
}

type projectResourceRepository struct {
}

func (p projectResourceRepository) PageByProjectIdAndType(num, size int, projectId string, resourceType string) (int, []model.ProjectResource, error) {
	var total int
	var projectResources []model.ProjectResource
	err := db.DB.
		Model(model.ProjectResource{}).
		Where(model.ProjectResource{ProjectID: projectId, ResourceType: resourceType}).
		Count(&total).
		Find(&projectResources).
		Offset((num - 1) * size).
		Limit(size).Error
	return total, projectResources, err
}

func (p projectResourceRepository) ListByProjectNameAndType(projectName string, resourceType string) ([]model.ProjectResource, error) {
	var project model.Project
	err := db.DB.Model(model.Project{}).Where(model.Project{Name: projectName}).First(&project).Error
	if err != nil {
		return nil, err
	}
	var projectResources []model.ProjectResource
	err = db.DB.Model(model.ProjectResource{}).Where(model.ProjectResource{ProjectID: project.ID, ResourceType: resourceType}).Find(&projectResources).Error
	if err != nil {
		return nil, err
	}
	return projectResources, nil
}

func (p projectResourceRepository) ListByResourceIdAndType(resourceId string, resourceType string) ([]model.ProjectResource, error) {
	var projectResources []model.ProjectResource
	err := db.DB.Model(model.ProjectResource{}).Where(model.ProjectResource{ResourceId: resourceId, ResourceType: resourceType}).Find(&projectResources).Error
	return projectResources, err
}

func (p projectResourceRepository) Batch(operation string, items []model.ProjectResource) error {
	switch operation {
	case constant.BatchOperationDelete:
		tx := db.DB.Begin()
		for _, item := range items {
			if item.ResourceType == constant.ResourceBackupAccount {
				var clusterResources []model.ProjectResource
				err := tx.Where(model.ProjectResource{ProjectID: item.ProjectID, ResourceType: constant.ResourceCluster}).Find(&clusterResources).Error
				if err != nil && !gorm.IsRecordNotFoundError(err) {
					tx.Rollback()
					return err
				}
				if len(clusterResources) > 0 {
					for _, clusterResource := range clusterResources {
						var backupStrategy model.ClusterBackupStrategy
						err = tx.Where(model.ClusterBackupStrategy{BackupAccountID: item.ResourceId, ClusterID: clusterResource.ResourceId}).Delete(&backupStrategy).Error
						if err != nil {
							tx.Rollback()
							return err
						}
					}
				}
			}
			err := tx.Delete(&item).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()

	case constant.BatchOperationCreate:
		tx := db.DB.Begin()
		for i, _ := range items {
			if err := tx.Model(model.ProjectResource{}).Create(&items[i]).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
	default:
		return constant.NotSupportedBatchOperation
	}
	return nil
}

func (p projectResourceRepository) Create(resource model.ProjectResource) error {
	return db.DB.Model(model.ProjectResource{}).Create(&resource).Error
}

func (p projectResourceRepository) DeleteByResourceIdAnyResourceType(resourceId string, resourceType string) error {
	var projectResources []model.ProjectResource
	if resourceId == "" {
		return nil
	}
	err := db.DB.Model(model.ProjectResource{}).Where(model.ProjectResource{ResourceId: resourceId, ResourceType: resourceType}).Find(&projectResources).Error
	if err != nil {
		return err
	}
	var ids []string
	for _, item := range projectResources {
		ids = append(ids, item.ID)
	}
	err = db.DB.Where("id in (?)", ids).Delete(&projectResources).Error
	if err != nil {
		return err
	}
	return nil
}
