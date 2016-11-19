package repositories

import "github.com/jerminb/gilgab/models"

//StoryMongoRepository is the MongoRepository realization for Story objects
type StoryMongoRepository struct {
	bmr *BaseMongoRepository
}

//FindByID return is a Story using its ID
func (smr *StoryMongoRepository) FindByID(ID string) (interface{}, error) {
	return smr.bmr.FindByID(ID, &models.Story{}, "stories")
}

//Insert adds a story to the database
func (smr *StoryMongoRepository) Insert(entity interface{}) (string, error) {
	s := entity.(models.Story)
	return smr.bmr.Insert(&s, "stories")
}

//Update updates an exisiting story
func (smr *StoryMongoRepository) Update(id string, entity interface{}) error {
	s := entity.(models.Story)
	return smr.bmr.Update(id, &s, "stories")
}
