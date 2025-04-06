package models

type User struct {
	ID                 string         `json:"_id"`
	AccessRoles        []string       `json:"accessRoles"`
	Avatar             Avatar         `json:"avatar"`
	Badge              []any          `json:"badge"`
	Bot                any            `json:"bot,omitempty"`
	Companies          []UserCompany  `json:"companies"`
	CreatedAt          string         `json:"createdAt"`
	Devices            Devices        `json:"devices"`
	Email              string         `json:"email"`
	EmailIsVerified    bool           `json:"emailIsVerified"`
	Extra              Extra          `json:"extra"`
	HierarchyLevel     string         `json:"hierarchyLevel"`
	IsActive           bool           `json:"isActive"`
	IsOnline           bool           `json:"isOnline"`
	IsPhoneVerified    bool           `json:"isPhoneVerified"`
	Job                string         `json:"job"`
	JobTitle           string         `json:"jobTitle"`
	LastRequestDate    string         `json:"lastRequestDate"`
	MessagesUnread     []any          `json:"messagesUnread"`
	ModifiedAt         string         `json:"modifiedAt"`
	Name               Name           `json:"name"`
	NeedPasswordChange bool           `json:"needPasswordChange"`
	Notifications      Notifications  `json:"notifications"`
	Permissions        []any          `json:"permissions"`
	Phone              string         `json:"phone"`
	ProfileInfo        []any          `json:"profileInfo"`
	Properties         []string       `json:"properties"`
	Provider           string         `json:"provider"`
	RequiredChanges    []any          `json:"requiredChanges"`
	Role               string         `json:"role"`
	Search             []string       `json:"search"`
	Settings           Settings       `json:"settings"`
	TaskManager        bool           `json:"taskManager"`
	TermsConditions    bool           `json:"termsConditions"`
	Extensions         map[string]any `json:"extensions"`
	PermissionsV2      []string       `json:"permissionsV2"`
}

type Settings struct {
	HideSummary  bool `json:"hideSummary"`
	HideContacts bool `json:"hideContacts"`
}

type Notifications struct {
	TurnOffChannel []any  `json:"turnOffChannel"`
	TurnOffGroup   []any  `json:"turnOffGroup"`
	Work           []Work `json:"work"`
	General        string `json:"general"`
}

type Work struct {
	Day    string `json:"day"`
	Active bool   `json:"active"`
	Start  string `json:"start"`
	End    string `json:"end"`
}

type Name struct {
	SecondLastName string `json:"secondLastName"`
	LastName       string `json:"lastName"`
	Names          string `json:"names"`
	DisplayName    string `json:"displayName"`
}

type Extra struct {
}

type Devices struct {
	IPhone  []any `json:"iphone"`
	Android []any `json:"android"`
	Web     bool  `json:"web"`
	Web2    any   `json:"web2,omitempty"`
}

type UserCompany struct {
	CompanyID string    `json:"companyId"`
	ID        string    `json:"_id"`
	Hierarchy Hierarchy `json:"hierarchy"`
}

type Hierarchy struct {
	Subordinate []any `json:"subordinate"`
	Peers       []any `json:"peers"`
	Boss        []any `json:"boss"`
}

type Avatar struct {
	Small    string `json:"small"`
	Original string `json:"original"`
	Square   string `json:"square"`
}

type UserQuery struct {
	IsActive *bool   `json:"isActive,omitempty"`
	JobTitle *string `json:"jobTitle,omitempty"`
}

type UserActivity struct {
	Date             string                       `json:"date"`
	CustomStatus     *UserActivityCustomStatus    `json:"customStatus,omitempty"`
	ConnectionStatus UserActivityConnectionStatus `json:"connectionStatus"`
}

type UserActivityCustomStatus struct {
	Display         *string `json:"display,omitempty"`
	DurationMinutes *string `json:"durationMinutes,omitempty"`
	Start           *string `json:"start,omitempty"`
	Type            *string `json:"type,omitempty"`
	End             *string `json:"end,omitempty"`
}

type UserActivityConnectionStatus struct {
	Online     *bool   `json:"online,omitempty"`
	LastOnline *string `json:"lastOnline,omitempty"`
}
