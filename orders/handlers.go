package orders

import (
	"net/http"
	"strconv"
)


func OrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "web/index.html")
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err!=nil{
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		product := r.FormValue("product")
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err!=nil{
			http.Error(w, "Invalid qty", http.StatusBadRequest)
			return
		}
		order := Order{
			Product: product,
			Quantity: quantity,
		}
		go Producer(order)
		w.Write([]byte("Order placed!!!"))
	}
}