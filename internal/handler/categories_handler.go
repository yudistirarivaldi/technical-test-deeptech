package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/dto"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/model"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/service"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/utils"
)

type CategoriesHandler struct {
	categoriesService *service.CategoriesService
}

func NewCategoriesHandler(categoriesService *service.CategoriesService) *CategoriesHandler {
	return &CategoriesHandler{
		categoriesService: categoriesService,
	}
}

func (h *CategoriesHandler) HandleInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, model.Response{
			ResponseCode: "01",
			Message:      "Method not allowed",
		})
		return
	}

	var req dto.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid request body",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Validation failed",
			Errors:       utils.FormatValidationErrors(err),
		})
		return
	}

	category := &model.Categories{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.categoriesService.InsertCategory(r.Context(), category); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to insert category",
		})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, model.Response{
		ResponseCode: "00",
		Message:      "Category created successfully",
	})
}

func (h *CategoriesHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, model.Response{
			ResponseCode: "01",
			Message:      "Method not allowed",
		})
		return
	}

	data, err := h.categoriesService.GetAll(r.Context())
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to get categories",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "Success",
		Data:         data,
	})
}

func (h *CategoriesHandler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, model.Response{
			ResponseCode: "01",
			Message:      "Method not allowed",
		})
		return
	}

	vars := mux.Vars(r)
	idParam := vars["id"]
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id <= 0 {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid category ID",
		})
		return
	}

	category, err := h.categoriesService.GetByID(r.Context(), id)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to get category",
		})
		return
	}
	if category == nil {
		utils.WriteJSON(w, http.StatusNotFound, model.Response{
			ResponseCode: "01",
			Message:      "Category not found",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "Success",
		Data:         category,
	})
}

func (h *CategoriesHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, model.Response{
			ResponseCode: "01",
			Message:      "Method not allowed",
		})
		return
	}

	vars := mux.Vars(r)
	idParam := vars["id"]
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id <= 0 {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid category ID",
		})
		return
	}

	var req dto.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid request body",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Validation failed",
			Errors:       utils.FormatValidationErrors(err),
		})
		return
	}

	category := &model.Categories{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.categoriesService.UpdateCategory(r.Context(), category); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to update category",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "Category updated successfully",
	})
}

func (h *CategoriesHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, model.Response{
			ResponseCode: "01",
			Message:      "Method not allowed",
		})
		return
	}

	vars := mux.Vars(r)
	idParam := vars["id"]
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id <= 0 {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid category ID",
		})
		return
	}

	if err := h.categoriesService.DeleteCategory(r.Context(), id); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to delete category",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "Category deleted successfully",
	})
}
