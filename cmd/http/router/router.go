package router

import (
	"github.com/gorilla/mux"
	"glowing-potato/infra/context/handler"
	"net/http"
)

const (
	POS = http.MethodPost
)

func InitRouter(handler *handler.HandlerCtx) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/create-order", handler.OrdersHandler.CreateOrder).Methods(POS)

	return r
}
