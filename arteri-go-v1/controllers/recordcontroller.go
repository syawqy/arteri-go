package controllers

import (
	"arteri-go-v1/database"
	"arteri-go-v1/entities"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var record entities.Record
	json.NewDecoder(r.Body).Decode(&record)
	database.Instance.Create(&record)
	json.NewEncoder(w).Encode(record)
}

func checkIfRecordExists(recordId string) bool {
	var record entities.Record
	database.Instance.First(&record, recordId)
	if record.ID == 0 {
		return false
	}
	return true
}

func GetRecordById(w http.ResponseWriter, r *http.Request) {
	recordId := mux.Vars(r)["id"]
	if checkIfCategoryExists(recordId) == false {
		json.NewEncoder(w).Encode("Record Not Found!")
		return
	}
	var record entities.Record
	database.Instance.First(&record, recordId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func GetRecords(w http.ResponseWriter, r *http.Request) {
	var records []entities.Record
	database.Instance.Find(&records)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(records)
}

func UpdateRecord(w http.ResponseWriter, r *http.Request) {
	recordId := mux.Vars(r)["id"]
	if checkIfCategoryExists(recordId) == false {
		json.NewEncoder(w).Encode("Record Not Found!")
		return
	}
	var record entities.Record
	database.Instance.First(&record, recordId)
	json.NewDecoder(r.Body).Decode(&record)
	database.Instance.Save(&record)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	recordId := mux.Vars(r)["id"]
	if checkIfCategoryExists(recordId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Record Not Found!")
		return
	}
	var record entities.Record
	database.Instance.Delete(&record, recordId)
	json.NewEncoder(w).Encode("Record Deleted Successfully!")
}
