package repositories

import "github.com/jerminb/gilgab/models"

//UserMongoRepository is the MongoRepository realization for User objects
type UserMongoRepository struct {
	bmr *BaseMongoRepository
}

//FindByID return is a User using its ID
func (umr *UserMongoRepository) FindByID(ID string) (interface{}, error) {
	return umr.bmr.FindByID(ID, &models.User{}, "users")
}

//Insert adds a user to the database
func (umr *UserMongoRepository) Insert(entity interface{}) (string, error) {
	u := entity.(models.User)
	return umr.bmr.Insert(&u, "users")
}

//Update updates an exisiting user
func (umr *UserMongoRepository) Update(id string, entity interface{}) error {
	u := entity.(models.User)
	return umr.bmr.Update(id, &u, "users")
}
