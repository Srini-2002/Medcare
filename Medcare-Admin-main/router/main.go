package router

import (
	"medcare-admin/controller"

	"github.com/gin-gonic/gin"
)

// Router creates and configures the Gin router.
func Router() *gin.Engine {
	router := gin.Default()

	router.Static("/admin", "./frontend/admin")

	router.POST("/getalldoctors", controller.GetAllDoctors)
	router.POST("/getapproveddoctors", controller.GetApprovedDoctors)
	router.POST("/getnotapproveddoctors", controller.GetNotApprovedDoctors)
	router.POST("/getdoctor", controller.GetDoctor)
	router.POST("/changestatus", controller.ChangeDoctorStatus)
	router.POST("/insertblog", controller.InsertBlog)

	return router
}
