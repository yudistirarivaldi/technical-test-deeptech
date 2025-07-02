package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/dto"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/model"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/service"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/utils"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{productService: s}
}

func (h *ProductHandler) HandleInsert(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid request body",
		})
		return
	}

	if err := validator.New().Struct(req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Validation failed",
			Errors:       utils.FormatValidationErrors(err),
		})
		return
	}

	p := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		CategoryID:  req.CategoryID,
		Stock:       req.Stock,
	}

	if err := h.productService.Insert(r.Context(), p); err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to insert product",
		})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, model.Response{ResponseCode: "00", Message: "Product created successfully"})
}

func (h *ProductHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, model.Response{
			ResponseCode: "01",
			Message:      "Method not allowed",
		})
		return
	}

	data, err := h.productService.GetAll(r.Context())
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to get products",
		})
		return
	}
	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "Success", Data: data,
	})
}

func (h *ProductHandler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{ResponseCode: "01", Message: "Invalid product ID"})
		return
	}

	product, err := h.productService.GetByID(r.Context(), id)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{ResponseCode: "01", Message: "Failed to get product"})
		return
	}
	if product == nil {
		utils.WriteJSON(w, http.StatusNotFound, model.Response{ResponseCode: "01", Message: "Product not found"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{ResponseCode: "00", Message: "Success", Data: product})
}

func (h *ProductHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid product ID",
		})
		return
	}

	existing, err := h.productService.GetByID(r.Context(), id)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to check product",
		})
		return
	}
	if existing == nil {
		utils.WriteJSON(w, http.StatusNotFound, model.Response{
			ResponseCode: "01",
			Message:      "Product not found",
		})
		return
	}

	var req dto.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid request body",
		})
		return
	}

	if err := validator.New().Struct(req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Validation failed",
			Errors:       utils.FormatValidationErrors(err),
		})
		return
	}

	p := &model.Product{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		CategoryID:  req.CategoryID,
		Stock:       req.Stock,
	}

	if err := h.productService.Update(r.Context(), p); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to update product",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "Product updated successfully",
	})
}

func (h *ProductHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid product ID",
		})
		return
	}

	existing, err := h.productService.GetByID(r.Context(), id)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to check product",
		})
		return
	}
	if existing == nil {
		utils.WriteJSON(w, http.StatusNotFound, model.Response{
			ResponseCode: "01",
			Message:      "Product not found",
		})
		return
	}

	if err := h.productService.Delete(r.Context(), id); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to delete product",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "Product deleted successfully",
	})
}
