//IT ALWAYS STARTS LIKE THIS
package main

//THE STANDARD GOLANG SHENANIGANS TO IMPORT PACKAGES

import (
	"database/sql"
	"log"
	"net/http" //TO HANDLE HTTP REQUESTS
	"tabler/backend/mypackages/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//Room Struct
type Room struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

var db *sql.DB
var err error

func main() {

	//MUX IS USED TO CREATE THE BACKEND ROUTES
	r := mux.NewRouter()

	//----------------------ROOM ROUTES--------------------------------------
	r.HandleFunc("/api/homePage", controllers.HomePage).Methods("GET")
	r.HandleFunc("/api/searchRooms/{tituloMesa}", controllers.SearchRooms).Methods("GET")
	r.HandleFunc("/api/createRoom", controllers.CreateRoom).Methods("POST")
	r.HandleFunc("/api/joinRoom", controllers.JoinRoom).Methods("POST")
	r.HandleFunc("/api/updateRoom", controllers.UpdateRoom).Methods("PUT")
	r.HandleFunc("/api/deleteRoom", controllers.DeleteRoom).Methods("DELETE")
	//-----------------------------------------------------------------------

	//----------------------USER ROUTES--------------------------------------
	r.HandleFunc("/api/createUser", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/api/updateProfile", controllers.UpdateProfile).Methods("PUT")
	r.HandleFunc("/api/uploadAvatar", controllers.UploadAvatar).Methods("POST")
	//-----------------------------------------------------------------------

	log.Println("Server Online!")
	log.Fatal(http.ListenAndServe(":8000", r))

}
