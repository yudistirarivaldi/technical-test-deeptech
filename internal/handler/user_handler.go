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

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, model.Response{
			ResponseCode: "01",
			Message:      "Method not allowed",
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

	user, err := h.userService.GetByID(r.Context(), int64(userID))
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to get user",
		})
		return
	}
	if user == nil {
		utils.WriteJSON(w, http.StatusNotFound, model.Response{
			ResponseCode: "01",
			Message:      "USer not found",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "User profile retrieved successfully",
		Data:         user,
	})
}

func (h *UserHandler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, model.Response{
			ResponseCode: "01",
			Message:      "Method not allowed",
		})
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		utils.WriteJSON(w, http.StatusUnauthorized, model.Response{
			ResponseCode: "01",
			Message:      "Unauthorized",
		})
		return
	}

	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid JSON format",
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

	birthDate, err := utils.ParseDate(req.DateOfBirth)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid birth_date format",
		})
		return
	}

	user := &model.Users{
		ID:          userID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		DateOfBirth: birthDate,
		Gender:      req.Gender,
	}

	if err := h.userService.Update(r.Context(), user); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to update user",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "User updated successfully",
	})
}
