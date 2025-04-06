package models

type Task struct {
	ID              string                 `json:"_id"`
	Editors         []string               `json:"editors"`
	Followers       []string               `json:"followers"`
	Visibility      []string               `json:"visibility"`
	Serial          interface{}            `json:"serial"` // number | number[]
	ClosedAt        string                 `json:"closedAt"`
	ModifiedStateAt string                 `json:"modifiedStateAt"`
	ChannelType     string                 `json:"channelType"` // 'bound' | 'unbound' | 'unbound-hierarchy'
	Status1         string                 `json:"status1,omitempty"`
	Status2         string                 `json:"status2,omitempty"`
	Status3         string                 `json:"status3,omitempty"`
	Status4         string                 `json:"status4,omitempty"`
	Status5         string                 `json:"status5,omitempty"`
	Answers         []string               `json:"answers"`
	Asset           string                 `json:"asset,omitempty"`
	Assignee        string                 `json:"assignee"`
	Validators      []string               `json:"validators"`
	Owner           string                 `json:"owner"`
	UserList        []string               `json:"userList,omitempty"`
	ModifiedAt      string                 `json:"modifiedAt"`
	CreatedAt       string                 `json:"createdAt"`
	CreatedBy       string                 `json:"createdBy"`
	ProjectCode     string                 `json:"projectCode"`
	Indentation     int                    `json:"indentation"`
	Parent          string                 `json:"parent,omitempty"`
	RelativeWeight  float64                `json:"relativeWeight"`
	Weight          float64                `json:"weight"`
	Name            string                 `json:"name"`
	Company         string                 `json:"company"`
	TaskGroup       string                 `json:"taskGroup"`
	Survey          string                 `json:"survey"`
	Child           []string               `json:"child"`
	IsActive        bool                   `json:"isActive"`
	IsValid         bool                   `json:"isValid"`
	Channel         string                 `json:"channel,omitempty"`
	ActiveSlas      []string               `json:"activeSlas"`
	Info            string                 `json:"info"`
	EstimatedTime   *int                   `json:"estimatedTime,omitempty"`
	SmStateMachine  string                 `json:"smStateMachine"`
	SmState         string                 `json:"smState"`
	Status          string                 `json:"status"`
	StartDate       *string                `json:"startDate,omitempty"`
	ResolutionDate  *string                `json:"resolutionDate,omitempty"`
	EndDate         *string                `json:"endDate,omitempty"`
	LastMessage     *TaskLastMessage       `json:"lastMessage,omitempty"`
	Color           string                 `json:"color"` // 'none' | 'red' | 'blue' | 'green' | 'yellow'
	Extensions      map[string]interface{} `json:"extensions"`
}

type TaskLastMessage struct {
	Content string `json:"content"`
	Date    string `json:"date"`
}

type TaskPostData struct {
	TaskGroup     string                 `json:"taskGroup"`
	Name          string                 `json:"name"`
	UserList      []string               `json:"userList,omitempty"`
	Assignee      string                 `json:"assignee"`
	Followers     []string               `json:"followers,omitempty"`
	Editors       []string               `json:"editors,omitempty"`
	StartDate     *string                `json:"startDate,omitempty"`
	EndDate       *string                `json:"endDate,omitempty"`
	EstimatedTime *int                   `json:"estimatedTime,omitempty"`
	Channel       string                 `json:"channel,omitempty"`
	Status1       string                 `json:"status1,omitempty"`
	Status2       string                 `json:"status2,omitempty"`
	Status3       string                 `json:"status3,omitempty"`
	Status4       string                 `json:"status4,omitempty"`
	Status5       string                 `json:"status5,omitempty"`
	Extensions    map[string]interface{} `json:"extensions,omitempty"`
}

type TaskPatchData struct {
	Name          string                 `json:"name,omitempty"`
	UserList      []string               `json:"userList,omitempty"`
	Assignee      string                 `json:"assignee,omitempty"`
	Followers     []string               `json:"followers,omitempty"`
	Editors       []string               `json:"editors,omitempty"`
	StartDate     *string                `json:"startDate,omitempty"`
	EndDate       *string                `json:"endDate,omitempty"`
	EstimatedTime *int                   `json:"estimatedTime,omitempty"`
	Channel       string                 `json:"channel,omitempty"`
	Status1       string                 `json:"status1,omitempty"`
	Status2       string                 `json:"status2,omitempty"`
	Status3       string                 `json:"status3,omitempty"`
	Status4       string                 `json:"status4,omitempty"`
	Status5       string                 `json:"status5,omitempty"`
	Extensions    map[string]interface{} `json:"extensions,omitempty"`
	IsActive      *bool                  `json:"isActive,omitempty"`
	SmState       string                 `json:"smState,omitempty"`
	Visibility    []string               `json:"visibility,omitempty"`
}

type TaskQuery struct {
	IsActive  *bool       `json:"isActive,omitempty"`
	ID        interface{} `json:"_id,omitempty"` // string | { $in: string[] }
	Assignee  string      `json:"assignee,omitempty"`
	Serial    interface{} `json:"serial,omitempty"`    // number | number[]
	UserList  interface{} `json:"userList,omitempty"`  // string | { $in: string[] }
	Channel   interface{} `json:"channel,omitempty"`   // string | { $in: string[] }
	SmState   interface{} `json:"smState,omitempty"`   // string | { $ne: string }
	Status    interface{} `json:"status,omitempty"`    // string | { $in: string[] }
	Status1   interface{} `json:"status1,omitempty"`   // string | { $in: string[] }
	Status2   interface{} `json:"status2,omitempty"`   // string | { $in: string[] }
	Status3   interface{} `json:"status3,omitempty"`   // string | { $in: string[] }
	Status4   interface{} `json:"status4,omitempty"`   // string | { $in: string[] }
	Status5   interface{} `json:"status5,omitempty"`   // string | { $in: string[] }
	CreatedAt interface{} `json:"createdAt,omitempty"` // string | { $gt: string, $gte: string, $lt: string, $lte: string }
	StartDate interface{} `json:"startDate,omitempty"` // string | { $gt: string, $gte: string, $lt: string, $lte: string }
	EndDate   interface{} `json:"endDate,omitempty"`   // string | { $gt: string, $gte: string, $lt: string, $lte: string }
	Name      interface{} `json:"name,omitempty"`      // string | { $regex: string }
	Limit     *int        `json:"limit,omitempty"`
	Skip      *int        `json:"skip,omitempty"`
	Fields    []string    `json:"fields,omitempty"`
}

type QueryTaskFilterOptions struct {
	Limit   *int    `json:"limit,omitempty"`
	LimitBy *string `json:"limitBy,omitempty"` // 'all' | 'group'
}

type FilteredTasks struct {
	ID    []FilteredTaskID `json:"_id"`
	Tasks []Task           `json:"tasks"`
}

type FilteredTaskID struct {
	Key     string `json:"key"`
	Display string `json:"display"`
	Type    string `json:"type"`
	ID      string `json:"id"`
}

type SMStateChange struct {
	Task         string `json:"task"`
	CurrentState string `json:"currentState"`
	CreatedAt    string `json:"createdAt"`
}
