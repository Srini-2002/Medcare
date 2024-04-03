package handlers

import (
	"log"
	"medcare/constants"
	"medcare/models"
	"medcare/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomerSignup(c *gin.Context) {
	var request *models.CustomerSignUp
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error"})
		return
	}

	err = services.CustomerSignUp(request)
	if err != nil {
		log.Println("Error creating customer:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": "Account Created"})
}

func CustomerSignIn(c *gin.Context) {
	var request *models.CustomerSignIn
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": "Internal Server error"})
		return
	}
	status, ans := services.IsValidUser(request)
	if status {
		token, err := services.CreateToken(request.EmailId, ans.CustomerID)
		if err != nil {
			log.Println("Error creating token:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		response, err := services.InsertToken(ans.CustomerID, request.EmailId, token)
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

func BookAppointment(c *gin.Context) {
	var request *models.BookAppointment
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Error in Binding", err)
		return
	}
	token := request.Token
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	Customerid, err := services.ExtractID(token, constants.SecretKey)
	if err != nil {
		log.Printf("Error extracting CustomerID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	request.CustomerID = Customerid
	request.Token = ""
	err = services.BookAppointment(request)
	if err != nil {
		log.Println("error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Appointment Successfully Created"})

}

func ListPatientReport(c *gin.Context) {
	var request *models.ListReport
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	token := request.Token
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	Customerid, err := services.ExtractID(token, constants.SecretKey)
	if err != nil {
		log.Printf("Error extracting CustomerID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	request.CustomerID = Customerid
	request.Token = ""
	reports, err := services.ListPatientReport(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": reports})
}

func ListPrescription(c *gin.Context) {
	var request *models.ListPrescription
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	token := request.Token
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	Customerid, err := services.ExtractID(token, constants.SecretKey)
	if err != nil {
		log.Printf("Error extracting CustomerID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	request.CustomerID = Customerid
	request.Token = ""
	prescriptions, err := services.ListPrescription(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": prescriptions})
}
