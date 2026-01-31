package handlers

import (
	"cashier/models"
	"cashier/services"
	"cashier/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type CategoriesHandler struct {
	service *services.CategoryService
}

func NewCategoriesHandler(service *services.CategoryService) *CategoriesHandler {
	return &CategoriesHandler{service: service}
}

// GetAll godoc
// @Summary      Ambil semua kategori
// @Description  Mengambil semua data kategori dari database
// @Tags         categories
// @Accept       json
// @Produce      json
// @Success      200  {object}   utils.DataResponse{data=[]models.Categories}
// @Failure      500  {object}   utils.MessageResponse
// @Router       /api/category [get]
func (h *CategoriesHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, http.StatusOK, categories)
}

// GetCategory godoc
// @Summary      Ambil kategori berdasarkan ID
// @Description  Mengambil data kategori spesifik berdasarkan ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Success      200  {object}  utils.DataResponse{data=models.Categories}
// @Failure      400  {object}  utils.MessageResponse
// @Failure      404  {object}  utils.MessageResponse
// @Router       /api/category/{id} [get]
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

// PostCategory godoc
// @Summary      Buat kategori baru
// @Description  Menambahkan data kategori baru ke database
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category body models.Categories true "Data Kategori"
// @Success      201  {object}  utils.DataResponse{data=models.Categories}
// @Failure      400  {object}  utils.MessageResponse
// @Router       /api/category [post]
func (h *CategoriesHandler) PostCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Categories
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
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

// DeleteCategory godoc
// @Summary      Hapus kategori
// @Description  Menghapus data kategori berdasarkan ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Success      200  {object}  utils.MessageResponse
// @Failure      400  {object}  utils.MessageResponse
// @Failure      500  {object}  utils.MessageResponse
// @Router       /api/category/{id} [delete]
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

// UpdateCategory godoc
// @Summary      Update kategori
// @Description  Memperbarui data kategori berdasarkan ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Param        category body models.Categories true "Data Kategori"
// @Success      200  {object}  utils.DataResponse{data=models.Categories}
// @Failure      400  {object}  utils.MessageResponse
// @Router       /api/category/{id} [put]
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
