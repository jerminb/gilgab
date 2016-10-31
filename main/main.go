package main

import (
	"net/http"

	"github.com/jerminb/gilgab/config"
	"github.com/jerminb/gilgab/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()
	// Get a user resource
	/*r.GET("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Stub an example user
		u := models.User{
			FirstName: "Bob",
			LastName:  "Smith",
			Gender:    "male",
			Age:       50,
			ID:        p.ByName("id"),
		}

		// Marshal provided interface into JSON structure
		uj, _ := json.Marshal(u)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", uj)
	})*/
	uc := controllers.NewControllerForType(config.ControllerTypeUser)
	r.GET("/users/:id", uc.GetByID)
	r.GET("/users/:id/story_views/:s_id", uc.GetPropertyByID)
	r.POST("/users", uc.Post)
	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}
