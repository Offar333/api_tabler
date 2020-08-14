package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

//Room Struct
type Room struct {
	//ID    int    `json:"id"`
	//Title string `json:"title"`
	//Desc  string `json:"desc"`
}

//HomePage exported
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var rooms []Room

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

	log.Println("Bateu aqui!")

	if err != nil {
		panic(err.Error())
	}

	result, err := db.Query("SELECT ID_MESA, TITULO_MESA, DESC_MESA FROM mesa")

	if err != nil {

		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var room Room
		//err := result.Scan(&room.ID, &room.Title, &room.Desc)
		//err := result.Scan(&room.Desc)

		if err != nil {
			panic(err.Error())
		}
		rooms = append(rooms, room)
	}

	json.NewEncoder(w).Encode(rooms)

}