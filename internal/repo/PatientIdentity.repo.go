package repo

import (
	"NIAD_SmartKiosk/internal/database"
	"NIAD_SmartKiosk/internal/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type PatientRepo struct {
	Patient *mongo.Collection
}

func NewPatientRepo() *PatientRepo {
	return &PatientRepo{
		Patient: database.GetCollection("Patient"),
	}
}

func (pr *PatientRepo) FindOne(ctx context.Context, filter interface{}) (*models.Patient, error) {
	var patient models.Patient

	err := pr.Patient.FindOne(ctx, filter).Decode(&patient)

	if err != nil {

		return nil, err
	}

	return &patient, nil
}

func (pr *PatientRepo) Save(ctx context.Context, patient *models.Patient) error {
	patient.CreatedAt = time.Now()
	patient.UpdatedAt = time.Now()

	_, err := pr.Patient.InsertOne(ctx, patient)
	if err != nil {
		log.Println("Error inserting document:", err)
		return err
	}

	return nil
}
