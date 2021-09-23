package entities

import "crud/database/mongodb"

// VersionStatus is used to know the config status, in order to lock it after release
type VersionStatus int

const (
	// Build status is used before release
	Build VersionStatus = iota + 1
	// Released status is used after release
	Released
)

// ViewType is used to know the type of view, to allow different components
type ViewType int

const (
	// Splash is used on the app presentation page
	Splash ViewType = iota + 1
	// Login is used on the app login page
	Login
	// List is used to show elements on a list design
	List
	// CardList is used to show elements on a card design
	CardList
	// Form is used to show inputs elements to receive data
	Form
)

// LayoutType is used to know the type of layout, where the menu will be etc...
type LayoutType int

const (
	// LeftMenu applications with menu showing from the left
	LeftMenu LayoutType = iota + 1
	// BottomMenu applications with menu showing on the bottom side
	BottomMenu
	// TopRightDots applications, will have access to specific pages through 3dots button on the top
	TopRightDots
)

// AuthType Api Authentication type
type AuthType int

const (
	// Basic used if user and password is provided
	Basic AuthType = iota + 1
	// Bearer used if token authentication is provided
	Bearer
)

// DataSourceType is used to know the origin of data to fullfil a view/component
type DataSourceType int

const (
	// FromAPI will be used to get data from a specific API
	FromAPI DataSourceType = iota + 1
	// FromJSON is used when data is statically sent
	FromJSON
)

// ComponentType is used to identify the component type
type ComponentType int

const (
	// Button element
	Button ComponentType = iota + 1
	// Popup element
	Popup
	// Text element
	Text
	// Input element
	Input
	// Image element
	Image
	// Frame element
	Frame
	// CloseSession element
	CloseSession
)

// EventType is used to execute an action on a user interaction
type EventType int

const (
	// OnClick on click event
	OnClick EventType = iota + 1
	// OnScroll User scroll
	OnScroll
	// OnBack User back button
	OnBack
	// SlideLeft slide to the left
	SlideLeft
	// SlideRight slide to the right
	SlideRight
)

// ActionType is used to execute an action on an event
type ActionType int

const (
	// GotoPage Navigate to a view element
	GotoPage ActionType = iota + 1
	// GetNext move to the next element
	GetNext
	// RequestAPI sending current element and action button
	RequestAPI
)

// Position is where a component should be located
type Position int

const (
	// TopLeft position
	TopLeft Position = iota + 1
	// TopCenter position
	TopCenter
	// TopRight position
	TopRight
	// CenterLeft position
	CenterLeft
	// CenterCenter position
	CenterCenter
	// CenterRight position
	CenterRight
	// BottomLeft position
	BottomLeft
	// BottomCenter position
	BottomCenter
	// BottomRight position
	BottomRight
)

// MobileConfig is the structure to save all new mobiles configurations
type MobileConfig struct {
	mongodb.MongoID `bson:",inline"`
	UserID          string        `json:"user_id" bson:"user_id" binding:"required"`
	AppID           string        `json:"app_id" bson:"app_id" binding:"required"`
	VersionID       string        `json:"version_id" bson:"version_id" binding:"required"`
	Status          VersionStatus `json:"status,omitempty" bson:"status,omitempty"`
	Views           []View        `json:"views,omitempty" bson:"views,omitempty"`
	Type            LayoutType    `json:"type" bson:"type" binding:"required"`
	Layout          Layout        `json:"layout" bson:"layout" binding:"required"`
}

// Layout is the structure for the main elements layout
type Layout struct {
	Menu           Menu   `json:"menu,omitempty" bson:"menu,omitempty"`
	PrimaryColor   string `json:"primary,omitempty" bson:"primary,omitempty"`
	SecondaryColor string `json:"secondary,omitempty" bson:"secondary,omitempty"`
	Logo           string `json:"logo,omitempty" bson:"logo,omitempty"`
}

// Menu is the list of menu elements
type Menu struct {
	Items []MenuItem `json:"items,omitempty" bson:"items,omitempty"`
}

// MenuItem is the element inside a menu
type MenuItem struct {
	Components []Component `json:"components,omitempty" bson:"components,omitempty"`
}

// View is the structure to save all new mobiles views with its own components
type View struct {
	ID         string      `json:"id" bson:"id"`
	Type       ViewType    `json:"type" bson:"type" binding:"required"`
	Components []Component `json:"components,omitempty" bson:"components,omitempty"`
	DataSource DataSource  `json:"datasource,omitempty" bson:"datasource,omitempty"`
}

// DataSource is the structure to know how the data will be provided
type DataSource struct {
	Type   DataSourceType `json:"type" bson:"type" binding:"required"`
	API    API            `json:"api,omitempty" bson:"api,omitempty"`
	Static string         `json:"static,omitempty" bson:"static,omitempty"`
}

// API is the structure to save all API configurations
type API struct {
	External  bool              `json:"external" bson:"external" binding:"required"`
	URL       string            `json:"url" bson:"url" binding:"required"`
	AuthType  AuthType          `json:"auth_type,omitempty" bson:"api,omitempty"`
	AuthUser  string            `json:"auth_user,omitempty" bson:"auth_user,omitempty"`
	AuthPass  string            `json:"auth_pass,omitempty" bson:"auth_pass,omitempty"`
	AuthToken string            `json:"auth_token,omitempty" bson:"auth_token,omitempty"`
	Method    string            `json:"method,omitempty" bson:"method,omitempty" binding:"required"`
	DataPath  string            `json:"data_path,omitempty" bson:"data_path,omitempty"`
	Mapping   map[string]string `json:"mapping,omitempty" bson:"mapping,omitempty"`
}

// Component is used to save specific components configutrations
type Component struct {
	Name           string        `json:"name,omitempty" bson:"name,omitempty"`
	Type           ComponentType `json:"type" bson:"type" binding:"required"`
	Content        string        `json:"content" bson:"content" binding:"required"`
	PrimaryColor   string        `json:"primary,omitempty" bson:"primary,omitempty"`
	SecondaryColor string        `json:"secondary,omitempty" bson:"secondary,omitempty"`
	Position       Position      `json:"position,omitempty" bson:"position,omitempty"`
	Childs         []Component   `json:"childs,omitempty" bson:"childs,omitempty"`
	Events         []Event       `json:"events,omitempty" bson:"events,omitempty"`
}

// Event is used to trigger events on the components an views
type Event struct {
	Type   EventType `json:"type" bson:"type" binding:"required"`
	Action Action    `json:"action" bson:"action" binding:"required"`
}

// Action is used to trigger an action after an event
type Action struct {
	Type ActionType `json:"type" bson:"type" binding:"required"`
	// GoToView is the view id to be redirected
	GoToView      string  `json:"go_to_view,omitempty" bson:"go_to_view,omitempty"`
	AfterFinished *Action `json:"after,omitempty" bson:"after,omitempty"`
}
