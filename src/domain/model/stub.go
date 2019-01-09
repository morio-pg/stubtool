package model

import (
	"encoding/json"
	"time"
)

// Stub データの構造体
type Stub struct {
	StubID      string    `json:"stubId" firestore:"stub_id" validate:"required"`
	UID         string    `json:"uid" firestore:"uid" validate:"required"`
	Filename    string    `json:"filename" firestore:"filename" validate:"min=3,max=30"`
	Description string    `json:"description" firestore:"description" validate:"max=40"`
	StatusCode  int       `json:"statusCode" firestore:"status_code" validate:"gte=100,lte=599"`
	Wait        int       `json:"wait" firestore:"wait" validate:"gte=0,lte=15000"`
	MimeTypeKey int       `json:"mimeTypeKey" firestore:"mime_type_key" validate:"gte=1,lte=2"`
	Content     string    `json:"content" firestore:"content" validate:"required,max=5000"`
	CreatedAt   time.Time `json:"createdAt" firestore:"created_at" validate:"required"`
	UpdatedAt   time.Time `json:"updatedAt" firestore:"updated_at" validate:"required"`
}

// MimeTypeToString MIMEタイプを文字列に変換
func (s *Stub) MimeType() (mimeType string) {
	switch s.MimeTypeKey {
	case 1:
		return "application/json"
	case 2:
		return "application/xml"
	default:
		return "unknown"
	}
}

// MarshalJSON JSONへの変換時の内容を指定
func (s *Stub) MarshalJSON() ([]byte, error) {
	type Alias Stub
	return json.Marshal(&struct {
		MimeType string `json:"mimeType"`
		*Alias
	}{
		MimeType: s.MimeType(),
		Alias:    (*Alias)(s),
	})
}
