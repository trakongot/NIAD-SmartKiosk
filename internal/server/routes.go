package server

import (
	"NIAD_SmartKiosk/internal/controller"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	PatientIdentityRouter := r.Group("/api/patient")
	{
		PatientIdentityRouter.POST("/add", controller.NewPatientController().AddPatient)
		PatientIdentityRouter.GET("/", controller.NewPatientController().GetOneUser)
	}

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	return r
}

// test server
func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
