package services

import (
	"github.com/jerminb/gilgab/models"
	"github.com/jerminb/gilgab/repositories"
	"gopkg.in/mgo.v2/bson"
)

//StoryService provides domain specific functionalities for Story Domain Model
type StoryService struct {
	storyRepo     repositories.MongoRepository
	storyPartRepo repositories.MongoRepository
}

//AddPart adds a new story part to strory identified by storyID
func (ss *StoryService) AddPart(storyID string, part models.StoryPart) (string, error) {
	pid, err := ss.storyPartRepo.Insert(part)
	if err != nil {
		return "", err
	}
	sInterface, err := ss.storyRepo.FindByID(storyID)
	if err != nil {
		return "", err
	}
	s := sInterface.(models.Story)
	if s.PartIDs == nil {
		s.PartIDs = make([]bson.ObjectId, 0)
	}
	oid := bson.ObjectIdHex(pid)
	s.PartIDs = append(s.PartIDs, oid)
	err = ss.storyRepo.Update(storyID, s)
	if err != nil {
		return "", err
	}
	return pid, nil
}
