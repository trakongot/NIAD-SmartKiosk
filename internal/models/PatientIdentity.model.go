package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PatientIdentity struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientID              string             `bson:"patient_id" json:"patient_id"`
	FullName               string             `bson:"full_name" json:"full_name"`
	DateOfBirth            string             `bson:"date_of_birth" json:"date_of_birth"`
	Gender                 string             `bson:"gender" json:"gender"`
	Nationality            string             `bson:"nationality" json:"nationality"`
	ResidencePlace         string             `bson:"residence_place" json:"residence_place"`
	ChipAuthen             string             `bson:"chip_authen" json:"chip_authen"`
	VerifySod              string             `bson:"verify_sod" json:"verify_sod"`
	IDCode                 string             `bson:"id_code" json:"id_code"`
	OldIDCode              string             `bson:"old_id_code" json:"old_id_code"`
	IssueDate              string             `bson:"issue_date" json:"issue_date"`
	ExpiryDate             string             `bson:"expiry_date" json:"expiry_date"`
	OriginPlace            string             `bson:"origin_place" json:"origin_place"`
	PersonalIdentification string             `bson:"personal_identification" json:"personal_identification"`
	Race                   string             `bson:"race" json:"race"`
	Religion               string             `bson:"religion" json:"religion"`
	MotherName             string             `bson:"mother_name" json:"mother_name"`
	FatherName             string             `bson:"father_name" json:"father_name"`
	WifeName               string             `bson:"wife_name" json:"wife_name"`
	HumanName              string             `bson:"human_name" json:"human_name"`
	CreatedAt              string             `bson:"created_at" json:"created_at"`
	UpdatedAt              string             `bson:"updated_at" json:"updated_at"`
}
