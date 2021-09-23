package entities

import (
	"crud/database/mongodb"
)

// App is the structure for the app
type App struct {
	mongodb.MongoID `bson:",inline"`
	Name            string    `json:"name,omitempty" bson:"name,omitempty"`
	Icon            string    `json:"icon,omitempty" bson:"icon,omitempty"`
	Package         string    `json:"package,omitempty" bson:"package,omitempty"`
	BundleID        string    `json:"bundle_id,omitempty" bson:"bundle_id,omitempty"`
	Description     string    `json:"description,omitempty" bson:"description,omitempty"`
	Versions        []Version `json:"versions,omitempty" bson:"versions,omitempty"`
}

// Version is to save different application versions
type Version struct {
	AppID              string `json:"app_id,omitempty" bson:"app_id,omitempty"`
	VersionString      string `json:"versionstring,omitempty" bson:"versionstring,omitempty"`
	AndroidCompilation string `json:"androidcompilation,omitempty" bson:"androidcompilation,omitempty"`
	IOsCompilation     string `json:"ioscompilation,omitempty" bson:"ioscompilation,omitempty"`
	ReleaseNotes       string `json:"releasenotes,omitempty" bson:"releasenotes,omitempty"`
}
