package controllers

import (
	"arteri-go-v1/database"
	"arteri-go-v1/entities"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var category entities.Category
	json.NewDecoder(r.Body).Decode(&category)
	database.Instance.Create(&category)
	json.NewEncoder(w).Encode(category)
}

func checkIfCategoryExists(categoryId string) bool {
	var category entities.Category
	database.Instance.First(&category, categoryId)
	return category.ID != 0
}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	categoryId := mux.Vars(r)["id"]
	if !checkIfCategoryExists(categoryId) {
		json.NewEncoder(w).Encode("Category Not Found!")
		return
	}
	var record entities.Category
	database.Instance.First(&record, categoryId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var category []entities.Category
	database.Instance.Find(&category)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := mux.Vars(r)["id"]
	if !checkIfCategoryExists(categoryId) {
		json.NewEncoder(w).Encode("Category Not Found!")
		return
	}
	var record entities.Category
	database.Instance.First(&record, categoryId)
	json.NewDecoder(r.Body).Decode(&record)
	database.Instance.Save(&record)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categoryId := mux.Vars(r)["id"]
	if !checkIfCategoryExists(categoryId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Category Not Found!")
		return
	}
	var record entities.Category
	database.Instance.Delete(&record, categoryId)
	json.NewEncoder(w).Encode("Category Deleted Successfully!")
}
