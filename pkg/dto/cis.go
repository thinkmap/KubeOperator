package dto

import "github.com/KubeOperator/KubeOperator/pkg/model"

type CisTask struct {
	model.CisTask
}

type CisResult struct {
	model.CisResult
}

type CisBatch struct {
	Items     []CisTask
	Operation string
}
