package models

type Channel struct {
	ID          string   `json:"_id"`
	Group       string   `json:"group"`
	NameCode    string   `json:"nameCode"`
	NameDisplay string   `json:"nameDisplay"`
	UserIDs     []string `json:"userIds"`
}

type ChannelPostBody struct {
	Group       string   `json:"group"`
	NameCode    string   `json:"nameCode"`
	NameDisplay string   `json:"nameDisplay"`
	UserIDs     []string `json:"userIds"`
}
