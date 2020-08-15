package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

//HomePage exported
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	EnableCors(&w)
	var rooms []Room

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

	log.Println("Bateu aqui!")

	if err != nil {
		panic(err.Error())
	}

	result, err := db.Query("SELECT ID_MESA, ADM_MESA, TITULO_MESA, DESC_MESA, QTDEJOG_MESA, FORMA_MESA, STATUS_MESA  FROM mesa")

	if err != nil {

		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var room Room

		/*
			ID      int    `json:"id"`
			AdmMesa string `json:"admMesa"`
			Title   string `json:"title"`
			Desc    string `json:"desc"`
			QtdeJog int    `json:"qtdeJog"`
			Formato string `json:"formato"`
			Status  int    `json:"status"`
		*/

		err := result.Scan(&room.ID, &room.AdmMesa, &room.Title, &room.Desc, &room.QtdeJog, &room.Formato, &room.Status)

		if err != nil {
			panic(err.Error())
		}
		rooms = append(rooms, room)
	}

	json.NewEncoder(w).Encode(rooms)
	w.WriteHeader(http.StatusOK)

}
