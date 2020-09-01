//IT ALWAYS STARTS LIKE THIS
package main

//THE STANDARD GOLANG SHENANIGANS TO IMPORT PACKAGES

import (
	"database/sql"
	"log"
	"mypackages/controllers"
	"net/http" //TO HANDLE HTTP REQUESTS

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

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
	r.HandleFunc("/api/roomData", controllers.RoomData).Methods("POST")
	r.HandleFunc("/api/roomFormat", controllers.RoomFormat).Methods("GET")
	//-----------------------------------------------------------------------

	//----------------------USER ROUTES--------------------------------------
	r.HandleFunc("/api/createUser", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/api/updateProfile", controllers.UpdateProfile).Methods("PUT")
	r.HandleFunc("/api/uploadAvatar", controllers.UploadAvatar).Methods("POST")
	r.HandleFunc("/api/login", controllers.CheckLogin).Methods("POST")
	r.HandleFunc("/api/userData", controllers.UserData).Methods("POST")
	r.HandleFunc("/api/roomsJoined", controllers.RoomsJoined).Methods("POST")
	//-----------------------------------------------------------------------

	//----------------------TEST ROUTES--------------------------------------
	r.HandleFunc("/api/jwtTest", controllers.TestJwt).Methods("POST")
	//-----------------------------------------------------------------------

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	handler := c.Handler(r)
	log.Println("Server Online!")
	log.Fatal(http.ListenAndServe(":8000", handler))

}
