package routers

import (
	"fmt"
	"github.com/Eydzhpee08/university/handlers/db"
	"github.com/Eydzhpee08/university/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// HandlerRouting Router function
func HandlerRouting() {
	r := mux.NewRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: utils.AddCors(r),
	}

	r.HandleFunc("/employees", db.GetEmployees).Methods("GET")
	r.HandleFunc("/employee", db.CreateEmployee).Methods("POST")
	r.HandleFunc("/employee/docx", db.GetAllDocx).Methods("GET")
	r.HandleFunc("/admin/signIn", db.SignIn).Methods("POST")
	r.HandleFunc("/admin/employees", db.GetAllEmployees).Methods("GET")
	fmt.Println("Server is start to listening on port: ", srv.Addr)
	log.Fatal("INFO ", srv.ListenAndServe())
}
