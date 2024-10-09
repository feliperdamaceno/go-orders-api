package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/feliperdamaceno/go-orders-api/internal/model"
	"github.com/feliperdamaceno/go-orders-api/internal/repository/order"
	"github.com/google/uuid"
)

type OrderHandler struct {
	Repo *order.RedisRepo
}

type OrderCreateResponse struct {
	CustomerId uuid.UUID         `json:"customerId"`
	OrderItems []model.OrderItem `json:"orderItems"`
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	body := OrderCreateResponse{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()

	order := model.Order{
		Id:         rand.Uint64(),
		CustomerId: body.CustomerId,
		OrderItems: body.OrderItems,
		CreatedAt:  &now,
	}

	err = h.Repo.Create(r.Context(), order)
	if err != nil {
		fmt.Println("failed to create order: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	res, err := json.Marshal(order)
	if err != nil {
		fmt.Println("failed to encode order: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

type OrderGetAllResponse struct {
	Orders []model.Order `json:"orders"`
	Next   uint64        `json:"next,omitempty"`
}

func (h *OrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	cursorStr := r.URL.Query().Get("cursor")
	if cursorStr == "" {
		cursorStr = "0"
	}

	const decimal = 10
	const bit = 64
	cursor, err := strconv.ParseUint(cursorStr, decimal, bit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	const size = 50
	res, err := h.Repo.GetAll(r.Context(), order.GetAllPage{
		Cursor: cursor,
		Size:   size,
	})

	if err != nil {
		fmt.Println("failed to get all orders: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := OrderGetAllResponse{
		Orders: res.Orders,
		Next:   res.Cursor,
	}

	data, err := json.Marshal(response)
	if err != nil {
		fmt.Println("failed to encode orders: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(data)
	w.WriteHeader(http.StatusFound)
}

func (h *OrderHandler) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Retrieve One")
}

func (h *OrderHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One")
}

func (h *OrderHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete")
}
