package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//RoomFormat exported
func RoomFormat(w http.ResponseWriter, r *http.Request) {
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

	var formats []GameFormats

	result, err := db.Query("SELECT FORM_DESC FROM formatos")

	if err != nil {
		panic(err.Error())
	}

	var format GameFormats

	for result.Next() {
		err := result.Scan(&format.DescFormat)
		if err != nil {
			panic(err.Error())
		}

		formats = append(formats, format)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(formats)
	w.WriteHeader(http.StatusOK)
}
