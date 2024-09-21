package handler

import (
	"fmt"
	"net/http"
)

type OrderHandler struct{}

func (o *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One")
}

func (o *OrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Retrieve All")
}

func (o *OrderHandler) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Retrieve One")
}

func (o *OrderHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One")
}

func (o *OrderHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete")
}
