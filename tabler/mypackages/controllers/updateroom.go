package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//UpdateRoom FUNCTION
func UpdateRoom(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

	stmtIns, err := db.Prepare("UPDATE mesa SET TITULO_MESA = ? , DESC_MESA = ? , QTDEJOG_MESA = ? , FORMA_MESA = ? WHERE ID_MESA = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idMesa := keyVal["ID_MESA"]
	tituloMesa := keyVal["TITULO_MESA"]
	descMesa := keyVal["DESC_MESA"]
	qtdejogMesa := keyVal["QTDEJOG_MESA"]
	formaMesa := keyVal["FORMA_MESA"]

	_, err = stmtIns.Exec(tituloMesa, descMesa, qtdejogMesa, formaMesa, idMesa)
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
}
