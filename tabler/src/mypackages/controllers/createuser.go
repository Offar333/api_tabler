package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//CreateUser exported
func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

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

	var doesExistApelido int
	var doesExistEmail int

	_ = db.QueryRow("SELECT COUNT(*) FROM usuario WHERE APELIDO_USUAR = ?", apelidoUsuar).Scan(&doesExistApelido)

	if doesExistApelido != 0 {
		res := DoesExist{JaExiste: "apelido"}

		json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusFound)

		return
	}

	_ = db.QueryRow("SELECT COUNT(*) FROM usuario WHERE EMAIL_USUAR = ?", emailUsuar).Scan(&doesExistEmail)

	if doesExistEmail != 0 {
		res := DoesExist{JaExiste: "email"}

		json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusFound)

		return

	}

	stmtIns, err := db.Prepare("INSERT INTO usuario(NOME_USUAR, APELIDO_USUAR, SENHA_USUAR, EMAIL_USUAR) VALUES (?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmtIns.Exec(nomeUsuar, apelidoUsuar, senhaUsuar, emailUsuar)
	if err != nil {
		panic(err.Error())
	}

	res := DoesExist{JaExiste: "usuarioCriado"}
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)

}
