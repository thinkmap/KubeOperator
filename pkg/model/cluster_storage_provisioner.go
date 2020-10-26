package model

import (
	"github.com/KubeOperator/KubeOperator/pkg/model/common"
	uuid "github.com/satori/go.uuid"
)

type ClusterStorageProvisioner struct {
	common.BaseModel
	ID        string `json:"_"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	Name      string `json:"name"    gorm:"not null;unique"`
	Message   string `json:"message" gorm:"type:text(65535)"`
	Vars      string `json:"_"    gorm:"type:text(65535)"`
	ClusterID string `json:"clusterId"`
}

func (c *ClusterStorageProvisioner) BeforeCreate() (err error) {
	c.ID = uuid.NewV4().String()
	return nil
}

func (c ClusterStorageProvisioner) TableName() string {
	return "ko_cluster_storage_provisioner"
}
