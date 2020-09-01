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

//CreateUser exported
func CreateUser(w http.ResponseWriter, r *http.Request) {
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
	avatarUsuar := "C:/FotoPadrao"

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

	stmtIns, err := db.Prepare("INSERT INTO usuario(NOME_USUAR, APELIDO_USUAR, SENHA_USUAR, EMAIL_USUAR, AVATAR_USUAR) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmtIns.Exec(nomeUsuar, apelidoUsuar, senhaUsuar, emailUsuar, avatarUsuar)
	if err != nil {
		panic(err.Error())
	}

	res := DoesExist{JaExiste: "usuarioCriado"}
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)

}
