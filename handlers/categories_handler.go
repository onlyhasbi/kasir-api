package handlers

import (
	"cashier/models"
	"cashier/services"
	"cashier/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CategoriesHandler struct {
	service *services.CategoryService
}

func NewCategoriesHandler(service *services.CategoryService) *CategoriesHandler {
	return &CategoriesHandler{service: service}
}

func (h *CategoriesHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, http.StatusOK, categories)
}

func (h *CategoriesHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "ID tidak valid, gunakan angka")
		return
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, category)
}

func (h *CategoriesHandler) PostCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Categories
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Printf("JSON Decode Error: %v", err)
		utils.ErrorResponse(w, http.StatusBadRequest, "Periksa kembali data yang dikirim")
		return
	}

	err = h.service.Create(&category)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, category)
}

func (h *CategoriesHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "ID tidak valid, gunakan angka")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONMessageResponse(w, http.StatusOK, "Category deleted successfully")
}

func (h *CategoriesHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "ID tidak valid, gunakan angka")
		return
	}

	var category models.Categories
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Periksa kembali data yang dikirim")
		return
	}

	category.ID = id
	err = h.service.Update(&category)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, category)
}
