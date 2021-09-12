package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/korathbasil/chat-app-api/controllers"
)

func InitializeUserRoute(subRouter *mux.Router) {
	subRouter.HandleFunc("/", userHandler).Methods("GET")

	userController := controllers.UserController{}

	subRouter.HandleFunc("/login", userController.LoginUserHandler)
}

func userHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "User Route")
}
