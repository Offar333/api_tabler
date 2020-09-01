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

//CreateRoom exported
func CreateRoom(w http.ResponseWriter, r *http.Request) {
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

	stmtIns, err := db.Prepare("INSERT INTO mesa(ADM_MESA, TITULO_MESA, DESC_MESA, QTDEJOG_MESA, FORMA_MESA, STATUS_MESA, LVLINIC_MESA, EXPJOGO_MESA, LINK_CHAT) VALUES (?,?,?,?,?,?,?,?,?) ")

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	admMesa := keyVal["ADM_MESA"]
	tituloMesa := keyVal["TITULO_MESA"]
	descMesa := keyVal["DESC_MESA"]
	qtdejogMesa := keyVal["QTDEJOG_MESA"]
	formaMesa := keyVal["FORMA_MESA"]
	statusMesa := keyVal["STATUS_MESA"]
	lvlinicMesa := keyVal["LVLINIC_MESA"]
	expJogoMesa := keyVal["EXPJOGO_MESA"]
	linkChat := keyVal["LINK_CHAT"]

	_, err = stmtIns.Exec(admMesa, tituloMesa, descMesa, qtdejogMesa, formaMesa, statusMesa, lvlinicMesa, expJogoMesa, linkChat)
	if err != nil {
		panic(err.Error())
	}

	result, err := db.Query("SELECT ID_MESA, ADM_MESA, TITULO_MESA, DESC_MESA, QTDEJOG_MESA, FORMA_MESA, STATUS_MESA, LVLINIC_MESA, EXPJOGO_MESA, LINK_CHAT  FROM mesa WHERE ID_MESA = LAST_INSERT_ID()")

	if err != nil {

		panic(err.Error())
	}

	defer result.Close()

	var room Room

	for result.Next() {

		err := result.Scan(&room.ID, &room.AdmMesa, &room.Title, &room.Desc, &room.QtdeJog, &room.Formato, &room.Status, &room.LvlInic, &room.ExpJogo, &room.LinkChat)

		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(room)
	w.WriteHeader(http.StatusOK)
}
