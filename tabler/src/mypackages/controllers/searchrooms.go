package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//SearchRooms exported
func SearchRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var rooms []Room
	params := mux.Vars(r)

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

	searchKey := "%" + params["tituloMesa"] + "%"

	result, err := db.Query("SELECT ID_MESA, TITULO_MESA, DESC_MESA FROM mesa WHERE TITULO_MESA LIKE ?", searchKey)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var room Room

	for result.Next() {
		err := result.Scan(&room.ID, &room.Title, &room.Desc)
		if err != nil {
			panic(err.Error())
		}

		rooms = append(rooms, room)
	}

	json.NewEncoder(w).Encode(rooms)
	w.WriteHeader(http.StatusOK)
}
