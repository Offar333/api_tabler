package controllers

import (
	"net/http"
)

//EnableCors exported
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
