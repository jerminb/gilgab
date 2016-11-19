package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jerminb/gilgab/config"
	"github.com/jerminb/gilgab/models"
	"github.com/jerminb/gilgab/repositories"
	"github.com/julienschmidt/httprouter"
)

//StoryController is the controller class for Story model
type StoryController struct {
	BaseController
	PartRepository repositories.MongoRepository
}

//GetByID returns a Story by its ID
func (sc *StoryController) GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("GetByID:%s\n", p.ByName("id"))
	story, err := sc.Repository.FindByID(p.ByName("id"))
	if err != nil {
		sc.SetHTTPCode(&w, config.HTTPCode500)
		sc.WriteError(&w, err.Error())
		return
	}

	// Write content-type, statuscode, payload
	sc.SetHTTPContentType(&w, config.HTTPContentTypeJSON)
	sc.SetHTTPCode(&w, config.HTTPCode200)
	sc.WriteResponse(&w, story)
}

//GetPropertyByID returns a StoryView by ID
func (sc *StoryController) GetPropertyByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	/*user, err := uc.Repository.FindByID(p.ByName("id"))
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
		if sv.StoryID.String() == p.ByName("s_id") {
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
	uc.SetHTTPCode(&w, config.HTTPCode404)*/
}

//Post creates a User Object
func (sc *StoryController) Post(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("Post\n")
	// Stub an user to be populated from the body
	s := models.Story{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&s)
	log.Printf("Story is %v\n", s)
	id, err := sc.Repository.Insert(s)
	if err != nil {
		sc.SetHTTPCode(&w, config.HTTPCode500)
		sc.WriteError(&w, err.Error())
		return
	}
	// Write content-type, statuscode, payload
	sc.SetHTTPContentType(&w, config.HTTPContentTypeJSON)
	sc.SetHTTPCode(&w, config.HTPPCode201)
	sc.WriteID(&w, id)
}

//Put updates a User Object
func (sc *StoryController) Put(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	sid := p.ByName("id")
	// Stub an user to be populated from the body
	s := models.Story{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&s)
	err := sc.Repository.Update(sid, s)
	if err != nil {
		sc.SetHTTPCode(&w, config.HTTPCode500)
		sc.WriteError(&w, err.Error())
		return
	}
	// Write content-type, statuscode, payload
	sc.SetHTTPContentType(&w, config.HTTPContentTypeJSON)
	sc.SetHTTPCode(&w, config.HTPPCode204)
}
