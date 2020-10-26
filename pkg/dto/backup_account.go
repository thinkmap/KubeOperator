package dto

import "github.com/KubeOperator/KubeOperator/pkg/model"

type BackupAccount struct {
	model.BackupAccount
	CredentialVars interface{} `json:"credentialVars"`
}

type BackupAccountOp struct {
	Operation string          `json:"operation" validate:"required"`
	Items     []BackupAccount `json:"items" validate:"required"`
}

type BackupAccountRequest struct {
	Name           string      `json:"name" validate:"required"`
	CredentialVars interface{} `json:"credentialVars" validate:"required"`
	Bucket         string      `json:"bucket" validate:"required"`
	Type           string      `json:"type" validate:"required"`
}

type CloudStorageRequest struct {
	CredentialVars interface{} `json:"credentialVars" validate:"required"`
	Type           string      `json:"type" validate:"required"`
}
