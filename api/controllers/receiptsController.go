package controllers

import (
	"encoding/json"
	"net/http"

	"main/internal/data"
	"main/internal/model"
	"main/internal/service"
)


func InitReceiptsController() {
	http.Handle("POST /receipts/process", http.HandlerFunc(ProcessReceipt))
	http.Handle("GET /receipts/{id}/points", http.HandlerFunc(GetPoints))
}

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var req model.ProcessReceiptRequest

	//Validate request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(),http.StatusBadRequest)
	} else {
		
	//Process request
		res := service.ProcessReceipt(req)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
	
}

func GetPoints(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Invalid Id",http.StatusBadRequest)
	} else {
		res:= model.GetPointsResponse{
			Points: data.GetDataById(id),
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}