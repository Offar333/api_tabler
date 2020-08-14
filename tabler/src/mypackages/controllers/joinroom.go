package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//JoinRoom exported
func JoinRoom(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idMesa := keyVal["ID_MESA"]
	idUsuar := keyVal["ID_USUAR"]
	mestreJoga := keyVal["MESTRE_JOGA"]
	charName := keyVal["NOMECHAR_JOGA"]
	userName := keyVal["UserName"]

	sheetPath := "CharSheet/" + userName + "/" + charName

	//---------------------CREATING FOLDER--------------------
	_, errFolder := os.Stat(sheetPath)

	if os.IsNotExist(errFolder) {
		errDir := os.MkdirAll(sheetPath, os.ModePerm)
		if errDir != nil {
			log.Fatal(errFolder)
		}
	}
	//--------------------------------------------------------

	//----------------------SAVING FILE----------------------
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile(sheetPath, charName+"-*.pdf")

	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	//------------------------------------------------------------------

	stmtIns, err := db.Prepare("INSERT INTO mesa_jogadores(ID_MESA, ID_USUAR, MESTRE_JOGA, FICHA_JOGA) VALUES (?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	//CHECK IF THERE'S ALREADY A DM AT THE TABLE
	var isThereDm int
	_ = db.QueryRow("SELECT COUNT(*) FROM mesa_jogadores WHERE ID_MESA = ? AND MESTRE_JOGA = 1", idMesa).Scan(&isThereDm)

	if err != nil {

		panic(err.Error())

	}

	if isThereDm != 0 && mestreJoga == "1" {

		fmt.Fprintf(w, "Já existe mestre nesta mesa!")

	} else { //IF THERE'S NO DM, INSERT THE PLAYER IN THE ROOM

		_, err = stmtIns.Exec(idMesa, idUsuar, mestreJoga, sheetPath)
		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintf(w, "Jogador inserido na mesa!")
		w.WriteHeader(http.StatusOK)
	}

}
