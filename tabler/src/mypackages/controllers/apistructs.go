package controllers

//Room Struct
type Room struct {
	ID      int    `json:"id"`
	AdmMesa string `json:"admMesa"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	QtdeJog int    `json:"qtdeJog"`
	Formato string `json:"formato"`
	Status  int    `json:"status"`
}
