package models

type FileContentType string

const (
	FileContentTypeImage    FileContentType = "image"
	FileContentTypeAudio    FileContentType = "audio"
	FileContentTypeVideo    FileContentType = "video"
	FileContentTypeDocument FileContentType = "document"
)

type FileStatus string

const (
	FileStatusPending    FileStatus = "pending"
	FileStatusProcessing FileStatus = "processing"
	FileStatusUploaded   FileStatus = "uploaded"
	FileStatusDeleted    FileStatus = "deleted"
	FileStatusError      FileStatus = "error"
)

type File struct {
	ID          string          `json:"_id"`
	ContentType FileContentType `json:"contentType"`
	Company     *string         `json:"company,omitempty"`
	Status      FileStatus      `json:"status"`
	User        *string         `json:"user,omitempty"`
	CreatedAt   string          `json:"createdAt"`
	ModifiedAt  string          `json:"modifiedAt"`
	Public      bool            `json:"public"`
	URL         *string         `json:"url,omitempty"`
	Extension   *string         `json:"extension,omitempty"`
	MimeType    *string         `json:"mimeType,omitempty"`
	FileName    *string         `json:"fileName,omitempty"`
	Size        *int64          `json:"size,omitempty"`
}

type FileUploaded struct {
	File
	Status    FileStatus `json:"status"` // Siempre será "uploaded"
	URL       string     `json:"url"`
	Extension string     `json:"extension"`
	MimeType  string     `json:"mimeType"`
	FileName  string     `json:"fileName"`
	Size      int64      `json:"size"`
}
