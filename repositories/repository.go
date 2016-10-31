package repositories

import (
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
}

//UserMongoRepository is the MongoRepository realization for User objects
type UserMongoRepository struct {
	session *mgo.Session
}

//FindByID return is a User using its ID
func (umr *UserMongoRepository) FindByID(ID string) (interface{}, error) {
	/*return models.User{
		FirstName: "Bob",
		LastName:  "Smith",
		Gender:    "male",
		Age:       50,
		ID:        ID,
	}, nil*/
	// Grab id

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(ID) {
		return nil, errMalformedID
	}

	// Grab id
	oid := bson.ObjectIdHex(ID)

	// Stub user
	u := models.User{}
	if err := umr.session.DB(config.DBName).C("users").FindId(oid).One(&u); err != nil {
		return nil, err
	}

	return u, nil
}

//Insert adds a user to the database
func (umr *UserMongoRepository) Insert(entity interface{}) (string, error) {
	u := entity.(models.User)
	// Add an Id
	u.ID = bson.NewObjectId()
	u.CreatedAt = time.Now()
	err := umr.session.DB(config.DBName).C("users").Insert(u)
	if err != nil {
		return "", err
	}
	return u.ID.String(), nil
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
			session: s,
		}
	}
	return nil
}
