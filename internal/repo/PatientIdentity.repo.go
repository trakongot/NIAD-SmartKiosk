package repo

import (
	"NIAD_SmartKiosk/internal/database"
	"NIAD_SmartKiosk/internal/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type PatientIdentityRepo struct {
	db *mongo.Client
}

func NewPatientIdentityRepo() *PatientIdentityRepo {
	return &PatientIdentityRepo{
		db: database.NewDB(),
	}
}

func (pr *PatientIdentityRepo) Save(ctx context.Context, patientIdentity *models.PatientIdentity) error {
	collection := pr.db.Database("NIAD").Collection("PatientIdentity")

	_, err := collection.InsertOne(ctx, patientIdentity)
	if err != nil {
		log.Println("Error inserting document:", err)
		return err
	}

	return nil
}
