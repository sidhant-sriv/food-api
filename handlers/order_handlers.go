package api

import (
	"encoding/json"
	"fmt"
	"foodapi/models"
	"foodapi/reddb"
	"net/http"

	"github.com/gorilla/mux"
)

func recordOrder(w http.ResponseWriter, r *http.Request) {
    var order models.Order
    json.NewDecoder(r.Body).Decode(&order)

    // Save order details to Redis
    err := reddb.SetOrder(order.OrderID, "Order recorded")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Order %s recorded", order.OrderID)
}

func prepareOrder(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    orderID := vars["orderID"]

    err := reddb.SetOrder(orderID, "Order being prepared")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Order %s is being prepared", orderID)
}

func dispatchOrder(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    orderID := vars["orderID"]

    err := reddb.SetOrder(orderID, "Order dispatched")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Order %s dispatched", orderID)
}

func getOrderDetails(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    orderID := vars["orderID"]

    orderStatus, err := reddb.GetOrder(orderID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    order := models.Order{
        OrderID:     orderID,
        OrderStatus: orderStatus,
    }

    response, err := json.Marshal(order)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func SetupRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/record", recordOrder).Methods("POST")
    router.HandleFunc("/prepare/{orderID}", prepareOrder).Methods("POST")
    router.HandleFunc("/dispatch/{orderID}", dispatchOrder).Methods("POST")
    router.HandleFunc("/order/{orderID}", getOrderDetails).Methods("GET")

    return router
}
