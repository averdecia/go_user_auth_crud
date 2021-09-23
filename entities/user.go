package entities

import (
	"crud/database/mongodb"
	"encoding/json"
)

// User is the structure for the user
type User struct {
	mongodb.MongoID `bson:",inline"`
	Firstname       string      `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname        string      `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Phone           int         `json:"phone,omitempty" bson:"phone,omitempty"`
	Email           string      `json:"email,omitempty" bson:"email,omitempty"`
	Password        string      `json:"password,omitempty" bson:"password,omitempty"`
	Tokens          string      `json:"tokens,omitempty" bson:"tokens,omitempty"`
	Social          interface{} `json:"social,omitempty" bson:"social,omitempty"`
}

// Auth is used in authentication method
type Auth struct {
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

// MarshalJSON rewrite in order to avoid pasword from being printed
func (u User) MarshalJSON() ([]byte, error) {
	type user User // prevent recursion
	x := user(u)
	x.Password = ""
	return json.Marshal(x)
}

// Activity is the structure for the user
type Activity struct {
	Timestamp     int64   `json:"timestamp"`
	Priority      int     `json:"priority,omitempty"`
	Type          string  `json:"type,omitempty"`
	User          string  `json:"user,omitempty"`
	AffectedUser  string  `json:"affecteduser,omitempty"`
	App           string  `json:"app,omitempty"`
	Subject       string  `json:"subject,omitempty"`
	SubjectParams string  `json:"subjectparams,omitempty"`
	Message       string  `json:"message,omitempty"`
	MessageParams string  `json:"messageparams,omitempty"`
	File          string  `json:"file,omitempty"`
	Link          string  `json:"link,omitempty"`
	ObjectType    string  `json:"object_type,omitempty"`
	ObjectID      float32 `json:"object_id,omitempty"`
}

// ElasticSearchResults is the structure for the Elastic Search results
type ElasticSearchResults struct {
	Took   int                        `json:"took"`
	Shards json.RawMessage            `json:"_shards"`
	Hits   Hits                       `json:"hits"`
	Aggs   map[string]json.RawMessage `json:"aggregations"`
}

// Hits is the structure for the elastic result elements
type Hits struct {
	Total int      `json:"total.value"`
	Hits  []Source `json:"hits"`
}

// Source is the structure for the elastic result elements
type Source struct {
	Source json.RawMessage `json:"_source"`
	Index  string          `json:"_index"`
	ID     string          `json:"_id"`
}

// Session is the structure for the user
type Session struct {
	User         User
	Token        string
	LastTimeUsed int
	Device       string
}
