package handlers

import (
	"cashier/models"
	"cashier/services"
	"cashier/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type ProductsHandler struct {
	service *services.ProductsService
}

func NewProductsHandler(service *services.ProductsService) *ProductsHandler {
	return &ProductsHandler{service: service}
}

func (h *ProductsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, products)
}

func (h *ProductsHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "ID tidak valid, gunakan angka")
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, product)
}

func (h *ProductsHandler) PostProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Products
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Periksa kembali data yang dikirim")
		return
	}

	err = h.service.Create(&product)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, product)
}

func (h *ProductsHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "ID tidak valid, gunakan angka")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONMessageResponse(w, http.StatusOK, "Product deleted successfully")
}

func (h *ProductsHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "ID tidak valid, gunakan angka")
		return
	}

	var product models.Products
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Periksa kembali data yang dikirim")
		return
	}

	product.ID = id
	err = h.service.Update(&product)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, product)
}
