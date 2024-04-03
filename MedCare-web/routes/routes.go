package routes

import (
	"medcare/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	pwd, _ := os.Getwd()

	r.Static("/signin", pwd+"/frontend/signin")
	r.Static("/home", pwd+"/frontend/home")
	r.Static("/signup", pwd+"/frontend/signup")
	r.Static("/doctorsignup", pwd+"/frontend/doctorsignup")
	r.Static("/doctorsignin", pwd+"/frontend/doctorsignin")
	r.Static("/doctorhome", pwd+"/frontend/doctorhome")
	r.Static("/patienthome", pwd+"/frontend/patienthome")
	r.Static("/patientcreateappointment", pwd+"/frontend/patientcreateappointment")
	r.Static("/doctorenroll", pwd+"/frontend/doctorenroll")
	r.Static("/listappointmentfordoctor", pwd+"/frontend/doctorlistappointment")
	r.Static("/createprescriptionfordoctor", pwd+"/frontend/doctorprescription")
	r.Static("/patientreport", pwd+"/frontend/patientreport")
	r.Static("/doctorreport", pwd+"/frontend/doctorreport")
	r.Static("/patientprescriptionlist", pwd+"/frontend/patientprescription")

	r.POST("/signupcustomer", handlers.CustomerSignup)
	r.POST("/signincustomer", handlers.CustomerSignIn)
	r.POST("/signupdoctor", handlers.DoctorSignup)
	r.POST("/signindoctor", handlers.DoctorSignin)
	r.POST("/bookappointment", handlers.BookAppointment)
	r.POST("/listappointments", handlers.ListAppointment)
	r.POST("/createprescription", handlers.CreatePrescription)
	r.POST("/listpatientreport", handlers.ListPatientReport)
	r.POST("/listpatientprescription", handlers.ListPrescription)
}
