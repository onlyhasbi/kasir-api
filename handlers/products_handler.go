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

// GetAll godoc
// @Summary      Ambil semua produk
// @Description  Mengambil semua data produk dari database
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {object}   utils.DataResponse{data=[]models.Products}
// @Failure      500  {object}   utils.MessageResponse
// @Router       /api/product [get]
func (h *ProductsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, products)
}

// GetProduct godoc
// @Summary      Ambil produk berdasarkan ID
// @Description  Mengambil data produk spesifik berdasarkan ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  utils.DataResponse{data=models.Products}
// @Failure      400  {object}  utils.MessageResponse
// @Failure      404  {object}  utils.MessageResponse
// @Router       /api/product/{id} [get]
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

// PostProduct godoc
// @Summary      Buat produk baru
// @Description  Menambahkan data produk baru ke database
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product body models.Products true "Data Produk"
// @Success      201  {object}  utils.DataResponse{data=models.Products}
// @Failure      400  {object}  utils.MessageResponse
// @Router       /api/product [post]
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

// DeleteProduct godoc
// @Summary      Hapus produk
// @Description  Menghapus data produk berdasarkan ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  utils.MessageResponse
// @Failure      400  {object}  utils.MessageResponse
// @Failure      404  {object}  utils.MessageResponse
// @Router       /api/product/{id} [delete]
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

// UpdateProduct godoc
// @Summary      Update produk
// @Description  Memperbarui data produk berdasarkan ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Param        product body models.Products true "Data Produk"
// @Success      200  {object}  utils.DataResponse{data=models.Products}
// @Failure      400  {object}  utils.MessageResponse
// @Router       /api/product/{id} [put]
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
