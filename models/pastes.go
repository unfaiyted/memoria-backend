package models

import "time"

// Paste represents a stored text snippet with metadata
// @Description A text snippet with formatting, expiration, and privacy settings
type Paste struct {
	ID              string    `gorm:"primaryKey" json:"id" example:"p12345abcde" binding:"required"`
	Title           string    `gorm:"not null" json:"title" example:"My Code Snippet" binding:"required"`
	Content         string    `gorm:"type:text;not null" json:"content" example:"console.log('Hello world');" binding:"required"`
	SyntaxHighlight string    `gorm:"default:'text'" json:"syntax_highlight" example:"javascript" binding:"required"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at" example:"2023-01-01T00:00:00Z"`
	ExpiresAt       time.Time `gorm:"index" json:"expires_at,omitempty" example:"2023-01-08T00:00:00Z"`
	Privacy         string    `gorm:"default:'public'" json:"privacy" example:"public" binding:"required,oneof=public private password"` // "public", "private", "password"
	Password        string    `gorm:"type:varchar(100)" json:"-"`                                                                        // Stored as hash, not returned
	UserID          string    `gorm:"index" json:"user_id,omitempty" example:"u98765zyxwv"`
}

// TableName specifies the database table name for the Paste model
func (Paste) TableName() string {
	return "pastes"
}

// PasteResponse represents the response structure for paste endpoints
// @Description Paste response wrapper
type PasteResponse struct {
	Data  *Paste `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

// PasteRequest
// @Description represents a request to get a paste
type PasteRequest struct {
	ID string `json:"id"`
}

// PasteListResponse represents a list of pastes in response
// @Description List of pastes response wrapper
type PasteListResponse struct {
	Data  []Paste `json:"data,omitempty"`
	Count int     `json:"count,omitempty"`
	Error string  `json:"error,omitempty"`
}
