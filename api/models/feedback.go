package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Feedback TODO
type Feedback struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Message   string        `json:"message"`
	Emoji     int           `json:"emoji"`
	Email     string        `json:"email"`
	UserID    string        `json:"user_id"`
	Domain    string        `json:"domain"`
	URL       string        `json:"url"`
	CreatedAt time.Time     `json:"created_at"`
}

// FeedbackList return a list of feedback
func FeedbackList(db *mgo.Database) ([]*Feedback, error) {
	feedback := []*Feedback{}
	c := db.C("feedback")
	if err := c.Find(nil).All(&feedback); err != nil {
		return nil, err
	}
	return feedback, nil
}

// Create a feedback from the object
func (f *Feedback) Insert(db *mgo.Database) error {
	f.CreatedAt = time.Now()
	col := db.C("feedback")

	if err := col.Insert(f); err != nil {
		return err
	}
	return nil
}

// Get return the spefic feedback from the DB
func (f *Feedback) Get(db *mgo.Database) error {
	col := db.C("feedback")
	if err := col.Find(bson.M{"_id": f.Id}).One(&f); err != nil {
		return err
	}
	return nil
}
