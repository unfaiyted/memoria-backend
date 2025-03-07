package models

import "time"

// Paste represents a stored text snippet with metadata
// @Description A text snippet with formatting, expiration, and privacy settings
type Paste struct {
	ID              uint64    `gorm:"primaryKey" json:"id" example:"123111" binding:"required"`
	Title           string    `gorm:"not null" json:"title" example:"My Code Snippet" binding:"required"`
	Content         string    `gorm:"type:text;not null" json:"content" example:"console.log('Hello world');" binding:"required"`
	SyntaxHighlight string    `gorm:"default:'text'" json:"syntaxHighlight" example:"javascript" binding:"required"`
	EditorType      string    `gorm:"default:'code';column:editor_type" json:"editorType" example:"code" binding:"required,oneof=code text"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"createdAt" example:"2023-01-01T00:00:00Z"`
	ExpiresAt       time.Time `gorm:"index" json:"expiresAt,omitempty" example:"2023-01-08T00:00:00Z"`
	Privacy         string    `gorm:"default:'public'" json:"privacy" example:"public" binding:"required,oneof=public private"` // "public", "private"
	PrivateAccessID string    `gorm:"type:varchar(64);uniqueIndex" json:"privateAccessId,omitempty" example:"abc123xyz456"`
	Password        string    `gorm:"type:varchar(100)" json:"-"` // Stored as hash, not returned
	UserID          string    `gorm:"index" json:"user_id,omitempty" example:"u98765zyxwv"`
}

type CreatePasteRequest struct {
	Title           string    `json:"title" binding:"required"`
	Content         string    `json:"content" binding:"required"`
	SyntaxHighlight string    `json:"syntaxHighlight,omitempty" `
	EditorType      string    `json:"editorType,omitempty" example:"code" binding:"oneof=code text"`
	ExpiresAt       time.Time `json:"expiresAt,omitempty" example:"2023-01-08T00:00:00Z"`
	Privacy         string    `json:"privacy" binding:"required,oneof=public private password"`
	Password        string    `json:"password,omitempty" example:"mySecurePassword123"`
}

type UpdatePasteRequest struct {
	Title           string    `json:"title" binding:"required"`
	ID              uint64    `json:"id" binding:"required"`
	Content         string    `json:"content" binding:"required"`
	SyntaxHighlight string    `json:"syntaxHighlight,omitempty" `
	EditorType      string    `json:"editorType,omitempty"`
	ExpiresAt       time.Time `json:"expiresAt,omitempty" example:"2023-01-08T00:00:00Z"`
	Privacy         string    `json:"privacy" binding:"required,oneof=public private password"`
	Password        string    `json:"password,omitempty" example:"mySecurePassword123"`
}

type PasteListRequest struct {
	Page    int `json:"page" form:"page" binding:"required"`
	PerPage int `json:"perPage" form:"perPage" binding:"required"`
}

// TableName specifies the database table name for the Paste model
func (Paste) TableName() string {
	return "pastes"
}

// // PasteData represents the response  data for paste endpoints
// // @Description Paste data response wrapper
type PasteData struct {
	Paste *Paste `json:"paste,omitempty"`
}

// PasteRequest
// @Description represents a request to get a paste
type PasteRequest struct {
	ID uint64 `json:"id"`
}

// PasteListResponse represents a list of pastes in response
// @Description List of pastes response wrapper
type PasteListData struct {
	Pastes []Paste `json:"pastes,omitempty"`
	Count  int     `json:"count,omitempty"`
}
