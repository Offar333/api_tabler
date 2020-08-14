package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//CreateUser exported
func CreateUser(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

	stmtIns, err := db.Prepare("INSERT INTO usuario(NOME_USUAR, APELIDO_USUAR, SENHA_USUAR, EMAIL_USUAR) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	nomeUsuar := keyVal["NOME_USUAR"]
	apelidoUsuar := keyVal["APELIDO_USUAR"]
	senhaUsuar := keyVal["SENHA_USUAR"]
	emailUsuar := keyVal["EMAIL_USUAR"]

	_, err = stmtIns.Exec(nomeUsuar, apelidoUsuar, senhaUsuar, emailUsuar)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Usuario Criado!")
	w.WriteHeader(http.StatusOK)

}
