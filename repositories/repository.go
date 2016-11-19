package repositories

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/jerminb/gilgab/config"
	"github.com/jerminb/gilgab/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var errMalformedID = errors.New("ID is not in bson hex format")

//MongoRepository is the interface signature for concrete repositories
type MongoRepository interface {
	FindByID(ID string) (interface{}, error)
	//FindByFilter(filter string) ([]interface{}, error)
	Insert(entity interface{}) (string, error)
	Update(id string, entity interface{}) error
}

//BaseMongoRepository is the collection of all the mongo repository functions with commonalities lumped in
type BaseMongoRepository struct {
	session *mgo.Session
}

//FindByID return is a User using its ID
func (bmr *BaseMongoRepository) FindByID(ID string, model models.GilGabModel, collection string) (interface{}, error) {
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(ID) {
		return nil, errMalformedID
	}

	// Grab id
	oid := bson.ObjectIdHex(ID)
	if err := bmr.session.DB(config.DBName).C(collection).FindId(oid).One(model); err != nil {
		return nil, err
	}

	return model, nil
}

//Insert adds a user to the database
func (bmr *BaseMongoRepository) Insert(entity models.GilGabModel, collection string) (string, error) {
	// Add an Id
	entity.SetID(bson.NewObjectId())
	entity.SetCreatedAt(time.Now())
	err := bmr.session.DB(config.DBName).C(collection).Insert(entity)
	if err != nil {
		return "", err
	}
	return entity.GetID().Hex(), nil
}

//Update updates an object
func (bmr *BaseMongoRepository) Update(ID string, entity models.GilGabModel, collection string) error {
	if !bson.IsObjectIdHex(entity.GetID().String()) {
		// Verify id is ObjectId, otherwise bail
		if !bson.IsObjectIdHex(ID) {
			return errMalformedID
		}
		entity.SetID(bson.ObjectIdHex(ID))
	}
	je, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	me := make(map[string]interface{})
	err = json.Unmarshal(je, &me)
	if err != nil {
		return err
	}
	change := bson.M{"$set": bson.M(me)}
	err = bmr.session.DB(config.DBName).C(collection).UpdateId(entity.GetID(), change)
	if err != nil {
		return err
	}
	return nil
}

//NewRepositoryForType is the factory class for MongoRepository Realization classes
func NewRepositoryForType(repoType string) MongoRepository {
	// Connect to our local mongo
	s, err := mgo.Dial(config.DBConnectionString)

	// Check if connection error, is mongo running?
	if err != nil {
	}
	switch repoType {
	case config.RepositoryTypeUser:
		return &UserMongoRepository{
			bmr: &BaseMongoRepository{s},
		}
	case config.RepositoryTypeStory:
		return &StoryMongoRepository{
			bmr: &BaseMongoRepository{s},
		}
	}

	return nil
}
