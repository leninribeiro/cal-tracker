package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/leninribeiro/cal-tracker/data"
	"github.com/leninribeiro/cal-tracker/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("name")
	filter := bson.M{}
	if userName != "" {
		filter = bson.M{"name": userName}
	}
	users, err := data.FindUser(filter)
	if err != nil {
		log.Printf("Error finding user: %s", err)
		failedJson, _ := json.Marshal(models.UserResponse{Success: "false"})
		rw.Write(failedJson)
	}
	response := models.UserResponse{
		Success: "true",
		Users:   users,
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Printf("Could not marshal json: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("INTERNAL SERVER ERROR"))
	}
	rw.Write(responseJson)
}
