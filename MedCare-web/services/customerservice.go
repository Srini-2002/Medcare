package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	generativeai "medcare/GenerativeAI"
	"medcare/database"

	"medcare/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CustomerSignUp(request *models.CustomerSignUp) error {
	ctx := context.Background()
	request.CustomerID = GenerateIDPA()
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"customerid", request.CustomerID}},
			bson.D{{"email", request.EmailID}},
		}},
	}

	existingCustomer := models.CustomerSignUp{}
	err := database.Customer_Collection.FindOne(ctx, filter).Decode(&existingCustomer)
	if err == nil {
		log.Println("Account Already Exists", err)
		return errors.New("account elready Exists")
	} else if err != mongo.ErrNoDocuments {
		return err
	}
	request.CreatedTime = CurrentTime()

	result, err := database.Customer_Collection.InsertOne(ctx, request)
	if err != nil {
		log.Printf("Failed to create customer: %v", err)
		return errors.New("failed to create customer")
	}
	// Find and return the newly inserted customer
	newCustomer := &models.CustomerSignUp{}
	err = database.Customer_Collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(newCustomer)
	if err != nil {
		return err
	}

	return nil
}

func IsValidUser(request *models.CustomerSignIn) (bool, models.CustomerSignUp) {
	ctx := context.Background()
	customer := models.CustomerSignUp{}
	query := bson.M{"emailid": request.EmailId}
	err := database.Customer_Collection.FindOne(ctx, query).Decode(&customer)
	if err != nil {
		log.Println("Error in fetching user details : ", err)
		return false, customer
	}
	if request.Password != customer.Password {
		return false, customer
	}
	return true, customer
}

func InsertToken(customerid, email, token string) (string, error) {
	dbToken := models.PatientToken{EmailID: email, Token: token, CustomerID: customerid}
	result, err := database.Customer_Token.InsertOne(context.Background(), dbToken)
	if err != nil {
		log.Printf("Error inserting token: %v", err)
		return "", err
	}
	token1 := models.PatientToken{}
	query := bson.M{"_id": result.InsertedID}
	err = database.Customer_Token.FindOne(context.Background(), query).Decode(&token1)
	if err != nil {
		log.Println("Error in fetching token : ", err)
		return "", err
	}
	return token1.Token, nil
}

func BookAppointment(request *models.BookAppointment) error {
	ctx := context.Background()
	request.AppointmentID = GenerateIDBA()
	specialization, err := generativeai.GenerativeAI(request)
	if err != nil {
		log.Println(err)
		return err
	}

	doctor := models.DoctorTimeline{}
	filter := bson.M{
		"specialization": specialization,
		"isavailable":    true,
	}

	cursor, err := database.DoctorAppointmentSchedule_Collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var doctor models.DoctorTimeline
		if err := cursor.Decode(&doctor); err != nil {
			return err
		}

		dateFound := false
		for _, cc := range doctor.CustomerCount {
			if cc.Date == request.Date {
				dateFound = true
				break
			}
		}

		if !dateFound {
			doctor.CustomerCount = append(doctor.CustomerCount, models.CustomerCount{
				Date:             request.Date,
				AppointmentCount: 0,
			})

			_, err := database.DoctorAppointmentSchedule_Collection.ReplaceOne(ctx, bson.M{"doctorid": doctor.DoctorID}, doctor)
			if err != nil {
				return err
			}
		}
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	filter1 := bson.M{"isavailable": true}
	filter2 := bson.M{"specialization": specialization}
	filter3 := bson.M{
		"customercount": bson.M{
			"$elemMatch": bson.M{
				"date": request.Date,
				"appointmentcount": bson.M{
					"$lt": 10, // Less than 10
				},
			},
		},
	}

	combinedFilter := bson.M{
		"$and": []bson.M{filter1, filter2, filter3},
	}
	err = database.DoctorAppointmentSchedule_Collection.FindOne(ctx, combinedFilter).Decode(&doctor)
	if err != nil {
		log.Println("Not found Doctor with this specialization", err)
		err = database.Doctor_Collection.FindOne(ctx, bson.M{"specialization": "Public Health and General Preventive Medicine"}).Decode(&doctor)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	if !doctor.IsApproved {
		log.Println(doctor.Name + " is not approved yet")
		return errors.New("doctor not approved yet")
	}

	filter = bson.M{
		"doctorid": doctor.DoctorID,
		"customercount": bson.M{
			"$elemMatch": bson.M{
				"date": request.Date,
			},
		},
	}

	update := bson.M{
		"$inc": bson.M{
			"customercount.$.appointmentcount": 1, // Increment by 1
		},
	}

	_, err = database.DoctorAppointmentSchedule_Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	var count models.CustomerCount
	for _, cc := range doctor.CustomerCount {
		if cc.Date == request.Date {
			count = cc
			break
		}
	}

	if count.AppointmentCount == 0 {
		request.FromDateTime = "09:00:00"
		request.ToDateTime = "09:30:00"
	} else if count.AppointmentCount == 1 {
		request.FromDateTime = "09:40:00"
		request.ToDateTime = "10:10:00"
	} else if count.AppointmentCount == 2 {
		request.FromDateTime = "10:20:00"
		request.ToDateTime = "10:50:00"
	} else if count.AppointmentCount == 3 {
		request.FromDateTime = "11:00:00"
		request.ToDateTime = "11:30:00"
	} else if count.AppointmentCount == 4 {
		request.FromDateTime = "11:40:00"
		request.ToDateTime = "12:10:00"
	} else if count.AppointmentCount == 5 {
		request.FromDateTime = "13:40:00"
		request.ToDateTime = "14:10:00"
	} else if count.AppointmentCount == 6 {
		request.FromDateTime = "14:20:00"
		request.ToDateTime = "14:50:00"
	} else if count.AppointmentCount == 7 {
		request.FromDateTime = "15:00:00"
		request.ToDateTime = "15:30:00"
	} else if count.AppointmentCount == 8 {
		request.FromDateTime = "15:40:00"
		request.ToDateTime = "16:10:00"
	} else if count.AppointmentCount == 9 {
		request.FromDateTime = "16:20:00"
		request.ToDateTime = "16:50:00"
	} else {
		log.Println(err)
		return errors.New("invalid appointment count")
	}

	request.DoctorSpecialization = specialization
	request.PreferredDoctorID = doctor.DoctorID
	request.DoctorEmail = doctor.EmailID
	request.DoctorName = doctor.Name

	res, err, icsdata := GetMeetLink(request)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("meeetlink", res)
	request.MeetLink = res

	// _, err = database.BookAppointment.UpdateOne(ctx, bson.M{"_id": result1.InsertedID}, bson.M{"$set": bson.M{"meetlink": res}})

	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	err = SendEmailforAppointment(request, "appointment.ics", icsdata)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}
	request.CreatedTime = CurrentTime()
	result1, err := database.BookAppointment.InsertOne(ctx, request)
	if err != nil {
		log.Println("error while inserting  patient appointment", err)
		return err
	}

	err = database.BookAppointment.FindOne(ctx, bson.M{"_id": result1.InsertedID}).Decode(&request)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func ListPatientReport(request *models.ListReport) (*models.ListReport, error) {
	ctx := context.Background()
	filter := bson.M{"customerid": request.CustomerID}

	// Define the options to sort by CreatedTime in descending order
	var result models.ListReport
	opts := options.FindOne().SetSort(bson.D{{"createdtime", -1}})
	err := database.BookAppointment.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {

			return nil, fmt.Errorf("could not find")
		}
		return nil, err
	}

	var prescriptions []models.CreatePrescription
	cursor, err := database.Prescription_Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Use cursor.All to decode all documents into the prescriptions slice
	if err := cursor.All(ctx, &prescriptions); err != nil {
		return nil, err
	}

	result.Prescriptions = prescriptions

	return &result, nil
}

func ListPrescription(request *models.ListPrescription) ([]models.ListPrescription, error) {
	ctx := context.Background()
	filter := bson.M{"customerid": request.CustomerID}
	cursor, err := database.Prescription_Collection.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)
	var result []models.ListPrescription

	err = cursor.All(ctx, &result)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

