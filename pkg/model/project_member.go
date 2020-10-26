package model

import (
	"github.com/KubeOperator/KubeOperator/pkg/model/common"
	uuid "github.com/satori/go.uuid"
)

type ProjectMember struct {
	common.BaseModel
	ID        string `json:"-" gorm:"type:varchar(64)"`
	ProjectID string `json:"projectId" gorm:"type:varchar(64)"`
	UserID    string `json:"userId" gorm:"type:varchar(64)"`
	Role      string `json:"role" gorm:"type:varchar(64)"`
	User      User   `json:"-" gorm:"save_associations:false"`
}

func (p *ProjectMember) BeforeCreate() (err error) {
	p.ID = uuid.NewV4().String()
	return err
}

func (p ProjectMember) TableName() string {
	return "ko_project_member"
}
