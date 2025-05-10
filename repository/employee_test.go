package repository

import (
	"context"
	"log"
	"os"
	"projectx/model"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
		mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in .env")
	}
	mongoTestClient, err := mongo.Connect(context.Background(),
	options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatal("error while connecting mongodb", err)
	}

	log.Println("mongodb Successfully connected.")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("ping failed", err)
	}
	log.Println("ping success")

	return mongoTestClient
}

func TestMongoOperation(t *testing.T) {
	mongoTestClient := newMongoClient()

	defer mongoTestClient.Disconnect(context.Background())

	// dummy data 
	emp1 := uuid.New().String()
	// emp2 := uuid.New().string()

	// connect to collection

	coll := mongoTestClient.Database("companydb").Collection("employee_test")

	empRepo := EmployeeRepo{MongoCollection: coll}

	// Insert Employee 1 data
	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := model.Employee{
			Name: 		"Tony Stark",
			Department: "Physics",
			Employee: emp1,
		}	

		result, err := empRepo.InsertEmployee(&emp)

		if err != nil {
			log.Fatal("Insert 1 operation failed", err)
		}
		t.Log("Insert 1 successful", result)
	})
	
	// get Employee 1 data

	t.Run("Get Employee 1", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)

		if err != nil {
			log.Fatal("get operation failed", err)
		}
		t.Log("emp 1", result.Name)
	})

	// get All Employee
	t.Run("Get Employee 1", func(t *testing.T) {
		results, err := empRepo.FindAllEmployee()

		if err != nil {
			log.Fatal("get operation failed", err)
		}
		t.Log("employees", results)
	})

	// Update Employee 1
	t.Run("update Employee 1 Name", func(t *testing.T) {
		emp := model.Employee{
			Name: 		"Tony Stark vs Iron Man",
			Department: "Physics",
			Employee: emp1,
		}	

		result, err := empRepo.UpdateEmployeeID(emp1, &emp)

		if err != nil {
			log.Fatal("Insert 1 operation failed", err)
		}
		t.Log("Insert 1 successful", result)
	})
	t.Run("Get Employee 1 after update", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)

		if err != nil {
			log.Fatal("get operation failed", err)
		}
		t.Log("emp 1", result.Name)
	})
}
