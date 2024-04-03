package models

type DoctorSignup struct {
	DoctorID        string `json:"doctorid" bson:"doctorid"`
	Name            string `json:"name" bson:"name" binding:"required"`
	DateofBirth     string `json:"dateofbirth" bson:"dateofbirth" binding:"required"`
	Gender          string `json:"gender" bson:"gender" binding:"required"`
	EmailID         string `json:"emailid" bson:"emailid" binding:"required"`
	PhoneNumber     string `json:"phonenumber" bson:"phonenumber" binding:"required"`
	MedicalSchool   string `json:"medicalschool" bson:"medicalschool" binding:"required"`
	GraduationDate  string `json:"graduationdate" bson:"graduationdate" binding:"required"`
	DegreeDocument  string `json:"degreedocument" bson:"degreedocument" binding:"required"`
	LicenseNumber   string `json:"licensenumber" bson:"licensenumber" binding:"required"`
	LicenseDocument string `json:"licensedocument" bson:"licensedocument" binding:"required"`
	WorkExperience  string `json:"workexperience" bson:"workexperience" binding:"required"`
	Specialization  string `json:"specialization" bson:"specialization" binding:"required"`
	Photo           string `json:"photo" bson:"photo" binding:"required"`
	Password        string `json:"password" bson:"password" binding:"required"`
	CreatedTime     string `json:"createdtime" bson:"createdtime"`
	IsApproved      bool   `json:"isapproved" bson:"isapproved"`
}

type DoctorSignin struct {
	EmailId  string `json:"emailid" bson:"emailid" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type DoctorToken struct {
	DoctorID string `json:"doctorid" bson:"doctorid"`
	EmailID  string `json:"emailid" bson:"emailid" binding:"required"`
	Token    string `json:"token" bson:"token" binding:"required"`
}

type DoctorTimeline struct {
	DoctorID       string          `json:"doctorid" bson:"doctorid"`
	EmailID        string          `json:"emailid" bson:"emailid" `
	PhoneNumber    string          `json:"phonenumber" bson:"phonenumber"`
	Specialization string          `json:"specialization" bson:"specialization"`
	Name           string          `json:"name" bson:"name" binding:"required"`
	IsAvailable    bool            `json:"isavailable" bson:"isavailable"`
	IsApproved     bool            `json:"isapproved" bson:"isapproved"`
	CustomerCount  []CustomerCount `json:"customercount" bson:"customercount"`
}
type CustomerCount struct {
	Date             string `json:"date" bson:"date"`
	AppointmentCount int    `json:"appointmentcount" bson:"appointmentcount"`
}

type ListAppointmentforDoctor struct {
	Token    string `json:"token" bson:"token" binding:"required"`
	DoctorID string `json:"doctorid" bson:"doctorid"`
}

type CreatePrescription struct {
	CustomerId      string       `json:"customerid" bson:"customerid"`
	Token           string       `json:"token,omitempty" bson:"token,omitempty"`
	AppointmentID   string       `json:"appointmentid" bson:"appointmentid"`
	DoctorID        string       `json:"doctorid" bson:"doctorid"`
	Instructions    string       `json:"instruction" bson:"instruction"`
	NextAppointment string       `json:"nextappointment" bson:"nextappointment"`
	Medication      []Medication `json:"medication" bson:"medication"`
	CreatedTime     string       `json:"createdtime" bson:"createdtime"`
}
type Medication struct {
	Name      string `json:"name" bson:"name"`
	TimeRange string `json:"timerange" bson:"timerange"`
}
