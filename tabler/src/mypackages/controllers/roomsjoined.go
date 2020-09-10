package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//RoomsJoined exported
func RoomsJoined(w http.ResponseWriter, r *http.Request) {
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	userName := keyVal["ID_USUAR"]

	result, err := db.Query("SELECT a.ADM_MESA, a.ID_MESA, a.TITULO_MESA, a.DESC_MESA , a.LVLINIC_MESA, a.EXPJOGO_MESA, a.FORMA_MESA, a.LINK_CHAT, a.QTDEJOG_MESA, a.STATUS_MESA FROM mesa a "+
		"INNER JOIN mesa_jogadores b ON a.ID_MESA = b.ID_MESA WHERE b.ID_USUAR = ?", userName)

	if err != nil {
		panic(err.Error())
	}

	var room Room
	/*
		ID       int    `json:"id"`
		AdmMesa  string `json:"admMesa"`
		Title    string `json:"title"`
		Desc     string `json:"desc"`
		QtdeJog  string `json:"qtdeJog"`
		Formato  string `json:"formato"`
		Status   int    `json:"status"`
		LvlInic  string `json:"lvlInic"`
		ExpJogo  string `json:"expJogo"`
		LinkChat string `json:"linkChat"`
	*/

	for result.Next() {
		err := result.Scan(&room.AdmMesa, &room.ID, &room.Title, &room.Desc, &room.LvlInic, &room.ExpJogo, &room.Formato, &room.LinkChat, &room.QtdeJog, &room.Status)
		if err != nil {
			panic(err.Error())
		}

		rooms = append(rooms, room)
	}

	defer result.Close()

	if len(rooms) < 1 {
		var noRooms CustomResponse
		noRooms.CustomResp = "NoRooms"
		json.NewEncoder(w).Encode(noRooms)
		w.WriteHeader(http.StatusOK)
	} else {
		json.NewEncoder(w).Encode(rooms)
		w.WriteHeader(http.StatusOK)
	}

}
