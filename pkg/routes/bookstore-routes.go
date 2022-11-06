package routes

import (
	"github.com/gorilla/mux"
	"github.com/nguyenkhanhvn/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(routes *mux.Router) {
	routes.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	routes.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	routes.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	routes.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	routes.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
