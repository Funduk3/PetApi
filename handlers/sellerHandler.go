package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"petstore-api/models"
	"petstore-api/services"

	"github.com/gorilla/mux"
)

type SellerHandler struct {
	service services.SellerService
}

func NewSellerHandler(service services.SellerService) *SellerHandler {
	return &SellerHandler{service: service}
}

// GetSellers godoc
// @Summary Get all sellers
// @Description Get list of all sellers with optional pets inclusion
// @Tags sellers
// @Accept json
// @Produce json
// @Param include_pets query bool false "Include pets in response"
// @Success 200 {object} Response{data=[]models.Seller}
// @Failure 500 {object} Response
// @Router /sellers [get]
func (h *SellerHandler) GetSellers(w http.ResponseWriter, r *http.Request) {
	includePets := r.URL.Query().Get("include_pets") == "true"

	sellers, err := h.service.GetAllSellers(includePets)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to fetch sellers")
		return
	}

	SendSuccessResponse(w, sellers, "")
}

// GetSeller godoc
// @Summary Get seller by ID
// @Description Get a single seller by ID with optional pets inclusion
// @Tags sellers
// @Accept json
// @Produce json
// @Param id path int true "Seller ID"
// @Param include_pets query bool false "Include pets in response"
// @Success 200 {object} Response{data=models.Seller}
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /sellers/{id} [get]
func (h *SellerHandler) GetSeller(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid seller ID")
		return
	}

	includePets := r.URL.Query().Get("include_pets") == "true"

	seller, err := h.service.GetSellerByID(uint(id), includePets)
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

// CreateSeller godoc
// @Summary Create a new seller
// @Description Create a new seller with the provided information
// @Tags sellers
// @Accept json
// @Produce json
// @Param seller body models.CreateSellerRequest true "Seller creation data"
// @Success 201 {object} Response{data=models.Seller}
// @Failure 400 {object} Response
// @Router /sellers [post]
func (h *SellerHandler) CreateSeller(w http.ResponseWriter, r *http.Request) {
	var req models.CreateSellerRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	seller, err := h.service.CreateSeller(&req)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	SendCreatedResponse(w, seller, "Seller created successfully")
}

// UpdateSeller godoc
// @Summary Update seller
// @Description Update an existing seller's information
// @Tags sellers
// @Accept json
// @Produce json
// @Param id path int true "Seller ID"
// @Param seller body models.UpdateSellerRequest true "Seller update data"
// @Success 200 {object} Response{data=models.Seller}
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /sellers/{id} [put]
func (h *SellerHandler) UpdateSeller(w http.ResponseWriter, r *http.Request) {
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

// DeleteSeller godoc
// @Summary Delete seller
// @Description Delete a seller by ID (only if no pets are associated)
// @Tags sellers
// @Accept json
// @Produce json
// @Param id path int true "Seller ID"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 409 {object} Response
// @Failure 500 {object} Response
// @Router /sellers/{id} [delete]
func (h *SellerHandler) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid seller ID")
		return
	}

	err = h.service.DeleteSeller(uint(id))
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
