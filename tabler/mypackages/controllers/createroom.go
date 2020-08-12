package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//CreateRoom exported
func CreateRoom(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

	stmtIns, err := db.Prepare("INSERT INTO mesa(ADM_MESA, TITULO_MESA, DESC_MESA, QTDEJOG_MESA, FORMA_MESA, STATUS_MESA) VALUES (?,?,?,?,?,?) ")

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

	_, err = stmtIns.Exec(admMesa, tituloMesa, descMesa, qtdejogMesa, formaMesa, statusMesa)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Mesa criada com sucesso!")
	w.WriteHeader(http.StatusOK)
}
