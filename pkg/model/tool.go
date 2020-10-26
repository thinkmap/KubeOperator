package model

import (
	"github.com/KubeOperator/KubeOperator/pkg/model/common"
	uuid "github.com/satori/go.uuid"
)

const (
	supportedArchitectureAll   = "all"
	supportedArchitectureAmd64 = "amd64"
	supportedArchitectureArm64 = "arm64"
)

type ClusterTool struct {
	common.BaseModel
	ID           string `json:"_" gorm:"type:varchar(64)"`
	Name         string `json:"name"`
	ClusterID    string `json:"cluster_id"`
	Version      string `json:"version"`
	Describe     string `json:"describe"`
	Status       string `json:"status"`
	Message      string `json:"message" gorm:"type:text(65535)"`
	Logo         string `json:"logo" `
	Vars         string `json:"_"  gorm:"type:text(65535)"`
	Frame        bool   `json:"frame"`
	Url          string `json:"url"`
	Architecture string `json:"architecture"`
}

func (c *ClusterTool) BeforeCreate() (err error) {
	c.ID = uuid.NewV4().String()
	return nil
}

func (c ClusterTool) TableName() string {
	return "ko_cluster_tool"
}
