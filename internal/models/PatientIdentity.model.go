package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Patient struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FullName       string             `json:"full_name" bson:"full_name"`
	DateOfBirth    time.Time          `json:"date_of_birth" bson:"date_of_birth"`
	Gender         string             `json:"gender" bson:"gender"`
	Nationality    string             `json:"nationality" bson:"nationality"`
	ResidencePlace string             `json:"residence_place" bson:"residence_place"`
	Phone          string             `json:"phone" bson:"phone"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
	DetailedInfo   DetailedInfo       `json:"detailed_info" bson:"detailed_info"`
}

type DetailedInfo struct {
	ChipAuthen             string    `json:"chip_authen" bson:"chip_authen"`
	VerifySOD              string    `json:"verify_sod" bson:"verify_sod"`
	IDCode                 string    `json:"id_code" bson:"id_code"`
	OldIDCode              string    `json:"old_id_code" bson:"old_id_code"`
	IssueDate              time.Time `json:"issue_date" bson:"issue_date"`
	ExpiryDate             time.Time `json:"expiry_date" bson:"expiry_date"`
	OriginPlace            string    `json:"origin_place" bson:"origin_place"`
	PersonalIdentification string    `json:"personal_identification" bson:"personal_identification"`
	Race                   string    `json:"race" bson:"race"`
	Religion               string    `json:"religion" bson:"religion"`
	MotherName             string    `json:"mother_name" bson:"mother_name"`
	FatherName             string    `json:"father_name" bson:"father_name"`
	WifeName               string    `json:"wife_name" bson:"wife_name"`
	HumanName              string    `json:"human_name" bson:"human_name"`
}
