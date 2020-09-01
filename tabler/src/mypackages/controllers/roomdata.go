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

//RoomData exported
func RoomData(w http.ResponseWriter, r *http.Request) {
	var roomData RoomPlayers

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
	idUser := keyVal["ID_USUAR"]

	//CHECK IF THE PLAYER IS AT THE TABLE
	roomData.AlreadyIn = "no"
	var isAtTable int
	_ = db.QueryRow("SELECT COUNT(*) FROM mesa_jogadores WHERE ID_MESA = ? AND ID_USUAR = ?", idMesa, idUser).Scan(&isAtTable)

	if isAtTable != 0 {
		roomData.AlreadyIn = "yes"
	}

	//CHECK IF THERE`S ALREADY A DM AT THE TABLE
	var isThereDm int

	_ = db.QueryRow("SELECT COUNT(*) FROM mesa_jogadores WHERE ID_MESA = ? AND MESTRE_JOGA = 1", idMesa).Scan(&isThereDm)

	if err != nil {

		panic(err.Error())

	}

	if isThereDm == 0 {
		//POPULATE THE STRUCT WHITH THE INFORMATION THAT THERE IS NO DM AVAILABLE
		roomData.DungeonMaster = "NO_DM"
		roomData.TablesJoined = 0

	} else if isThereDm != 0 {

		var idDm int

		//POPULATE THE STRUCT WITH THE DM INFO
		//RETURNS THE NAME OF THE DM AT THE TABLE
		dmInfo, err := db.Query("SELECT a.ID_USUAR, a.APELIDO_USUAR AS MESTRE FROM usuario a "+
			"INNER JOIN mesa_jogadores b ON a.ID_USUAR = b.ID_USUAR "+
			"WHERE b.MESTRE_JOGA = 1 AND b.ID_MESA = ?", idMesa)

		if err != nil {
			panic(err.Error())
		}
		for dmInfo.Next() {

			err := dmInfo.Scan(&idDm, &roomData.DungeonMaster)
			if err != nil {
				panic(err.Error())
			}

		}

		_ = db.QueryRow("SELECT COUNT(*) FROM mesa_jogadores WHERE ID_USUAR = ? AND MESTRE_JOGA = 1", idDm).Scan(&roomData.TablesJoined)

		if err != nil {

			panic(err.Error())

		}

		defer dmInfo.Close()

	}

	//CHECK IF THERE ARE PLAYERS AT THE TABLE
	var isTherePlayers int
	_ = db.QueryRow("SELECT COUNT(*) FROM mesa_jogadores WHERE ID_MESA = ? AND MESTRE_JOGA = 0", idMesa).Scan(&isTherePlayers)

	if err != nil {

		panic(err.Error())

	}

	if isTherePlayers == 0 { //OK

		//POPULATE THE STRUCT WHITH THE INFORMATION THAT THERE IS NO PLAYERS AT THE TABLE
		roomData.Players = []PlayersInfo{
			{
				PlayerName:  "NO_PLAYERS",
				PlayerChar:  "NO_PLAYERS",
				PlayerClass: "NO_PLAYERS",
			},
		}

	} else if isTherePlayers != 0 { //OK

		result, err := db.Query("SELECT a.APELIDO_USUAR, b.NOMECHAR_JOGA, b.CLASSECHAR_JOGA FROM usuario a "+
			"INNER JOIN mesa_jogadores b ON a.ID_USUAR = b.ID_USUAR "+
			"WHERE b.ID_MESA = ? AND b.NOMECHAR_JOGA <> '' ", idMesa)

		if err != nil {
			panic(err.Error())
		}

		var info PlayersInfo

		for result.Next() {
			err := result.Scan(&info.PlayerName, &info.PlayerChar, &info.PlayerClass)

			if err != nil {
				panic(err.Error())
			}

			roomData.Players = append(roomData.Players, info)

		}

		defer result.Close()

	}

	json.NewEncoder(w).Encode(roomData)

	w.WriteHeader(http.StatusOK)

}
