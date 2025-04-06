package models

type Survey struct {
	Chat []QuestionChat `json:"chat"`
	ID   string         `json:"_id"`
	Code string         `json:"code"`
}

type PopulatedSurvey struct {
	Survey
	Chat []QuestionChat `json:"chat"`
}

type QuestionChat struct {
	ID           string     `json:"_id"`
	ContentArray []Question `json:"contentArray"`
	IsActive     bool       `json:"isActive"`
}

type SurveyChat struct {
	ID            string   `json:"_id,omitempty"`
	IsActive      bool     `json:"isActive"`
	IsSystemModel bool     `json:"isSystemModel"`
	ContentType   string   `json:"contentType"` // "application/vnd.cotalker.survey"
	Sender        string   `json:"sender"`      // "#system"|"#user"
	Survey        string   `json:"survey"`
	ContentArray  []string `json:"contentArray"`
	Order         int      `json:"order"`
}

type Question struct {
	Identifier  string   `json:"identifier"`
	Display     []string `json:"display"`
	ContentType string   `json:"contentType"`
	Code        string   `json:"code"`
}

type SurveyAPI struct {
	ID            string      `json:"_id"`
	Code          interface{} `json:"code"` // string | ObjectId
	Display       string      `json:"display"`
	Subproperties interface{} `json:"subproperties"` // []Object | ObjectId | string
}
