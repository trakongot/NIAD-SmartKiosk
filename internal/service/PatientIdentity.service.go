package service

import (
	"NIAD_SmartKiosk/internal/models"
	"NIAD_SmartKiosk/internal/repo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PatientService struct {
	patientRepo *repo.PatientRepo
}

func NewPatientService() *PatientService {
	return &PatientService{
		patientRepo: repo.NewPatientRepo(),
	}
}

func (ps *PatientService) Save(ctx context.Context, patient *models.Patient) error {

	return ps.patientRepo.Save(ctx, patient)
}

func (ps *PatientService) FindOneByID(ctx context.Context, id string) (*models.Patient, error) {
	// Chuyển ID từ string sang ObjectID
	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objectID}

	patient, err := ps.patientRepo.FindOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	return patient, nil
}
