package controllers

import (
	"fmt"
	"github.com/ssjlee93/fitworks-data-user/dtos"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ssjlee93/fitworks-data-user/repositories"
)

var validPath = regexp.MustCompile("^/(user|users)/([0-9]+)?$")

type UserController struct {
	r repositories.Repository[dtos.User]
}

func NewUserController(repo repositories.UserRepository) *UserController {
	return &UserController{r: &repo}
}

func (userController *UserController) ReadAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UserController.ReadAllHandler called")
	res, _ := userController.r.ReadAll()
	marshalResponse(res, w)
}

func (userController *UserController) readOneHandler(w http.ResponseWriter, r *http.Request, id int64) {
	log.Println("UserController.readOneHandler called")
	res, _ := userController.r.ReadOne(id)
	marshalResponse(res, w)
}

func (userController *UserController) createHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UserController.createHandler called")
	user, err := unmarshalRequest(w, r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = userController.r.Create(*user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusCreated)
}

func (userController *UserController) updateHandler(w http.ResponseWriter, r *http.Request, id int64) {
	log.Println("UserController.updateHandler called")
	user, err := unmarshalRequest(w, r)
	user.UserID = id
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = userController.r.Update(*user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusNoContent)
}

func (userController *UserController) deleteHandler(w http.ResponseWriter, r *http.Request, id int64) {
	log.Println("| UserController.deleteHandler called")
	err := userController.r.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusNoContent)
}

func (userController *UserController) Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("UserController.Handler called")
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

// 36  func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
// 37  	p, err := loadPage(title)
// 38  	if err != nil {
// 39  		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
// 40  		return
// 41  	}
// 42  	renderTemplate(w, "view", p)
// 43  }
// 44
// 45  func editHandler(w http.ResponseWriter, r *http.Request, title string) {
// 46  	p, err := loadPage(title)
// 47  	if err != nil {
// 48  		p = &Page{Title: title}
// 49  	}
// 50  	renderTemplate(w, "edit", p)
// 51  }
// 52
// 53  func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
// 54  	body := r.FormValue("body")
// 55  	p := &Page{Title: title, Body: []byte(body)}
// 56  	err := p.save()
// 57  	if err != nil {
// 58  		http.Error(w, err.Error(), http.StatusInternalServerError)
// 59  		return
// 60  	}
// 61  	http.Redirect(w, r, "/view/"+title, http.StatusFound)
// 62  }
