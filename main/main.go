package main

import (
	"log"
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
	sc := controllers.NewControllerForType(config.ControllerTypeStory)
	r.GET("/users/:id", uc.GetByID)
	r.GET("/users/:id/story_views/:s_id", uc.GetPropertyByID)
	r.POST("/users", uc.Post)
	r.PUT("/users/:id", uc.Put)
	r.GET("/stories/:id", sc.GetByID)
	r.GET("/stories/:id/story_parts/:s_id", sc.GetPropertyByID)
	r.POST("/stories", sc.Post)
	r.PUT("/stories/:id", sc.Put)
	// Fire up the server
	log.Println("Server is starting on port 3000.")
	http.ListenAndServe("localhost:3000", r)
}
