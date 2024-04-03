package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"medcare-admin/config"
	"medcare-admin/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GenerateID(request string) string {
	length := 4
	randomBytes := make([]byte, length)
	// Use crypto/rand to generate random bytes
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err) // Handle error appropriately
	}
	// Convert random bytes to a hexadecimal string
	id := hex.EncodeToString(randomBytes)

	return request + id
}

func CurrentTime() string {
	currentTime := time.Now()
	currentTimeStr := currentTime.Format("2006-01-02 15:04:05")
	return currentTimeStr
}

func GetAllDoctors() ([]models.DoctorSignup, error) {
	var Doctors []models.DoctorSignup
	cursor, err := config.Doctor_Collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var Doctor models.DoctorSignup
		err := cursor.Decode(&Doctor)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		Doctors = append(Doctors, Doctor)
	}
	return Doctors, nil
}
func GetApprovedDoctors() ([]models.DoctorSignup, error) {
	var Doctors []models.DoctorSignup
	cursor, err := config.Doctor_Collection.Find(context.Background(), bson.M{"isapproved": true})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var Doctor models.DoctorSignup
		err := cursor.Decode(&Doctor)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		Doctors = append(Doctors, Doctor)
	}
	return Doctors, nil
}

func GetNotApprovedDoctors() ([]models.DoctorSignup, error) {
	var Doctors []models.DoctorSignup
	cursor, err := config.Doctor_Collection.Find(context.Background(), bson.M{"isapproved": false})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var Doctor models.DoctorSignup
		err := cursor.Decode(&Doctor)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		Doctors = append(Doctors, Doctor)
	}
	return Doctors, nil
}

func GetDoctor(id string) (models.DoctorSignup, error) {
	var Doctor models.DoctorSignup
	err := config.Doctor_Collection.FindOne(context.Background(), bson.M{"doctorid": id}).Decode(&Doctor)
	if err != nil {
		log.Println(err)
		return Doctor, err
	}

	return Doctor, nil
}

func UpdateDoctorStatus(data models.ChangeDoctorStatus) (string, error) {
	filter := bson.M{"doctorid": data.DoctorID}
	if data.Status == "approved" {
		update := bson.M{"$set": bson.M{"isapproved": true}}
		_, err := config.Doctor_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "", err
		}
		var Doctor models.DoctorSignup
		err = config.Doctor_Collection.FindOne(context.Background(), bson.M{"doctorid": data.DoctorID}).Decode(&Doctor)
		if err != nil {
			log.Println(err)
			return "", err
		}
		currentDate := time.Now().Format("2006-01-02")
		docappointmentsch := models.DoctorTimeline{
			DoctorID:       Doctor.DoctorID,
			EmailID:        Doctor.EmailID,
			PhoneNumber:    Doctor.PhoneNumber,
			Specialization: Doctor.Specialization,
			Name:           Doctor.Name,
			IsAvailable:    true,
			IsApproved:     Doctor.IsApproved,
		}
		docappointmentsch.CustomerCount = append(docappointmentsch.CustomerCount, models.CustomerCount{
			Date:  currentDate,
			Count: 0,
		})
		_, err = config.DoctorAppointmentSchedule_Collection.InsertOne(context.Background(), docappointmentsch)
		if err != nil {
			log.Println(err)
			return "", err
		}
		return "Approved Successfully", nil
	} else if data.Status == "declined" {
		_, err := config.Doctor_Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return "", err
		}
		return "Deleted Successfully", nil
	}
	return "Wrong Input", nil
}

func InsertBlog(request models.InsertBlog) error {
	ctx := context.Background()
	request.BlogID = GenerateID("BLG")
	request.CreatedTime = CurrentTime()
	res, err := config.BlogCollection.InsertOne(ctx, request)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Blog Inserted Id: ", res.InsertedID)

	return nil
}



