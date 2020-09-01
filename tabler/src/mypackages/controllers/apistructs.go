package controllers

//Room Struct
type Room struct {
	ID       int    `json:"id"`
	AdmMesa  string `json:"admMesa"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	QtdeJog  string `json:"qtdeJog"`
	Formato  string `json:"formato"`
	Status   int    `json:"status"`
	LvlInic  string `json:"lvlInic"`
	ExpJogo  string `json:"expJogo"`
	LinkChat string `json:"linkChat"`
}

//RoomPlayers Struct
type RoomPlayers struct {
	DungeonMaster string        `json:"dungeonMaster"`
	TablesJoined  int           `json:"tablesJoined"`
	AlreadyIn     string        `json:"alreadyIn"`
	Players       []PlayersInfo `json:"players"`
}

//PlayersInfo Struct
type PlayersInfo struct {
	PlayerName  string `json:"playerName"`
	PlayerChar  string `json:"playerChar"`
	PlayerClass string `json:"playerClass"`
}

//GameFormats Struct
type GameFormats struct {
	DescFormat string `json:"descFormat"`
}

//User Struct
type User struct {
	ID         string `json:"id"`
	Nome       string `json:"nome"`
	Apelido    string `json:"apelido"`
	Email      string `json:"email"`
	AvatarPath string `json:"avatarpath"`
}

//DoesExist Struct
type DoesExist struct {
	JaExiste string `json:"jaExiste"`
}

//CustomResponse Struct
type CustomResponse struct {
	CustomResp string `json:"customResp"`
}
