package controllers

import (
	"fmt"
	"github.com/ssjlee93/fitworks-data-user/models"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ssjlee93/fitworks-data-user/services"
)

var validPath = regexp.MustCompile("^/(user|users)/([0-9]+)?$")

type UserController struct {
	s services.Service[models.User]
}

func NewUserController(svc services.UserService) *UserController {
	return &UserController{s: &svc}
}

func (userController *UserController) ReadAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("| UserController.ReadAllHandler")
	res, _ := userController.s.ReadAll()
	marshalResponse(res, w)
}

func (userController *UserController) readOneHandler(w http.ResponseWriter, r *http.Request, id int64) {
	log.Println("| UserController.readOneHandler")
	res, _ := userController.s.ReadOne(id)
	marshalResponse(res, w)
}

func (userController *UserController) createHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("| UserController.createHandler")
	user, err := unmarshalRequest(w, r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = userController.s.Create(*user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusCreated)
}

func (userController *UserController) updateHandler(w http.ResponseWriter, r *http.Request, id int64) {
	log.Println("| UserController.updateHandler")
	user, err := unmarshalRequest(w, r)
	user.UserID = id
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = userController.s.Update(*user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusNoContent)
}

func (userController *UserController) deleteHandler(w http.ResponseWriter, r *http.Request, id int64) {
	log.Println("| UserController.deleteHandler")
	err := userController.s.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusNoContent)
}

func (userController *UserController) Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("UserController.Handler")
	// validate URL
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		log.Println("UserController.Handler error on valid path")
		http.NotFound(w, r)
		return
	}

	// handle POST method separately
	if r.Method == http.MethodPost {
		userController.createHandler(w, r)
		return
	}

	// prepare id from path param
	id, err := extractId(m[2])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		userController.readOneHandler(w, r, id)
	case http.MethodPut:
		userController.updateHandler(w, r, id)
	case http.MethodDelete:
		userController.deleteHandler(w, r, id)
	default:
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}

func extractId(idStr string) (int64, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		e := fmt.Errorf("could not parse user id %s", idStr)
		fmt.Println(e)
		return -1, err
	}
	return int64(id), nil
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, int64)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		id, err := strconv.Atoi(m[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, int64(id))
	}
}
