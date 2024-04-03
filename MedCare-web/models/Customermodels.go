package models

type CustomerSignUp struct {
	CustomerID  string `json:"customerid" bson:"customerid"`
	Name        string `json:"name" bson:"name" binding:"required"`
	EmailID     string `json:"emailid" bson:"emailid" binding:"required"`
	Password    string `json:"password" bson:"password" binding:"required"`
	CreatedTime string `json:"createdtime" bson:"createdtime"`
}

type CustomerSignIn struct {
	EmailId  string `json:"emailid" bson:"emailid" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type PatientToken struct {
	CustomerID string `json:"customerid" bson:"customerid"`
	EmailID    string `json:"emailid" bson:"emailid" binding:"required"`
	Token      string `json:"token" bson:"token" binding:"required"`
}

type BookAppointment struct {
	AppointmentID           string `json:"appointmentid" bson:"appointmentid"`
	Token                   string `json:"token,omitempty" bson:"token,omitempty"`
	CustomerID              string `json:"customerid,omitempty" bson:"customerid,omitempty"`
	Name                    string `json:"name" bson:"name" binding:"required"`
	DateOfBirth             string `json:"dateofbirth" bson:"dateofbirth" binding:"required"`
	Gender                  string `json:"gender" bson:"gender" binding:"required"`
	PhoneNumber             string `json:"phonenumber" bson:"phonenumber" binding:"required"`
	EmailID                 string `json:"emailid" bson:"emailid" binding:"required"`
	Address                 string `json:"address" bson:"address" binding:"required"`
	BriefDescription        string `json:"briefdescription" bson:"briefdescription" binding:"required"`
	Symptoms                string `json:"symptoms" bson:"symptoms" binding:"required"`
	Date                    string `json:"date" bson:"date" binding:"required"`
	FromDateTime            string `json:"fromdatetime" bson:"fromdatetime"`
	ToDateTime              string `json:"todatetime" bson:"todatetime"`
	ExistingConditions      string `json:"existingconditions" bson:"existingconditions" binding:"required"`
	Medications             string `json:"medications" bson:"medications" binding:"required"`
	PastSurgeriesTreatments string `json:"pastsurgeriestreatments" bson:"pastsurgeriestreatments" binding:"required"`
	EmergencyContactName    string `json:"emergencycontactname" bson:"emergencycontactname" binding:"required"`
	EmergencyPhoneNumber    string `json:"emergencycontactphonenumber" bson:"emergencycontactphonenumber" binding:"required"`
	Notes                   string `json:"notes" bson:"notes"`
	DoctorSpecialization    string `json:"doctorspecialization" bson:"doctorspecialization"`
	PreferredDoctorID       string `json:"preferreddoctorid" bson:"preferreddoctorid"`
	DoctorEmail             string `json:"doctoremail" bson:"doctoremail"`
	MeetLink                string `json:"meetlink" bson:"meetlink"`
	CreatedTime             string `json:"createdtime" bson:"createdtime"`
	DoctorName              string `json:"doctorname" bson:"doctorname"`
}

type ListReport struct {
	AppointmentID           string               `json:"appointmentid" bson:"appointmentid"`
	Token                   string               `json:"token,omitempty" bson:"token,omitempty"`
	CustomerID              string               `json:"customerid,omitempty" bson:"customerid,omitempty"`
	Name                    string               `json:"name" bson:"name"`
	DateOfBirth             string               `json:"dateofbirth" bson:"dateofbirth"`
	Gender                  string               `json:"gender" bson:"gender"`
	PhoneNumber             string               `json:"phonenumber" bson:"phonenumber"`
	EmailID                 string               `json:"emailid" bson:"emailid"`
	Address                 string               `json:"address" bson:"address"`
	BriefDescription        string               `json:"briefdescription" bson:"briefdescription"`
	Symptoms                string               `json:"symptoms" bson:"symptoms"`
	Date                    string               `json:"date" bson:"date"`
	FromDateTime            string               `json:"fromdatetime" bson:"fromdatetime"`
	ToDateTime              string               `json:"todatetime" bson:"todatetime"`
	ExistingConditions      string               `json:"existingconditions" bson:"existingconditions"`
	Medications             string               `json:"medications" bson:"medications"`
	PastSurgeriesTreatments string               `json:"pastsurgeriestreatments" bson:"pastsurgeriestreatments"`
	EmergencyContactName    string               `json:"emergencycontactname" bson:"emergencycontactname"`
	EmergencyPhoneNumber    string               `json:"emergencycontactphonenumber" bson:"emergencycontactphonenumber"`
	Notes                   string               `json:"notes" bson:"notes"`
	DoctorSpecialization    string               `json:"doctorspecialization" bson:"doctorspecialization"`
	PreferredDoctorID       string               `json:"preferreddoctorid" bson:"preferreddoctorid"`
	DoctorEmail             string               `json:"doctoremail" bson:"doctoremail"`
	Prescriptions           []CreatePrescription `json:"prescriptions" bson:"prescriptions"`
}

type ListPrescription struct {
	CustomerID      string       `json:"customerid" bson:"customerid"`
	Token           string       `json:"token,omitempty" bson:"token,omitempty"`
	AppointmentID   string       `json:"appointmentid" bson:"appointmentid"`
	DoctorID        string       `json:"doctorid" bson:"doctorid"`
	Instructions    string       `json:"instruction" bson:"instruction"`
	NextAppointment string       `json:"nextappointment" bson:"nextappointment"`
	Medication      []Medication `json:"medication" bson:"medication"`
	CreatedTime     string       `json:"createdtime" bson:"createdtime"`
}
