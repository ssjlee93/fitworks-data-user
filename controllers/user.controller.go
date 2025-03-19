package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ssjlee93/fitworks-data-user/repositories"
)

var validPath = regexp.MustCompile("^(/(user|users)(/[0-9]+)?$)")

type UserController struct {
	r repositories.UserRepository
}

func NewUserController(repo repositories.UserRepository) *UserController {
	return &UserController{r: repo}
}

func (userController *UserController) ReadAllHandler(w http.ResponseWriter, r *http.Request) {
	res, _ := userController.r.ReadAll()
	marshalResponse(res, w)
}

func (userController *UserController) ReadOneHandler(w http.ResponseWriter, r *http.Request, id int64) {
	res, _ := userController.r.ReadOne(id)
	marshalResponse(res, w)
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

func (userController *UserController) Handler(w http.ResponseWriter, r *http.Request) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	fmt.Println(m)

	if m == nil {
		http.NotFound(w, r)
		return
	}
	id, err := strconv.Atoi(m[3][1:])
	if err != nil {
		e := fmt.Errorf("could not parse user id %s", m[3][1:])
		fmt.Println(e)
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		userController.ReadOneHandler(w, r, int64(id))
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	default:
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
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
