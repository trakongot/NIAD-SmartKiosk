package service

import (
	"NIAD_SmartKiosk/internal/models"
	"NIAD_SmartKiosk/internal/repo"
	"context"
)

type PatientIdentityService struct {
	patientRepo *repo.PatientIdentityRepo
}

func NewPatientIdentityService() *PatientIdentityService {
	return &PatientIdentityService{
		patientRepo: repo.NewPatientIdentityRepo(),
	}
}

func (ps *PatientIdentityService) Save(ctx context.Context, patientIdentity *models.PatientIdentity) error {

	return ps.patientRepo.Save(ctx, patientIdentity)
}
