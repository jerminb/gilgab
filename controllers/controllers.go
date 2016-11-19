package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jerminb/gilgab/config"
	"github.com/jerminb/gilgab/models"
	"github.com/jerminb/gilgab/repositories"
	"github.com/julienschmidt/httprouter"
)

//ErrorResponse encapsulates an error into a json container
type ErrorResponse struct {
	Error string `json:"error"`
}

//IDResponse is encapsulates an ID into a json container
type IDResponse struct {
	ID string `json:"id"`
}

//Controller is the interface that all the conceret controllers will have to realize
type Controller interface {
	GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetPropertyByID(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Post(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Put(w http.ResponseWriter, r *http.Request, p httprouter.Params)
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

//WriteResponse json-marshalizes response and writes into the response pipeline
func (bc *BaseController) WriteResponse(w *http.ResponseWriter, response interface{}) {
	rj, _ := json.Marshal(response)
	fmt.Fprintf(*w, "%s", rj)
}

//WriteID writes id as a post response
func (bc *BaseController) WriteID(w *http.ResponseWriter, id string) {
	idResponse := IDResponse{
		ID: id,
	}
	bc.WriteResponse(w, idResponse)
}

//WriteError writes error as a post response
func (bc *BaseController) WriteError(w *http.ResponseWriter, err string) {
	errResponse := ErrorResponse{
		Error: err,
	}
	bc.WriteResponse(w, errResponse)
}

//UserController is the controller class for User model
type UserController struct {
	BaseController
}

//GetByID returns a User by its ID
func (uc *UserController) GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("GetByID:%s\n", p.ByName("id"))
	user, err := uc.Repository.FindByID(p.ByName("id"))
	if err != nil {
		uc.SetHTTPCode(&w, config.HTTPCode500)
		uc.WriteError(&w, err.Error())
		return
	}

	// Write content-type, statuscode, payload
	uc.SetHTTPContentType(&w, config.HTTPContentTypeJSON)
	uc.SetHTTPCode(&w, config.HTTPCode200)
	uc.WriteResponse(&w, user)
}

//GetPropertyByID returns a StoryView by ID
func (uc *UserController) GetPropertyByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, err := uc.Repository.FindByID(p.ByName("id"))
	if err != nil {
		uc.SetHTTPCode(&w, config.HTTPCode500)
		uc.WriteError(&w, err.Error())
		return
	}

	userModel := user.(models.User)
	if userModel.StoryViews != nil {
		uc.SetHTTPCode(&w, config.HTTPCode500)
		return
	}

	for _, sv := range userModel.StoryViews {
		if sv.StoryID.String() == p.ByName("s_id") {
			// Write content-type, statuscode, payload
			uc.SetHTTPContentType(&w, config.HTTPContentTypeJSON)
			uc.SetHTTPCode(&w, config.HTTPCode200)
			uc.WriteResponse(&w, sv)
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
		uc.WriteError(&w, err.Error())
		return
	}
	// Write content-type, statuscode, payload
	uc.SetHTTPContentType(&w, config.HTTPContentTypeJSON)
	uc.SetHTTPCode(&w, config.HTPPCode201)
	uc.WriteID(&w, id)
}

//Put updates a User Object
func (uc *UserController) Put(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid := p.ByName("id")
	// Stub an user to be populated from the body
	u := models.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)
	err := uc.Repository.Update(uid, u)
	if err != nil {
		uc.SetHTTPCode(&w, config.HTTPCode500)
		uc.WriteError(&w, err.Error())
		return
	}
	// Write content-type, statuscode, payload
	uc.SetHTTPContentType(&w, config.HTTPContentTypeJSON)
	uc.SetHTTPCode(&w, config.HTPPCode204)
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
	case config.ControllerTypeStory:
		return &StoryController{
			BaseController{
				Repository: repositories.NewRepositoryForType(config.RepositoryTypeStory),
			},
			nil,
		}
	}
	return nil
}
