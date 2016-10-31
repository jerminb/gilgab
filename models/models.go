package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/jerminb/gilgab/config"
)

//User is the representation of a user from domain models
type User struct {
	FirstName  string        `json:"first_name" bson:"first_name"`
	LastName   string        `json:"last_name" bson:"last_name"`
	Gender     string        `json:"gender" bson:"gender"`
	Age        int           `json:"age" bson:"age"`
	ID         bson.ObjectId `json:"id" bson:"_id"`
	CreatedAt  time.Time     `json:"created_at" bson:"created_at"`
	StoryViews []StoryView   `json:"story_views" bson:"story_views"`
}

//Comment is the representation of a comment from domain models
type Comment struct {
	ID           bson.ObjectId
	UserID       bson.ObjectId
	UserFullName string
	StoryID      bson.ObjectId
	CreatedAt    time.Time
	Content      string
}

//StoryPart is the representation of a story-part from domain models
type StoryPart struct {
	ID             bson.ObjectId
	Title          string
	Content        string
	TitleImagePath string
	RevisionNumber int
	IsPublished    bool
	CreatedAt      time.Time
	LastSavedAt    time.Time
	Sequence       int
	AutheredByID   bson.ObjectId
	ContentType    config.ContentType
}

//StoryCover is the representation of a story-cover from domain models
type StoryCover struct {
	Title          string `json:"title" bson:"title"`
	Synopsis       string `json:"synopsis" bson:"synopsis"`
	CoverImagePath string `json:"cover_image_path" bson:"cover_image_path"`
}

//Story is the representation of a story from domain models
type Story struct {
	StoryCover  StoryCover
	PartIDs     []bson.ObjectId
	Tags        []string
	ID          bson.ObjectId
	CreatedByID bson.ObjectId
}

//StoryView encapsulates user story views
type StoryView struct {
	StoryID    string     `json:"story_id" bson:"story_id"`
	LastPartID string     `json:"last_part_id" bson:"last_part_id"`
	StoryCover StoryCover `json:"story_cover" bson:"story_cover"`
	ViewDate   time.Time  `json:"view_date" bson:"view_date"`
	ViewCount  int        `json:"view_count" bson:"view_count"`
}
