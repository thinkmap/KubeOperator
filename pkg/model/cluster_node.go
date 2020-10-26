package model

import (
	"github.com/KubeOperator/KubeOperator/pkg/model/common"
	"github.com/KubeOperator/KubeOperator/pkg/util/ssh"
	"github.com/KubeOperator/kobe/api"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ClusterNode struct {
	common.BaseModel
	ID        string `json:"_"`
	Name      string `json:"name"`
	HostID    string `json:"_"`
	Host      Host   `json:"_" gorm:"save_associations:false"`
	ClusterID string `json:"clusterId"`
	Role      string `json:"role"`
	Status    string `json:"status"`
}

func (n *ClusterNode) BeforeCreate() (err error) {
	n.ID = uuid.NewV4().String()
	return nil
}

func (n ClusterNode) ToKobeHost() *api.Host {
	password, privateKey, _ := n.Host.GetHostPasswordAndPrivateKey()
	return &api.Host{
		Ip:         n.Host.Ip,
		Name:       n.Name,
		Port:       int32(n.Host.Port),
		User:       n.Host.Credential.Username,
		Password:   password,
		PrivateKey: string(privateKey),
	}
}

func (n ClusterNode) ToSSHConfig() ssh.Config {
	password, privateKey, _ := n.Host.GetHostPasswordAndPrivateKey()
	return ssh.Config{
		User:        n.Host.Credential.Username,
		Host:        n.Host.Ip,
		Port:        n.Host.Port,
		PrivateKey:  privateKey,
		Password:    password,
		DialTimeOut: 5 * time.Second,
		Retry:       3,
	}
}
func (n ClusterNode) TableName() string {
	return "ko_cluster_node"
}
