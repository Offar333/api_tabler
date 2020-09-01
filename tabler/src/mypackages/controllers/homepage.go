package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//HomePage exported
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var rooms []Room

	//NEED TO MAKE A FUNCTION FOR THIS, BUT I DUNNO HOW YET ):
	//LOADING .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//CONNECTING TO DB
	db, err = sql.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_CONN"))
	//-----------------------------------------------------------------

	if err != nil {
		panic(err.Error())
	}

	result, err := db.Query("SELECT ID_MESA, ADM_MESA, TITULO_MESA, DESC_MESA, QTDEJOG_MESA, FORMA_MESA, STATUS_MESA, LVLINIC_MESA, EXPJOGO_MESA  FROM mesa")

	if err != nil {

		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var room Room

		err := result.Scan(&room.ID, &room.AdmMesa, &room.Title, &room.Desc, &room.QtdeJog, &room.Formato, &room.Status, &room.LvlInic, &room.ExpJogo)

		if err != nil {
			panic(err.Error())
		}
		rooms = append(rooms, room)
	}

	json.NewEncoder(w).Encode(rooms)
	w.WriteHeader(http.StatusOK)

}
