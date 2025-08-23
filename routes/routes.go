package routes

import (
	"net/http"

	"petstore-api/handlers"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SetupRoutes(sellerHandler *handlers.SellerHandler, petHandler *handlers.PetHandler) http.Handler {
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api := r.PathPrefix("/").Subrouter()

	api.HandleFunc("/sellers", sellerHandler.GetSellers).Methods("GET")
	api.HandleFunc("/sellers/{id}", sellerHandler.GetSeller).Methods("GET")
	api.HandleFunc("/sellers", sellerHandler.CreateSeller).Methods("POST")
	api.HandleFunc("/sellers/{id}", sellerHandler.UpdateSeller).Methods("PUT")
	api.HandleFunc("/sellers/{id}", sellerHandler.DeleteSeller).Methods("DELETE")

	api.HandleFunc("/pets", petHandler.GetPets).Methods("GET")
	api.HandleFunc("/pets/{id}", petHandler.GetPet).Methods("GET")
	api.HandleFunc("/pets", petHandler.CreatePet).Methods("POST")
	api.HandleFunc("/pets/{id}", petHandler.UpdatePet).Methods("PUT")
	api.HandleFunc("/pets/{id}", petHandler.DeletePet).Methods("DELETE")

	return enableCORS(r)
}
