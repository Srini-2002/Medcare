package database

import (
	"context"
	"medcare/constants"

	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Customer_Collection *mongo.Collection
var Customer_Token *mongo.Collection
var Doctor_Collection *mongo.Collection
var Doctor_Token *mongo.Collection
var BookAppointment *mongo.Collection

func init() {
	clientoption := options.Client().ApplyURI(constants.MongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected-(patient)")
	Customer_Collection = client.Database("MedCare").Collection("CustomerProfile")

	fmt.Println("CustomerProfile Connected")
}

func init() {
	clientoption := options.Client().ApplyURI(constants.MongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected-(patient)")
	Customer_Token = client.Database("MedCare").Collection("CustomerProfile-JWT-Token")

	fmt.Println("CustomerProfile-JWT-Token Connected")
}

func init() {
	clientoption := options.Client().ApplyURI(constants.MongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected-(patient)")
	BookAppointment = client.Database("MedCare").Collection("Patient-BookAppointment")

	fmt.Println("Patient-BookAppointment Connected")
}

// =================================================================
// --------------->   Doctor   <----------------------
// =================================================================

func init() {
	clientoption := options.Client().ApplyURI(constants.MongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected-(doctor)")
	Doctor_Collection = client.Database("MedCare").Collection("DoctorProfile")

	fmt.Println("DoctorProfile Connected")
}

func init() {
	clientoption := options.Client().ApplyURI(constants.MongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected-(doctor)")
	Doctor_Token = client.Database("MedCare").Collection("DoctorProfile-JWT-Token")

	fmt.Println("DoctorProfile-JWT-Token Connected")
}

var DoctorAppointmentSchedule_Collection *mongo.Collection

func init() {
	clientoption := options.Client().ApplyURI(constants.MongoURI)

	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected")
	DoctorAppointmentSchedule_Collection = client.Database("MedCare").Collection("DoctorAppointmentSchedule")

	fmt.Println("DoctorAppointmentSchedule_Collection Connected")
}

var Prescription_Collection *mongo.Collection

func init() {
	clientoption := options.Client().ApplyURI(constants.MongoURI)

	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected")
	Prescription_Collection = client.Database("MedCare").Collection("Prescription")

	fmt.Println("Prescription_Collection Connected")
}
