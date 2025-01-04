package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ssjlee93/fitworks-data-user/dtos"
	"github.com/ssjlee93/fitworks-data-user/repositories"
)

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

func marshalResponse[T []dtos.User | *dtos.User](res T, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("could not marshal res: %w", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return nil
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

// 73  var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
// 74
// 75  func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
// 76  	return func(w http.ResponseWriter, r *http.Request) {
// 77  		m := validPath.FindStringSubmatch(r.URL.Path)

// 78  		if m == nil {
// 79  			http.NotFound(w, r)
// 80  			return
// 81  		}
// 82  		fn(w, r, m[2])
// 83  	}
// 84  }
