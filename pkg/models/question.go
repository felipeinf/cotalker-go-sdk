package models

type QuestionContentType string

const (
	QuestionContentTypeText         QuestionContentType = "application/vnd.cotalker.survey+text"
	QuestionContentTypeListQuestion QuestionContentType = "application/vnd.cotalker.survey+listquestion"
	QuestionContentTypeTextInput    QuestionContentType = "application/vnd.cotalker.survey+textinput"
	QuestionContentTypeTextNumber   QuestionContentType = "application/vnd.cotalker.survey+textnumber"
	QuestionContentTypeDateTime     QuestionContentType = "application/vnd.cotalker.survey+datetime"
	QuestionContentTypeProperty     QuestionContentType = "application/vnd.cotalker.survey+property"
	QuestionContentTypePerson       QuestionContentType = "application/vnd.cotalker.survey+person"
	QuestionContentTypeCalendar     QuestionContentType = "application/vnd.cotalker.survey+calendar"
	QuestionContentTypeGPS          QuestionContentType = "application/vnd.cotalker.survey+gps"
	QuestionContentTypeImage        QuestionContentType = "application/vnd.cotalker.survey+image"
	QuestionContentTypeSignature    QuestionContentType = "application/vnd.cotalker.survey+signature"
	QuestionContentTypeFile         QuestionContentType = "application/vnd.cotalker.survey+file"
	QuestionContentTypeAPI          QuestionContentType = "application/vnd.cotalker.survey+api"
	QuestionContentTypeSurvey       QuestionContentType = "application/vnd.cotalker.survey+survey"
)

type QuestionExec struct {
	Context *string `json:"context,omitempty"`
	Src     *string `json:"src,omitempty"`
}

type QuestionExecFilter struct {
	Context *string `json:"context,omitempty"`
	Filter  *string `json:"filter,omitempty"`
	Src     *string `json:"src,omitempty"`
}

type QuestionCommand struct {
	Commands         []string               `json:"commands"`
	ResetIdentifiers []string               `json:"resetIdentifiers,omitempty"`
	IsCommanded      string                 `json:"isCommanded"`
	Values           []QuestionCommandValue `json:"values"`
}

type QuestionCommandValue struct {
	Op     string `json:"op"` // '='|'regex'
	Target string `json:"target"`
}

type QuestionDetail struct {
	ID                 string              `json:"_id"`
	ContentType        QuestionContentType `json:"contentType"`
	Display            []string            `json:"display"`
	Code               []string            `json:"code"`
	Identifier         string              `json:"identifier"`
	Symbolizes         *string             `json:"symbolizes,omitempty"`
	Group              string              `json:"group"`
	Company            string              `json:"company"`
	Min                int                 `json:"min"`
	Max                int                 `json:"max"`
	IsActive           bool                `json:"isActive"`
	ModifiedAt         string              `json:"modifiedAt"`
	HideEmpty          bool                `json:"hideEmpty"`
	IsSystemModel      bool                `json:"isSystemModel"`
	IsReadOnly         bool                `json:"isReadOnly"`
	Required           bool                `json:"required"`
	TextAlign          *string             `json:"textAlign,omitempty"` // 'left'|'right'|'center'
	Responses          []string            `json:"responses,omitempty"`
	Command            QuestionCommand     `json:"command"`
	Exec               *QuestionExecs      `json:"exec,omitempty"`
	Subtype            *string             `json:"subtype,omitempty"`
	SkipCodeValidation *bool               `json:"skipCodeValidation,omitempty"`
}

type QuestionExecs struct {
	Preload   *QuestionExec       `json:"preload,omitempty"`
	OnDisplay *QuestionExec       `json:"onDisplay,omitempty"`
	Filter    *QuestionExecFilter `json:"filter,omitempty"`
	Validate  *QuestionExec       `json:"validate,omitempty"`
	OnChange  *QuestionExec       `json:"onChange,omitempty"`
	Presave   *QuestionExec       `json:"presave,omitempty"`
	Postsave  *QuestionExec       `json:"postsave,omitempty"`
}
