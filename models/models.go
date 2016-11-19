package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/jerminb/gilgab/config"
)

//GilGabModel is the interface that works as a template for repository classes
type GilGabModel interface {
	SetID(id bson.ObjectId)
	GetID() bson.ObjectId
	SetCreatedAt(createdAt time.Time)
	GetCreatedAt() time.Time
}

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

//SetID is the setter function for ID
func (u *User) SetID(id bson.ObjectId) {
	u.ID = id
}

//GetID is the getter function for ID
func (u *User) GetID() bson.ObjectId {
	return u.ID
}

//SetCreatedAt is the setter function for CreatedAt
func (u *User) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}

//GetCreatedAt is the getter function for CreatedAt
func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
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
	ID             bson.ObjectId      `json:"id" bson:"_id"`
	Title          string             `json:"title" bson:"title"`
	Content        string             `json:"content" bson:"content"`
	TitleImagePath string             `json:"title_image_path" bson:"title_image_path"`
	RevisionNumber int                `json:"revision_number" bson:"revision_number"`
	IsPublished    bool               `json:"is_published" bson:"is_published"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	LastSavedAt    time.Time          `json:"last_saved_at" bson:"last_saved_at"`
	Sequence       int                `json:"sequence" bson:"sequence"`
	AutheredByID   bson.ObjectId      `json:"authered_by_id" bson:"authered_by_id"`
	ContentType    config.ContentType `json:"content_type" bson:"content_type"`
	UserViews      []UserView         `json:"user_views" bson:"user_views"`
}

//SetID is the setter function for ID
func (sp *StoryPart) SetID(id bson.ObjectId) {
	sp.ID = id
}

//GetID is the getter function for ID
func (sp *StoryPart) GetID() bson.ObjectId {
	return sp.ID
}

//SetCreatedAt is the setter function for CreatedAt
func (sp *StoryPart) SetCreatedAt(createdAt time.Time) {
	sp.CreatedAt = createdAt
}

//GetCreatedAt is the getter function for CreatedAt
func (sp *StoryPart) GetCreatedAt() time.Time {
	return sp.CreatedAt
}

//StoryCover is the representation of a story-cover from domain models
type StoryCover struct {
	Title          string `json:"title" bson:"title"`
	Synopsis       string `json:"synopsis" bson:"synopsis"`
	CoverImagePath string `json:"cover_image_path" bson:"cover_image_path"`
}

//Story is the representation of a story from domain models
type Story struct {
	StoryCover  StoryCover      `json:"story_cover" bson:"story_cover"`
	PartIDs     []bson.ObjectId `json:"part_ids" bson:"part_ids"`
	Tags        []string        `json:"tags" bson:"tags"`
	ID          bson.ObjectId   `json:"id" bson:"_id"`
	CreatedByID bson.ObjectId   `json:"created_by" bson:"created_by"`
	CreatedAt   time.Time       `json:"created_at" bson:"created_at"`
	UserViews   []UserView      `json:"user_views" bson:"user_views"`
}

//SetID is the setter function for ID
func (s *Story) SetID(id bson.ObjectId) {
	s.ID = id
}

//GetID is the getter function for ID
func (s *Story) GetID() bson.ObjectId {
	return s.ID
}

//SetCreatedAt is the setter function for CreatedAt
func (s *Story) SetCreatedAt(createdAt time.Time) {
	s.CreatedAt = createdAt
}

//GetCreatedAt is the getter function for CreatedAt
func (s *Story) GetCreatedAt() time.Time {
	return s.CreatedAt
}

//StoryView encapsulates user story views
type StoryView struct {
	StoryID    bson.ObjectId `json:"story_id" bson:"story_id"`
	LastPartID bson.ObjectId `json:"last_part_id" bson:"last_part_id"`
	StoryCover StoryCover    `json:"story_cover" bson:"story_cover"`
	ViewDate   time.Time     `json:"view_date" bson:"view_date"`
	ViewCount  int           `json:"view_count" bson:"view_count"`
}

//UserView is the reverse of StoryView
type UserView struct {
	UserID    bson.ObjectId `json:"user_id" bson:"user_id"`
	ViewDate  time.Time     `json:"view_date" bson:"view_date"`
	ViewCount int           `json:"view_count" bson:"view_count"`
}
