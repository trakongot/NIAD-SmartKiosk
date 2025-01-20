package controller

import (
	"NIAD_SmartKiosk/internal/models"
	"NIAD_SmartKiosk/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PatientController struct {
	patientService *service.PatientService
}

func NewPatientController() *PatientController {
	return &PatientController{
		patientService: service.NewPatientService(),
	}
}

func (pc *PatientController) AddPatient(c *gin.Context) {
	var Patient models.Patient
	if err := c.ShouldBindBodyWithJSON(&Patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := pc.patientService.Save(c.Request.Context(), &Patient)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save patient identity"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Patient identity created"})
}

func (pc *PatientController) GetOneUser(c *gin.Context) {
	id := c.Query("id")
	patient, err := pc.patientService.FindOneByID(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Not found Patient",
		})
		return
	}

	c.JSON(http.StatusOK, patient)

}
