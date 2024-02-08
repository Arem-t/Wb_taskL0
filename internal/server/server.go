package server

import (
	"fmt"
	"net/http"
)

var jsonCache []byte

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderId := r.URL.Query().Get("id")
	jsonCache = GetCache(orderId)
	fmt.Fprintf(w, "Order details for ID %s", orderId)
	fmt.Fprintf(w, "\n", string(jsonCache))
}

func StartHandler() {
	http.HandleFunc("/order", getOrderHandler)
	http.Handle("/", http.FileServer(http.Dir("static")))
	fmt.Println("Server is starting...")
	if err := http.ListenAndServe(":9191", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
