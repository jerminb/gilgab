package repositories

import "github.com/jerminb/gilgab/models"

//StoryPartMongoRepository is the MongoRepository realization for StoryPart objects
type StoryPartMongoRepository struct {
	bmr *BaseMongoRepository
}

//FindByID return is a User using its ID
func (spmr *StoryPartMongoRepository) FindByID(ID string) (interface{}, error) {
	return spmr.bmr.FindByID(ID, &models.StoryPart{}, "story_parts")
}

//Insert adds a user to the database
func (spmr *StoryPartMongoRepository) Insert(entity interface{}) (string, error) {
	sp := entity.(models.StoryPart)
	return spmr.bmr.Insert(&sp, "story_parts")
}
