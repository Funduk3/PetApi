package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"petstore-api/models"
	"petstore-api/services"

	"github.com/gorilla/mux"
)

type PetHandler struct {
	service services.PetService
}

func NewPetHandler(service services.PetService) *PetHandler {
	return &PetHandler{service: service}
}

// GetPets godoc
// @Summary Get all pets
// @Description Get list of all pets with optional seller inclusion and filtering
// @Tags pets
// @Accept json
// @Produce json
// @Param include_seller query bool false "Include seller information in response"
// @Param seller_id query int false "Filter pets by seller ID"
// @Success 200 {object} Response{data=[]models.Pet}
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /pets [get]
func (h *PetHandler) GetPets(w http.ResponseWriter, r *http.Request) {
	includeSeller := r.URL.Query().Get("include_seller") == "true"

	var sellerID *uint
	if sellerIDStr := r.URL.Query().Get("seller_id"); sellerIDStr != "" {
		id, err := strconv.Atoi(sellerIDStr)
		if err != nil {
			SendErrorResponse(w, http.StatusBadRequest, "Invalid seller ID")
			return
		}
		sellerIDUint := uint(id)
		sellerID = &sellerIDUint
	}

	pets, err := h.service.GetAllPets(includeSeller, sellerID)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to fetch pets")
		return
	}

	SendSuccessResponse(w, pets, "")
}

// GetPet godoc
// @Summary Get pet by ID
// @Description Get a single pet by ID with optional seller inclusion
// @Tags pets
// @Accept json
// @Produce json
// @Param id path int true "Pet ID"
// @Param include_seller query bool false "Include seller information in response"
// @Success 200 {object} Response{data=models.Pet}
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /pets/{id} [get]
func (h *PetHandler) GetPet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid pet ID")
		return
	}

	includeSeller := r.URL.Query().Get("include_seller") == "true"

	pet, err := h.service.GetPetByID(uint(id), includeSeller)
	if err != nil {
		if err.Error() == "pet not found" {
			SendErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to fetch pet")
		return
	}

	SendSuccessResponse(w, pet, "")
}

// CreatePet godoc
// @Summary Create a new pet
// @Description Create a new pet with the provided information
// @Tags pets
// @Accept json
// @Produce json
// @Param pet body models.CreatePetRequest true "Pet creation data"
// @Success 201 {object} Response{data=models.Pet}
// @Failure 400 {object} Response
// @Router /pets [post]
func (h *PetHandler) CreatePet(w http.ResponseWriter, r *http.Request) {
	var req models.CreatePetRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	pet, err := h.service.CreatePet(&req)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	SendCreatedResponse(w, pet, "Pet created successfully")
}

// UpdatePet godoc
// @Summary Update pet
// @Description Update an existing pet's information
// @Tags pets
// @Accept json
// @Produce json
// @Param id path int true "Pet ID"
// @Param pet body models.UpdatePetRequest true "Pet update data"
// @Success 200 {object} Response{data=models.Pet}
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /pets/{id} [put]
func (h *PetHandler) UpdatePet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid pet ID")
		return
	}

	var req models.UpdatePetRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	pet, err := h.service.UpdatePet(uint(id), &req)
	if err != nil {
		if err.Error() == "pet not found" {
			SendErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, pet, "Pet updated successfully")
}

// DeletePet godoc
// @Summary Delete pet
// @Description Delete a pet by ID
// @Tags pets
// @Accept json
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /pets/{id} [delete]
func (h *PetHandler) DeletePet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid pet ID")
		return
	}

	err = h.service.DeletePet(uint(id))
	if err != nil {
		if err.Error() == "pet not found" {
			SendErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, nil, "Pet deleted successfully")
}
