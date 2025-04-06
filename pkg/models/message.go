package models

type SendMsgBody struct {
	Channel     string `json:"channel"`
	Content     string `json:"content"`
	ContentType string `json:"contentType"` // "text/system" | "text/plain"
	IsSaved     int    `json:"isSaved"`
	SentBy      string `json:"sentBy"`
}

type EditMsgBody struct {
	Channel string `json:"channel"`
	Content string `json:"content"`
	IsSaved int    `json:"isSaved"`
}
