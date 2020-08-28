package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

//TestJwt exported
func TestJwt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//NEED TO MAKE A FUNCTION FOR THIS, BUT I DUNNO HOW YET ):
	//LOADING .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//CONNECTING TO DB
	db, err = sql.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_CONN"))
	//-----------------------------------------------------------------)

	var user User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	userName := keyVal["APELIDO_USUAR"]
	userPassword := keyVal["SENHA_USUAR"]

	var doesExist int
	_ = db.QueryRow("SELECT COUNT(*) FROM usuario WHERE APELIDO_USUAR = ? AND SENHA_USUAR = ?", userName, userPassword).Scan(&doesExist)

	if err != nil {

		panic(err.Error())

	}

	if doesExist != 0 {

		userData, err := db.Query("SELECT ID_USUAR, NOME_USUAR, APELIDO_USUAR, EMAIL_USUAR, AVATAR_USUAR FROM usuario WHERE APELIDO_USUAR = ?", userName)

		if err != nil {
			panic(err.Error())
		}

		for userData.Next() {

			err := userData.Scan(&user.ID, &user.Nome, &user.Apelido, &user.Email, &user.AvatarPath)
			if err != nil {
				panic(err.Error())
			}
		}

		// Create a new token object, specifying signing method and the claims
		// you would like it to contain.
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":         user.ID,
			"nome":       user.Nome,
			"apelido":    user.Apelido,
			"email":      user.Email,
			"avatarpath": user.AvatarPath,
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte(os.Getenv("API_SECRET")))

		fmt.Println(tokenString, err)

		json.NewEncoder(w).Encode(tokenString)
		w.WriteHeader(http.StatusOK)

	} else {
		res := DoesExist{JaExiste: "UsuarioInexistente"}

		json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusNotFound)
	}

}
