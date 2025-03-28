package controllers

import "net/http"

type Controller interface {
	ReadAllHandler(w http.ResponseWriter, r *http.Request)
	Handler(w http.ResponseWriter, r *http.Request)
}
