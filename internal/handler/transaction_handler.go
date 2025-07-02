package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/dto"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/middleware"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/model"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/service"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/utils"
)

type TransactionHandler struct {
	transactionService *service.TransactionService
}

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: s}
}

func (h *TransactionHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTransactionRequest
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

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok || userID == 0 {
		utils.WriteJSON(w, http.StatusUnauthorized, model.Response{
			ResponseCode: "01",
			Message:      "Unauthorized",
		})
		return
	}

	req.UserID = userID
	if err := h.transactionService.Create(r.Context(), &req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      err.Error(),
		})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, model.Response{
		ResponseCode: "00",
		Message:      "Transaction created successfully",
	})
}
