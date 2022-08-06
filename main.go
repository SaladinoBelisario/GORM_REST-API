package main

import (
	"github.com/SaladinoBelisario/GORM_REST-API/db"
	"github.com/SaladinoBelisario/GORM_REST-API/models"
	"github.com/SaladinoBelisario/GORM_REST-API/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// database connection
	db.DBConnection()
	// db.DB.Migrator().DropTable(models.User{})
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	subrouter := router.PathPrefix("/api").Subrouter()

	// tasks routes
	subrouter.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	subrouter.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	subrouter.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")

	// users routes
	subrouter.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	subrouter.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	subrouter.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	subrouter.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	_ = http.ListenAndServe(":4000", router)
}
