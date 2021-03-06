package config

//ContentType distinguishes between free and premium contents
type ContentType int

const (
	//FreeContentType is a content that is available for all the users
	FreeContentType ContentType = iota
	//PremiumContentType requires a premium account to view
	PremiumContentType
)

const (
	//HTTPContentTypeJSON is application/json
	HTTPContentTypeJSON string = "application/json"

	//HTTPCode200 is the success http code
	HTTPCode200 int = 200
	//HTTPCode404 is the not found http code
	HTTPCode404 int = 404
	//HTTPCode500 is internal server error http code
	HTTPCode500 int = 500
	//HTPPCode201 is resource created http code
	HTPPCode201 int = 201
	//HTPPCode204 is resource updated http code
	HTPPCode204 int = 204

	//ControllerTypeUser defines a controller type for User objects
	ControllerTypeUser string = "user"
	//ControllerTypeStory defines a controller type for Story objects
	ControllerTypeStory string = "story"

	//RepositoryTypeUser defines a repository type for User objects
	RepositoryTypeUser string = "user"
	//RepositoryTypeStory defines a repository type for story objects
	RepositoryTypeStory string = "story"
	//RepositoryTypeStoryPart defines a repository type for StoryType objects
	RepositoryTypeStoryPart string = "story_part"

	//DBName is the Database name in mongodb
	DBName string = "gilgab"
	//DBConnectionString is the connection string to mongodb
	DBConnectionString string = "mongodb://localhost"
)
