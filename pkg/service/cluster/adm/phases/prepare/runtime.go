package prepare

import (
	"github.com/KubeOperator/KubeOperator/pkg/service/cluster/adm/phases"
	"github.com/KubeOperator/KubeOperator/pkg/util/kobe"
	"io"
)

const (
	playbookNameContainerRuntime = "02-runtime.yml"
)

type ContainerRuntimePhase struct {
}

func (s ContainerRuntimePhase) Name() string {
	return "Install Container Runtime"
}

func (s ContainerRuntimePhase) Run(b kobe.Interface, writer io.Writer) error {
	return phases.RunPlaybookAndGetResult(b, playbookNameContainerRuntime, writer)
}
