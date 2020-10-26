package service

import (
	"encoding/json"
	"errors"
	"github.com/KubeOperator/KubeOperator/pkg/cloud_provider/client"
	"github.com/KubeOperator/KubeOperator/pkg/controller/page"
	"github.com/KubeOperator/KubeOperator/pkg/dto"
	"github.com/KubeOperator/KubeOperator/pkg/model"
	"github.com/KubeOperator/KubeOperator/pkg/model/common"
	"github.com/KubeOperator/KubeOperator/pkg/repository"
)

var (
	DeleteRegionError = "DELETE_REGION_FAILED_RESOURCE"
	RegionNameExist   = "NAME_EXISTS"
)

type RegionService interface {
	Get(name string) (dto.Region, error)
	List() ([]dto.Region, error)
	Page(num, size int) (page.Page, error)
	Delete(name string) error
	Create(creation dto.RegionCreate) (*dto.Region, error)
	Batch(op dto.RegionOp) error
	ListDatacenter(creation dto.RegionDatacenterRequest) ([]string, error)
}

type regionService struct {
	regionRepo repository.RegionRepository
	zoneRepo   repository.ZoneRepository
}

func NewRegionService() RegionService {
	return &regionService{
		regionRepo: repository.NewRegionRepository(),
	}
}

func (r regionService) Get(name string) (dto.Region, error) {
	var regionDTO dto.Region
	mo, err := r.regionRepo.Get(name)
	if err != nil {
		return regionDTO, err
	}

	m := make(map[string]interface{})
	regionDTO.Region = mo
	json.Unmarshal([]byte(mo.Vars), &m)
	regionDTO.RegionVars = m

	return regionDTO, err
}

func (r regionService) List() ([]dto.Region, error) {
	var regionDTOs []dto.Region
	mos, err := r.regionRepo.List()
	if err != nil {
		return regionDTOs, err
	}
	for _, mo := range mos {
		regionDTOs = append(regionDTOs, dto.Region{Region: mo})
	}
	return regionDTOs, err
}
func (r regionService) Page(num, size int) (page.Page, error) {
	var page page.Page
	var regionDTOs []dto.Region
	total, mos, err := r.regionRepo.Page(num, size)
	if err != nil {
		return page, err
	}
	for _, mo := range mos {
		regionDTO := new(dto.Region)
		m := make(map[string]interface{})
		regionDTO.Region = mo
		json.Unmarshal([]byte(mo.Vars), &m)
		regionDTO.RegionVars = m
		regionDTOs = append(regionDTOs, *regionDTO)
	}
	page.Total = total
	page.Items = regionDTOs
	return page, err
}

func (r regionService) Delete(name string) error {
	region, err := r.regionRepo.Get(name)
	if err != nil {
		return err
	}

	regions, err := r.zoneRepo.ListByRegionId(region.ID)
	if err != nil {
		return err
	}
	if len(regions) > 0 {
		errors.New(DeleteRegionError)
	}
	err = r.regionRepo.Delete(name)
	if err != nil {
		return err
	}
	return nil
}

func (r regionService) Create(creation dto.RegionCreate) (*dto.Region, error) {

	old, _ := r.Get(creation.Name)
	if old.ID != "" {
		return nil, errors.New(RegionNameExist)
	}

	vars, _ := json.Marshal(creation.RegionVars)
	region := model.Region{
		BaseModel:  common.BaseModel{},
		Name:       creation.Name,
		Vars:       string(vars),
		Datacenter: creation.Datacenter,
		Provider:   creation.Provider,
	}

	err := r.regionRepo.Save(&region)
	if err != nil {
		return nil, err
	}
	return &dto.Region{Region: region}, err
}

func (r regionService) Batch(op dto.RegionOp) error {
	var deleteItems []model.Region
	for _, item := range op.Items {
		deleteItems = append(deleteItems, model.Region{
			BaseModel: common.BaseModel{},
			ID:        item.ID,
			Name:      item.Name,
		})
	}
	err := r.regionRepo.Batch(op.Operation, deleteItems)
	if err != nil {
		return err
	}
	return nil
}

func (r regionService) ListDatacenter(creation dto.RegionDatacenterRequest) ([]string, error) {
	cloudClient := client.NewCloudClient(creation.RegionVars.(map[string]interface{}))
	var result []string
	if cloudClient != nil {
		result, err := cloudClient.ListDatacenter()
		if err != nil {
			return result, err
		}
		return result, err
	}
	return result, nil
}
