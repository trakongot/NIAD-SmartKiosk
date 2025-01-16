package controller

import (
	"NIAD_SmartKiosk/internal/models"
	"NIAD_SmartKiosk/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PatientIdentityController struct {
	patientService *service.PatientIdentityService
}

func NewPatientIdentityController() *PatientIdentityController {
	return &PatientIdentityController{
		patientService: service.NewPatientIdentityService(),
	}
}

func (pc *PatientIdentityController) AddPatient(c *gin.Context) {
	var patientIdentity models.PatientIdentity
	if err := c.ShouldBindBodyWithJSON(&patientIdentity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := pc.patientService.Save(c.Request.Context(), &patientIdentity)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save patient identity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient identity created"})

}
