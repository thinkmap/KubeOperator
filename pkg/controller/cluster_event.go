package controller

import (
	"github.com/KubeOperator/KubeOperator/pkg/service"
	"github.com/kataras/iris/v12/context"
)

type ClusterEventController struct {
	Ctx                 context.Context
	ClusterEventService service.ClusterEventService
}

func NewClusterEventController() *ClusterEventController {
	return &ClusterEventController{
		ClusterEventService: service.NewClusterEventService(),
	}
}
func (c ClusterEventController) GetNpdBy(clusterName string) (bool, error) {
	return c.ClusterEventService.GetNpd(clusterName)
}

func (c ClusterEventController) PostNpdDeleteBy(clusterName string) (bool, error) {
	return c.ClusterEventService.DeleteNpd(clusterName)
}
func (c ClusterEventController) PostNpdCreateBy(clusterName string) (bool, error) {
	return c.ClusterEventService.CreateNpd(clusterName)
}
