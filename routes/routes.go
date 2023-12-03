package routes

import (
	"bitcoin-klever-api/controllers"
	"bitcoin-klever-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/health", controllers.HealthCheckHandler).Methods("Get")
	r.HandleFunc("/details/{address}", controllers.GetBitcoinData).Methods("Get")
	r.HandleFunc("/balance/{address}", controllers.GetBalance).Methods("Get")
	r.HandleFunc("/send", controllers.SendHandler).Methods("Post")
	r.HandleFunc("/tx/{tx}", controllers.GetTransactionData).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
