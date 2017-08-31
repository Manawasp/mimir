package models

// Feedback TODO
type Feedback struct {
	Message string `json:"message"`
	Emoji   int    `json:"emoji" gorethink:"emoji"`
}
