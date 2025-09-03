package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"petstore-api/models"
	"petstore-api/services"
	"strconv"
)

type BuyerHandler struct {
	service services.SellerService
}

func NewBuyerHandler(service services.SellerService) *BuyerHandler {
	return &BuyerHandler{service: service}
}

func (h *BuyerHandler) GetBuyers(w *http.ResponseWriter, r *http.Request) {
	sellers, err := h.service.GetAllSellers()
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to fetch sellers")
		return
	}

	SendSuccessResponse(w, sellers, "")
}

func (h *BuyerHandler) GetBuyer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid seller ID")
		return
	}

	seller, err := h.service.GetSellerByID(uint(id))
	if err != nil {
		if err.Error() == "seller not found" {
			SendErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to fetch seller")
		return
	}

	SendSuccessResponse(w, seller, "")
}

func (h *BuyerHandler) CreateBuyer(w http.ResponseWriter, r *http.Request) {
	var req models.CreateSellerRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	seller, err := h.service.CreateBuyer(&req)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	SendCreatedResponse(w, seller, "Seller created successfully")
}

func (h *BuyerHandler) UpdateBuyer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid seller ID")
		return
	}

	var req models.UpdateSellerRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	seller, err := h.service.UpdateSeller(uint(id), &req)
	if err != nil {
		if err.Error() == "seller not found" {
			SendErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, seller, "Seller updated successfully")
}

func (h *BuyerHandler) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid seller ID")
		return
	}

	err = h.service.DeleteBuyer(uint(id))
	if err != nil {
		if err.Error() == "seller not found" {
			SendErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		if err.Error() == "cannot delete seller with existing pets" {
			SendErrorResponse(w, http.StatusConflict, err.Error())
			return
		}
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, nil, "Seller deleted successfully")
}
