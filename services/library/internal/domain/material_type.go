package domain

type MaterialType string

const (
	MaterialImage    MaterialType = "image"
	MaterialDocument MaterialType = "document"
	MaterialText     MaterialType = "text"
	MaterialVideo    MaterialType = "video"
)
