package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jerminb/gilgab/config"
	"github.com/jerminb/gilgab/models"
	"github.com/jerminb/gilgab/repositories"
	"github.com/julienschmidt/httprouter"
)

//Controller is the interface that all the conceret controllers will have to realize
type Controller interface {
	GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetPropertyByID(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Post(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

//BaseController is the base class for gilgab controllers
type BaseController struct {
	Repository repositories.MongoRepository
}

//SetHTTPContentType sets content type of a http response object
func (bc *BaseController) SetHTTPContentType(w *http.ResponseWriter, httpContentType string) {
	(*w).Header().Set("Content-Type", httpContentType)
}

//SetHTTPCode sets the return HTTP code (200 for success).
func (bc *BaseController) SetHTTPCode(w *http.ResponseWriter, code int) {
	(*w).WriteHeader(code)
}

//UserController is the controller class for User model
type UserController struct {
	BaseController
}

//GetByID returns a User by its ID
func (uc *UserController) GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, err := uc.Repository.FindByID(p.ByName("id"))
	if err != nil {
		uc.SetHTTPCode(&w, config.HTTPCode500)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	// Marshal provided interface into JSON structure
	uj, err := json.Marshal(user)
	if err != nil {
		uc.SetHTTPCode(&w, config.HTTPCode500)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	// Write content-type, statuscode, payload
	uc.SetHTTPContentType(&w, config.HTTPContentTypeJSON)
	uc.SetHTTPCode(&w, config.HTTPCode200)
	fmt.Fprintf(w, "%s", uj)
}

//GetPropertyByID returns a StoryView by ID
func (uc *UserController) GetPropertyByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, err := uc.Repository.FindByID(p.ByName("id"))
	if err != nil {
		uc.SetHTTPCode(&w, config.HTTPCode500)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	userModel := user.(models.User)
	if userModel.StoryViews != nil {
		uc.SetHTTPCode(&w, config.HTTPCode500)
		return
	}

	for _, sv := range userModel.StoryViews {
		if sv.StoryID == p.ByName("s_id") {
			svj, errSV := json.Marshal(sv)
			if errSV != nil {
				uc.SetHTTPCode(&w, config.HTTPCode500)
				fmt.Fprintf(w, "%s", errSV.Error())
				return
			}

			// Write content-type, statuscode, payload
			uc.SetHTTPContentType(&w, config.HTTPContentTypeJSON)
			uc.SetHTTPCode(&w, config.HTTPCode200)
			fmt.Fprintf(w, "%s", svj)
			return
		}
	}

	// Write content-type, statuscode, payload
	uc.SetHTTPContentType(&w, config.HTTPContentTypeJSON)
	uc.SetHTTPCode(&w, config.HTTPCode404)
}

//Post creates a User Object
func (uc *UserController) Post(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body
	u := models.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)
	id, err := uc.Repository.Insert(u)
	if err != nil {
		uc.SetHTTPCode(&w, config.HTTPCode500)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
	// Write content-type, statuscode, payload
	uc.SetHTTPContentType(&w, config.HTTPContentTypeJSON)
	uc.SetHTTPCode(&w, config.HTPPCode201)
	fmt.Fprintf(w, "%s", id)
}

//NewControllerForType is the factory for Controller Realization classes
func NewControllerForType(controllerType string) Controller {
	switch controllerType {
	case config.ControllerTypeUser:
		return &UserController{
			BaseController{
				Repository: repositories.NewRepositoryForType(config.RepositoryTypeUser),
			},
		}
	}
	return nil
}
