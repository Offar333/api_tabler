package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

//RoomFormat exported
func RoomFormat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var formats []GameFormats

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

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
