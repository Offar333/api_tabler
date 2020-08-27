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

//UpdateProfile exported
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
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

	var user User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	nomeUsuar := keyVal["NOME_USUAR"]
	apelidoUsuar := keyVal["APELIDO_USUAR"]
	emailUsuar := keyVal["EMAIL_USUAR"]
	avatarUsuar := keyVal["AVATAR_USUAR"]
	idUsuar := keyVal["ID_USUAR"]

	//NEED TO CHECK IF THE NEW USERNAME ALREADY EXISTS, IMPLEMENT LATER

	stmtIns, err := db.Prepare("UPDATE usuario SET NOME_USUAR = ? , APELIDO_USUAR = ? , EMAIL_USUAR = ? , AVATAR_USUAR = ? WHERE ID_USUAR = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmtIns.Exec(nomeUsuar, apelidoUsuar, emailUsuar, avatarUsuar, idUsuar)
	if err != nil {
		panic(err.Error())
	}

	userData, err := db.Query("SELECT ID_USUAR, NOME_USUAR, APELIDO_USUAR, EMAIL_USUAR, AVATAR_USUAR FROM usuario WHERE ID_USUAR = ?", idUsuar)

	if err != nil {
		panic(err.Error())
	}

	for userData.Next() {

		err := userData.Scan(&user.ID, &user.Nome, &user.Apelido, &user.Email, &user.AvatarPath)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)
}
