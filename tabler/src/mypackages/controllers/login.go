package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//User Struct
type User struct {
	ID         string `json:"id"`
	Nome       string `json:"nome"`
	Apelido    string `json:"apelido"`
	Email      string `json:"email"`
	AvatarPath string `json:"avatarpath"`
}

//CheckLogin exported
func CheckLogin(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

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
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Println("doesExist? = " + strconv.Itoa(doesExist))
	}

}
