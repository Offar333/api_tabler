package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//RoomData exported
func RoomData(w http.ResponseWriter, r *http.Request) {

	var room Room

	params := mux.Vars(r)

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

	//searchKey := params["idMesa"]

	result, err := db.Query("SELECT ID_MESA, ADM_MESA, TITULO_MESA, DESC_MESA, QTDEJOG_MESA, FORMA_MESA, STATUS_MESA FROM mesa WHERE ID_MESA LIKE ?", params["idMesa"])

	if err != nil {
		panic(err.Error())
	}

	/*
		ID      int    `json:"id"`
		AdmMesa string `json:"admMesa"`
		Title   string `json:"title"`
		Desc    string `json:"desc"`
		QtdeJog int    `json:"qtdeJog"`
		Formato string `json:"formato"`
		Status  int    `json:"status"`
	*/

	for result.Next() {
		err := result.Scan(&room.ID, &room.AdmMesa, &room.Title, &room.Desc, &room.QtdeJog, &room.Formato, &room.Status)
		if err != nil {
			panic(err.Error())
		}
	}

	defer result.Close()

	json.NewEncoder(w).Encode(room)
	w.WriteHeader(http.StatusOK)

}
