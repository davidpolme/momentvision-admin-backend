package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/davidpolme/momentvision-admin-backend/aws"
	"github.com/davidpolme/momentvision-admin-backend/models"
)

func CreateProgram(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateProgram")
	w.Header().Set("Content-Type", "application/json")
	var program models.Program
	err := json.NewDecoder(r.Body).Decode(&program)

	if err != nil {
		// Handle decoding error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the event data
	fmt.Printf("Received event: %+v\n", program)

	// Send a response back
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Event received successfully")

	//check if start date is before end date
	if program.StartDate > program.EndDate {
		log.Fatal("Start date is after end date")
	}


	if program.StartDate == program.EndDate {
	//check if start hour is before end hour
		if program.StartHour > program.EndHour {
			log.Fatal("Start hour is after end hour")
		}
	}

	// if start date and start time is before time now, set program to inactiiive
	// if start date and start time is after time now, set program to active
	program.Active = !(program.StartDate < time.Now().Format("2006-01-02") && program.StartHour < time.Now().Format("15:04") )

	log.Printf("Program: %+v\n", program)
	log.Printf("Program Active: %+v\n", program.Active)

	//convert program to json
	programJson, err := json.Marshal(program)
	if err != nil {
		log.Fatal(err)
	}
	
	//send program data to SNS
	aws.SendToSns(string(programJson))
}

func CreateHighlights(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var highlight models.Highlight
	_ = json.NewDecoder(r.Body).Decode(&highlight)
}
