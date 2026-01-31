package handlers

import (
	"cashier/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Categories)
	return
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID tidak valid, gunakan angka", http.StatusBadRequest)
		return
	}

	for _, c := range models.Categories {
		if c.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	http.Error(w, "Category tidak ditemukan.", http.StatusNotFound)
}

func PostCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Periksa kembali data yang dikirim", http.StatusBadRequest)
		return
	}

	newCategory.ID = len(models.Categories) + 1
	models.Categories = append(models.Categories, newCategory)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID tidak valid, gunakan angka", http.StatusBadRequest)
		return
	}

	for i, c := range models.Categories {
		if c.ID == id {
			models.Categories = append(models.Categories[:i], models.Categories[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Kategori berhasil dihapus",
			})
			return
		}
	}

	http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID tidak valid, gunakan angka", http.StatusBadRequest)
		return
	}

	var updateCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)

	if err != nil {
		http.Error(w, "Periksa kembali data yang dikirim", http.StatusBadRequest)
		return
	}

	for i, c := range models.Categories {
		if c.ID == id {
			updateCategory.ID = id
			models.Categories[i] = updateCategory

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Kategori berhasil diperbarui",
			})
			return
		}
	}

	http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
}
