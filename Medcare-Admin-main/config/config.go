package config

import (
	"context"
	"fmt"
	"log"
	"medcare-admin/constants"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Doctor_Collection *mongo.Collection

func init() {
	clientoption := options.Client().ApplyURI(constants.Connectstring)

	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected")
	Doctor_Collection = client.Database(constants.DB_Name).Collection(constants.Doctor_collection)

	fmt.Println("Doctor_Collection Connected")
}

var DoctorAppointmentSchedule_Collection *mongo.Collection

func init() {
	clientoption := options.Client().ApplyURI(constants.Connectstring)

	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected")
	DoctorAppointmentSchedule_Collection = client.Database(constants.DB_Name).Collection("DoctorAppointmentSchedule")

	fmt.Println("DoctorAppointmentSchedule_Collection Connected")
}


var BlogCollection *mongo.Collection

func init() {
	clientoption := options.Client().ApplyURI(constants.Connectstring)

	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected")
	BlogCollection = client.Database(constants.DB_Name).Collection("Medcare-Blogs")

	fmt.Println("BlogCollection Connected")
}