package models

type AnswerData struct {
	Code        []string `json:"code"`
	Identifier  string   `json:"identifier"`
	ContentType string   `json:"contentType"`
	Process     []string `json:"process"`
	Responses   []string `json:"responses"`
	Group       string   `json:"group"`
	User        string   `json:"user"`
}

type Answer struct {
	ID                string       `json:"_id"`
	UUID              string       `json:"uuid"`
	Survey            string       `json:"survey"`
	Channel           string       `json:"channel"`
	User              string       `json:"user"`
	Properties        []string     `json:"properties"`
	PropertyTypes     []string     `json:"propertyTypes"`
	IdentifiersNeeded []string     `json:"identifiersNeeded"`
	ExtendsAnswer     []string     `json:"extendsAnswer"`
	RExtendsAnswer    []string     `json:"rExtendsAnswer"`
	Data              []AnswerData `json:"data"`
	CreatedAt         string       `json:"createdAt"`
	ModifiedAt        string       `json:"modifiedAt"`
	Score             AnswerScore  `json:"score"`
}

type AnswerScore struct {
	Main   float64           `json:"main"`
	Scores []AnswerScoreItem `json:"scores"`
}

type AnswerScoreItem struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
