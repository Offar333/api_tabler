package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//RoomData exported
func RoomData(w http.ResponseWriter, r *http.Request) {
	var roomData RoomPlayers

	//var room Room

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idMesa := keyVal["ID_MESA"]

	//FIRST --> CHECK IF THERE`S ALREADY A DM AT THE TABLE
	var isThereDm int
	_ = db.QueryRow("SELECT COUNT(*) FROM mesa_jogadores WHERE ID_MESA = ? AND MESTRE_JOGA = 1", idMesa).Scan(&isThereDm)

	if err != nil {

		panic(err.Error())

	}

	if isThereDm == 0 {
		//POPULATE THE STRUCT WHITH THE INFORMATION THAT THERE IS NO DM AVAILABLE
		roomData.DungeonMaster = "NO_DM"
		roomData.TablesJoined = 0

		//TESTING

		roomData.Players = []PlayersInfo{
			{
				PlayerName:  "Teste_NOMEJOGADOR",
				PlayerChar:  "Teste_CHAR",
				PlayerClass: "Teste_CLASSE",
			},
			{
				PlayerName:  "Teste_NOMEJOGADOR 2",
				PlayerChar:  "Teste_CHAR 2",
				PlayerClass: "Teste_CLASSE 2",
			},
			{
				PlayerName:  "Teste_NOMEJOGADOR 3",
				PlayerChar:  "Teste_CHAR 3",
				PlayerClass: "Teste_CLASSE 3",
			},
		}
		/* roomData.Players.PlayerName = "Teste_NOMEJOGADOR"
		roomData.Players.PlayerChar = "Teste_CHAR"
		roomData.Players.PlayerClass = "Teste_CLASSE" */
		//------------------------------------------------

	} else if isThereDm != 0 {

		var idDm int

		//POPULATE THE STRUCT WITH THE DM INFO
		//RETURNS THE NAME OF THE DM AT THE TABLE
		dmInfo, err := db.Query("SELECT a.ID_USUAR, a.NOME_USUAR AS MESTRE FROM usuario a "+
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

		//TESTING
		/* roomData.Players.PlayerName = "Teste_NOMEJOGADOR"
		roomData.Players.PlayerChar = "Teste_CHAR"
		roomData.Players.PlayerClass = "Teste_CLASSE" */
		//------------------------------------------------

	}

	json.NewEncoder(w).Encode(roomData)
	w.WriteHeader(http.StatusOK)

}
