package controllers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//UploadAvatar exported
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
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

	var picPath string
	var userName string

	userName = "Offar"

	picPath = "ProfilePics/" + userName

	//---------------------CREATING FOLDER--------------------
	//Sets folder path and folder name
	_, err = os.Stat(picPath)

	//Checks if that folder path already exists
	//If the nested folder exists, do nothing
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(picPath, os.ModePerm)
		stmtUpdt, err := db.Prepare("UPDATE usuario SET AVATAR_USUAR = ? WHERE APELIDO_USUAR = ?")
		_, err = stmtUpdt.Exec(picPath, userName)
		if err != nil {
			panic(err.Error())
		}
		if errDir != nil {
			log.Fatal(err)
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

	tempFile, err := ioutil.TempFile(picPath, userName+"-*.png")

	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)
	w.WriteHeader(http.StatusOK)
	//------------------------------------------------------------------
}

//--------------------------------------------------------------------------
