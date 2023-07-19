package services

import (
	"github.com/bheemeshkammak/compage_module/compage_module/pkg/rest/server/daos"
	"github.com/bheemeshkammak/compage_module/compage_module/pkg/rest/server/models"
)

type CbnService struct {
	cbnDao *daos.CbnDao
}

func NewCbnService() (*CbnService, error) {
	cbnDao, err := daos.NewCbnDao()
	if err != nil {
		return nil, err
	}
	return &CbnService{
		cbnDao: cbnDao,
	}, nil
}

func (cbnService *CbnService) CreateCbn(cbn *models.Cbn) (*models.Cbn, error) {
	return cbnService.cbnDao.CreateCbn(cbn)
}

func (cbnService *CbnService) UpdateCbn(id int64, cbn *models.Cbn) (*models.Cbn, error) {
	return cbnService.cbnDao.UpdateCbn(id, cbn)
}

func (cbnService *CbnService) DeleteCbn(id int64) error {
	return cbnService.cbnDao.DeleteCbn(id)
}

func (cbnService *CbnService) ListCbns() ([]*models.Cbn, error) {
	return cbnService.cbnDao.ListCbns()
}

func (cbnService *CbnService) GetCbn(id int64) (*models.Cbn, error) {
	return cbnService.cbnDao.GetCbn(id)
}
