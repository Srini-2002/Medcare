package services

import (
	"context"
	"errors"
	"log"
	"medcare/database"
	"medcare/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DoctorSignup(request *models.DoctorSignup) error {
	ctx := context.Background()
	request.DoctorID = GenerateIDDOC()
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"doctorid", request.DoctorID}},
			bson.D{{"email", request.EmailID}},
		}},
	}

	existingdoctor := models.DoctorSignup{}
	err := database.Doctor_Collection.FindOne(ctx, filter).Decode(&existingdoctor)
	if err == nil {
		log.Println("Account Already Exists", err)
		return errors.New("account elready Exists")
	} else if err != mongo.ErrNoDocuments {
		return err
	}
	request.CreatedTime = CurrentTime()

	result, err := database.Doctor_Collection.InsertOne(ctx, request)
	if err != nil {
		log.Printf("Failed to create customer: %v", err)
		return errors.New("failed to create customer")
	}
	// Find and return the newly inserted customer
	newdoctor := &models.DoctorSignup{}
	err = database.Doctor_Collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(newdoctor)
	if err != nil {
		return err
	}

	return nil
}

func IsValidUserDoctor(request *models.DoctorSignin) (bool, models.DoctorSignup) {
	ctx := context.Background()
	doctor := models.DoctorSignup{}
	query := bson.M{"emailid": request.EmailId}
	err := database.Doctor_Collection.FindOne(ctx, query).Decode(&doctor)
	if err != nil {
		log.Println("Error in fetching user details : ", err)
		return false, doctor
	}
	if request.Password != doctor.Password {
		return false, doctor
	}
	if !doctor.IsApproved {
		return false, doctor
	}
	return true, doctor
}

func InsertTokenDoctor(doctorid, email, token string) (string, error) {
	dbToken := models.DoctorToken{EmailID: email, Token: token, DoctorID: doctorid}
	result, err := database.Doctor_Token.InsertOne(context.Background(), dbToken)
	if err != nil {
		log.Printf("Error inserting token: %v", err)
		return "", err
	}
	token1 := models.DoctorToken{}
	query := bson.M{"_id": result.InsertedID}
	err = database.Doctor_Token.FindOne(context.Background(), query).Decode(&token1)
	if err != nil {
		log.Println("Error in fetching token : ", err)
		return "", err
	}
	return token1.Token, nil
}

func ListAppointment(request *models.ListAppointmentforDoctor) ([]models.BookAppointment, error) {
	ctx := context.Background()
	currentTime := time.Now()
	currentDateStr := currentTime.Format("2006-01-02")
	currentTimeStr := currentTime.Format("15:04:05")
	filter := bson.M{
		"preferreddoctorid": request.DoctorID,
		"$or": []bson.M{
			// Appointments on the current date and starting from the current time onwards
			bson.M{"date": currentDateStr, "fromdatetime": bson.M{"$gte": currentTimeStr}},
			// Appointments on future dates
			bson.M{"date": bson.M{"$gt": currentDateStr}},
		},
	}

	cursor, err := database.BookAppointment.Find(ctx, filter)
	if err != nil {
		log.Println("error finding", err)
		return nil, err
	}
	var appointments []models.BookAppointment
	for cursor.Next(ctx) {
		var appointment models.BookAppointment
		err := cursor.Decode(&appointment)
		if err != nil {
			log.Println("error decoding", err)
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func CreatePrescription(request models.CreatePrescription) error {
	ctx := context.Background()
	request.Token = ""
	var appointment models.BookAppointment
	err := database.BookAppointment.FindOne(ctx, bson.M{"appointmentid": request.AppointmentID}).Decode(&appointment)
	if err != nil {
		log.Println(err)
		return err
	}
	request.CustomerId = appointment.CustomerID
	request.CreatedTime = CurrentTime()
	response, err := database.Prescription_Collection.InsertOne(ctx, request)
	if err != nil {
		log.Println("error inserting", err)
		return err
	}
	var Prescription *models.CreatePrescription
	err = database.Prescription_Collection.FindOne(ctx, bson.M{"_id": response.InsertedID}).Decode(&Prescription)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
