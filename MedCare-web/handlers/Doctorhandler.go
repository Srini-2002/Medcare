package handlers

import (
	"log"
	"medcare/constants"
	"medcare/models"
	"medcare/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoctorSignup(c *gin.Context) {
	var request *models.DoctorSignup
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error"})
		return
	}
	request.IsApproved = false

	err = services.DoctorSignup(request)
	if err != nil {
		log.Println("Error creating customer:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": "Account Created"})
}

func DoctorSignin(c *gin.Context) {
	var request *models.DoctorSignin
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": "Internal Server error"})
		return
	}
	status, ans := services.IsValidUserDoctor(request)
	if status {
		token, err := services.CreateToken(request.EmailId, ans.DoctorID)
		if err != nil {
			log.Println("Error creating token:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		response, err := services.InsertTokenDoctor(ans.DoctorID, request.EmailId, token)
		if err != nil {
			log.Println("Error inserting token:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": response})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
	}
}

func ListAppointment(c *gin.Context) {
	var request *models.ListAppointmentforDoctor
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	token := request.Token
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	Doctorid, err := services.ExtractID(token, constants.SecretKey)
	if err != nil {
		log.Printf("Error extracting CustomerID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	request.DoctorID = Doctorid
	request.Token = ""

	response, err := services.ListAppointment(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": response})

}

func CreatePrescription(c *gin.Context) {
	var request models.CreatePrescription
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	token := request.Token
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	Doctorid, err := services.ExtractID(token, constants.SecretKey)
	if err != nil {
		log.Printf("Error extracting CustomerID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	request.DoctorID = Doctorid
	request.Token = ""
	err = services.CreatePrescription(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Inserted Successfully"})

}
