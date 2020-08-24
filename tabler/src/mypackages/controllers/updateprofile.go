package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//UpdateProfile exported
func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

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
