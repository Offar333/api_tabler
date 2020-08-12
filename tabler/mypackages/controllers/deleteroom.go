package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//DeleteRoom exported
func DeleteRoom(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

	stmt, err := db.Prepare("DELETE FROM mesa WHERE ID_MESA = ?")

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

	_, err = stmt.Exec(idMesa)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Mesa com o ID = %s foi deletada", idMesa)
	w.WriteHeader(http.StatusOK)
}
