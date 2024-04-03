package controller

import (
	"log"
	"medcare-admin/models"
	"medcare-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDoctors(c *gin.Context) {
	Doctors, err := service.GetAllDoctors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": Doctors})

}
func GetApprovedDoctors(c *gin.Context) {
	Doctors, err := service.GetApprovedDoctors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": Doctors})

}
func GetNotApprovedDoctors(c *gin.Context) {
	Doctors, err := service.GetNotApprovedDoctors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": Doctors})

}

func GetDoctor(c *gin.Context) {
	var id models.GetDoctor
	if err := c.ShouldBindJSON(&id); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	Doctor, err := service.GetDoctor(id.DoctorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"message": Doctor})

}

func ChangeDoctorStatus(c *gin.Context) {
	var data models.ChangeDoctorStatus
	if err := c.ShouldBindJSON(&data); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	UpdatedStatus, err := service.UpdateDoctorStatus(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"message": UpdatedStatus})
}

func InsertBlog(c *gin.Context) {
	var request models.InsertBlog
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
		return
	}
	err = service.InsertBlog(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert blog"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully inserted the blog"})
}
