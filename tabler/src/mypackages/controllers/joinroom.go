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

//JoinRoom exported
func JoinRoom(w http.ResponseWriter, r *http.Request) {
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
	idMesa := keyVal["ID_MESA"]
	idUsuar := keyVal["ID_USUAR"]
	mestreJoga := keyVal["MESTRE_JOGA"]
	charName := keyVal["NOMECHAR_JOGA"]
	charClass := keyVal["CLASSECHAR_JOGA"]
	userName := keyVal["UserName"]

	sheetPath := "CharSheet/" + userName + "/" + charName

	/* //---------------------CREATING FOLDER--------------------
	_, errFolder := os.Stat(sheetPath)

	if os.IsNotExist(errFolder) {
		errDir := os.MkdirAll(sheetPath, os.ModePerm)
		if errDir != nil {
			log.Fatal(errFolder)
		}
	}
	//-------------------------------------------------------- */

	/* //----------------------SAVING FILE----------------------
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
	//------------------------------------------------------------------ */

	//CHECK IF THERE'S ALREADY A DM AT THE TABLE
	var isThereDm int
	_ = db.QueryRow("SELECT COUNT(*) FROM mesa_jogadores WHERE ID_MESA = ? AND MESTRE_JOGA = 1", idMesa).Scan(&isThereDm)

	if err != nil {

		panic(err.Error())

	}

	if isThereDm != 0 && mestreJoga == "1" {

		res := DoesExist{JaExiste: "mestre"}
		json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusFound)

	} else { //IF THERE'S NO DM, INSERT THE PLAYER IN THE ROOM

		stmtIns, err := db.Prepare("INSERT INTO mesa_jogadores(ID_MESA, ID_USUAR, MESTRE_JOGA, FICHA_JOGA, NOMECHAR_JOGA, CLASSECHAR_JOGA) VALUES (?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}

		/* _, err = stmtIns.Exec(idMesa, idUsuar, mestreJoga, "C:/SheetPath") */ //Tests
		if mestreJoga == "1" {
			_, err = stmtIns.Exec(idMesa, idUsuar, mestreJoga, "Mestre", "Mestre", "Mestre")
		} else {
			_, err = stmtIns.Exec(idMesa, idUsuar, mestreJoga, sheetPath, charName, charClass)
		}

		if err != nil {
			panic(err.Error())
		}

		res := DoesExist{JaExiste: "cadastrado"}
		json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusOK)
	}

}
