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

//UserData exported
func UserData(w http.ResponseWriter, r *http.Request) {
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
	userID := keyVal["ID_USUAR"]

	var doesExist int
	_ = db.QueryRow("SELECT COUNT(*) FROM usuario WHERE ID_USUAR = ?", userID).Scan(&doesExist)

	if err != nil {

		panic(err.Error())

	}

	if doesExist != 0 {

		userData, err := db.Query("SELECT ID_USUAR, NOME_USUAR, APELIDO_USUAR, EMAIL_USUAR, AVATAR_USUAR FROM usuario WHERE ID_USUAR = ?", userID)

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

	} else {
		res := DoesExist{JaExiste: "UsuarioInexistente"}

		json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusNotFound)
	}

}
