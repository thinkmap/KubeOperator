package model

import (
	"github.com/KubeOperator/KubeOperator/pkg/model/common"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ClusterLog struct {
	common.BaseModel
	ID        string    `json:"_"`
	ClusterID string    `json:"clusterId"`
	Type      string    `json:"type"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Message   string    `json:"message" gorm:"type:text(65535)"`
	Status    string    `json:"status"`
}

func (n *ClusterLog) BeforeCreate() (err error) {
	n.ID = uuid.NewV4().String()
	return nil
}

func (n ClusterLog) TableName() string {
	return "ko_cluster_log"
}
