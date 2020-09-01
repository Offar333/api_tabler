package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//SearchRooms exported
func SearchRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//NEED TO MAKE A FUNCTION FOR THIS, BUT I DUNNO HOW YET ):
	//LOADING .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//CONNECTING TO DB
	db, err = sql.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_CONN"))
	//-----------------------------------------------------------------

	var rooms []Room
	params := mux.Vars(r)

	searchKey := "%" + params["tituloMesa"] + "%"

	result, err := db.Query("SELECT ID_MESA, TITULO_MESA, DESC_MESA , LVLINIC_MESA, EXPJOGO_MESA FROM mesa WHERE TITULO_MESA LIKE ?", searchKey)

	if err != nil {
		panic(err.Error())
	}

	var room Room

	for result.Next() {
		err := result.Scan(&room.ID, &room.Title, &room.Desc, &room.LvlInic, &room.ExpJogo)
		if err != nil {
			panic(err.Error())
		}

		rooms = append(rooms, room)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(rooms)
	w.WriteHeader(http.StatusOK)
}
