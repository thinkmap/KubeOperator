package service

import (
	"encoding/json"
	"errors"
	"github.com/KubeOperator/KubeOperator/pkg/cloud_storage"
	"github.com/KubeOperator/KubeOperator/pkg/constant"
	"github.com/KubeOperator/KubeOperator/pkg/controller/page"
	"github.com/KubeOperator/KubeOperator/pkg/dto"
	"github.com/KubeOperator/KubeOperator/pkg/model"
	"github.com/KubeOperator/KubeOperator/pkg/model/common"
	"github.com/KubeOperator/KubeOperator/pkg/repository"
	"os"
)

var (
	CheckFailed            = "CHECK_FAILED"
	BackupAccountNameExist = "NAME_EXISTS"
)

type BackupAccountService interface {
	Get(name string) (*dto.BackupAccount, error)
	List(projectName string) ([]dto.BackupAccount, error)
	Page(num, size int) (page.Page, error)
	Create(creation dto.BackupAccountRequest) (*dto.BackupAccount, error)
	Update(creation dto.BackupAccountRequest) (*dto.BackupAccount, error)
	Batch(op dto.BackupAccountOp) error
	GetBuckets(request dto.CloudStorageRequest) ([]interface{}, error)
	Delete(name string) error
}

type backupAccountService struct {
	backupAccountRepo repository.BackupAccountRepository
}

func NewBackupAccountService() BackupAccountService {
	return &backupAccountService{
		backupAccountRepo: repository.NewBackupAccountRepository(),
	}
}

func (b backupAccountService) Get(name string) (*dto.BackupAccount, error) {
	var backupAccountDTO dto.BackupAccount
	mo, err := b.backupAccountRepo.Get(name)
	if err != nil {
		return nil, err
	}
	backupAccountDTO = dto.BackupAccount{
		BackupAccount: *mo,
	}
	return &backupAccountDTO, nil
}

func (b backupAccountService) List(projectName string) ([]dto.BackupAccount, error) {
	var backupAccountDTOs []dto.BackupAccount
	mos, err := b.backupAccountRepo.List(projectName)
	if err != nil {
		return nil, err
	}
	for _, mo := range mos {
		backupAccountDTOs = append(backupAccountDTOs, dto.BackupAccount{BackupAccount: mo})
	}
	return backupAccountDTOs, nil
}

func (b backupAccountService) Page(num, size int) (page.Page, error) {
	var page page.Page
	var backupAccountDTOs []dto.BackupAccount
	total, mos, err := b.backupAccountRepo.Page(num, size)
	if err != nil {
		return page, err
	}
	for _, mo := range mos {
		backupDTO := new(dto.BackupAccount)
		vars := make(map[string]interface{})
		json.Unmarshal([]byte(mo.Credential), &vars)
		backupDTO.CredentialVars = vars
		backupDTO.BackupAccount = mo

		backupAccountDTOs = append(backupAccountDTOs, *backupDTO)
	}
	page.Total = total
	page.Items = backupAccountDTOs
	return page, err
}

func (b backupAccountService) Create(creation dto.BackupAccountRequest) (*dto.BackupAccount, error) {

	old, _ := b.Get(creation.Name)
	if old != nil && old.ID != "" {
		return nil, errors.New(BackupAccountNameExist)
	}

	err := b.CheckValid(creation)
	if err != nil {
		return nil, err
	}
	credential, _ := json.Marshal(creation.CredentialVars)
	backupAccount := model.BackupAccount{
		Name:       creation.Name,
		Bucket:     creation.Bucket,
		Type:       creation.Type,
		Credential: string(credential),
		Status:     constant.Valid,
	}

	err = b.backupAccountRepo.Save(&backupAccount)
	if err != nil {
		return nil, err
	}

	return &dto.BackupAccount{BackupAccount: backupAccount}, err
}

func (b backupAccountService) Update(creation dto.BackupAccountRequest) (*dto.BackupAccount, error) {

	err := b.CheckValid(creation)
	if err != nil {
		return nil, err
	}
	credential, _ := json.Marshal(creation.CredentialVars)
	old, err := b.backupAccountRepo.Get(creation.Name)
	if err != nil {
		return nil, err
	}
	backupAccount := model.BackupAccount{
		ID:         old.ID,
		Name:       creation.Name,
		Bucket:     creation.Bucket,
		Type:       creation.Type,
		Credential: string(credential),
		Status:     constant.Valid,
	}

	err = b.backupAccountRepo.Save(&backupAccount)
	if err != nil {
		return nil, err
	}

	return &dto.BackupAccount{BackupAccount: backupAccount}, err
}

func (b backupAccountService) Batch(op dto.BackupAccountOp) error {
	var deleteItems []model.BackupAccount
	for _, item := range op.Items {
		deleteItems = append(deleteItems, model.BackupAccount{
			BaseModel: common.BaseModel{},
			ID:        item.ID,
			Name:      item.Name,
		})
	}
	err := b.backupAccountRepo.Batch(op.Operation, deleteItems)
	if err != nil {
		return err
	}
	return nil
}

func (b backupAccountService) GetBuckets(request dto.CloudStorageRequest) ([]interface{}, error) {
	vars := request.CredentialVars.(map[string]interface{})
	vars["type"] = request.Type
	client, err := cloud_storage.NewCloudStorageClient(vars)
	if err != nil {
		return nil, err
	}
	return client.ListBuckets()
}

func (b backupAccountService) CheckValid(create dto.BackupAccountRequest) error {
	vars := create.CredentialVars.(map[string]interface{})
	vars["type"] = create.Type
	vars["bucket"] = create.Bucket
	client, err := cloud_storage.NewCloudStorageClient(vars)
	if err != nil {
		return err
	}
	file, err := os.Create(constant.DefaultFireName)
	if err != nil {
		return err
	}
	defer file.Close()
	success, err := client.Upload(constant.DefaultFireName, constant.DefaultFireName)
	if err != nil {
		return err
	}
	if !success {
		return errors.New(CheckFailed)
	} else {
		deleteSuccess, err := client.Delete(constant.DefaultFireName)
		if err != nil {
			return err
		}
		if !deleteSuccess {
			return errors.New(CheckFailed)
		}
	}
	return nil
}

func (b backupAccountService) Delete(name string) error {
	return b.backupAccountRepo.Delete(name)
}
