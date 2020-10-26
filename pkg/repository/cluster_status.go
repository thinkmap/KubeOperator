package repository

import (
	"github.com/KubeOperator/KubeOperator/pkg/db"
	"github.com/KubeOperator/KubeOperator/pkg/model"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ClusterStatusRepository interface {
	Get(id string) (model.ClusterStatus, error)
	Save(status *model.ClusterStatus) error
	Delete(id string) error
}

func NewClusterStatusRepository() ClusterStatusRepository {
	return &clusterStatusRepository{
		conditionRepo: NewClusterStatusConditionRepository(),
	}
}

type clusterStatusRepository struct {
	conditionRepo ClusterStatusConditionRepository
}

func (c clusterStatusRepository) Get(id string) (model.ClusterStatus, error) {
	status := model.ClusterStatus{
		ID: id,
	}
	if err := db.DB.
		First(&status).
		Order("last_probe_time asc").
		Related(&status.ClusterStatusConditions).
		Error; err != nil {
		return status, err
	}
	return status, nil
}

func (c clusterStatusRepository) Save(status *model.ClusterStatus) error {
	tx := db.DB.Begin()
	if db.DB.NewRecord(status) {
		if err := db.DB.Create(&status).Error; err != nil {
			return err
		}
	} else {
		// 先记录原来的状态
		var oldStatus model.ClusterStatus
		db.DB.First(&oldStatus)
		if status.Phase != oldStatus.Phase {
			status.PrePhase = oldStatus.Phase
		}
		if err := db.DB.Save(&status).Error; err != nil {
			return err
		}
	}
	// 先清空所有 condition
	var oldConditions []model.ClusterStatusCondition
	notFound := tx.Model(model.ClusterStatusCondition{}).Where(model.ClusterStatusCondition{ClusterStatusID: status.ID}).Find(&oldConditions).RecordNotFound()
	if !notFound {
		for _, c := range oldConditions {
			tx.Delete(&c)
		}
	}
	// 保存最新的conditons
	for i, _ := range status.ClusterStatusConditions {
		status.ClusterStatusConditions[i].ClusterStatusID = status.ID
		if tx.NewRecord(status.ClusterStatusConditions[i]) {
			var temp model.ClusterStatusCondition
			if tx.Where(model.ClusterStatusCondition{ClusterStatusID: status.ClusterStatusConditions[i].ClusterStatusID, Name: status.ClusterStatusConditions[i].Name}).
				First(&temp).
				RecordNotFound() {
				status.ClusterStatusConditions[i].CreatedAt = time.Now()
				status.ClusterStatusConditions[i].UpdatedAt = time.Now()
				status.ClusterStatusConditions[i].ID = uuid.NewV4().String()
				if err := tx.Create(status.ClusterStatusConditions[i]).Error; err != nil {
					tx.Rollback()
					return err
				}
			} else {
				status.ClusterStatusConditions[i].ID = temp.ID
				status.ClusterStatusConditions[i].UpdatedAt = time.Now()
				if err := tx.Save(status.ClusterStatusConditions[i]).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
		} else {
			if err := tx.Save(status.ClusterStatusConditions[i]).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()
	return nil
}

func (c clusterStatusRepository) Delete(id string) error {
	if err := db.DB.
		First(&model.Cluster{ID: id}).
		Delete(model.Cluster{}).Error; err != nil {
		return err
	}
	return nil
}
