package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/nrednav/cuid2"
	"github.com/yash91989201/go_microservice/model"
	"github.com/yash91989201/go_microservice/repository/order"
)

type Order struct {
	Repo *order.RedisRepo
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CustomerID string           `json:"customer_id"`
		LineItems  []model.LineItem `json:line_items`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()

	order := model.Order{
		CreatedAt:  &now,
		OrderID:    cuid2.Generate(),
		CustomerID: body.CustomerID,
		LineItems:  body.LineItems,
	}

	err := o.Repo.Insert(r.Context(), order)
	if err != nil {
		fmt.Println("failed to insert")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(order)
	if err != nil {
		fmt.Println("failed to marshal :%w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
	w.WriteHeader(http.StatusCreated)
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {

}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {

}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {

}

func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {

}
