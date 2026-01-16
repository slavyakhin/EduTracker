package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/slavyakhin/EduTracker/services/library/internal/domain"
)

type Material struct {
	ID      uuid.UUID
	OwnerID int64

	Title       string
	Description *string

	MaterialType domain.MaterialType
	MimeType     string
	Extension    string
	SizeBytes    int64

	Bucket       string  // S3 bucket
	ObjectKey    string  // S3 key
	ThumbnailKey *string // S3 key for thumbnail (nil if not exists)
	UploadStatus domain.UploadStatus

	Metadata map[string]any // json metadata

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
