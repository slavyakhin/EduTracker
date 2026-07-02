package domain

import "fmt"

type MaterialType int

const (
	MaterialUnknown MaterialType = iota
	MaterialImage
	MaterialDocument
	MaterialText
	MaterialVideo
)

func (mt MaterialType) String() string {
	switch mt {
	case MaterialUnknown:
		return "UNKNOWN"
	case MaterialImage:
		return "IMAGE"
	case MaterialDocument:
		return "DOCUMENT"
	case MaterialText:
		return "TEXT"
	case MaterialVideo:
		return "VIDEO"
	default:
		return fmt.Sprintf("MaterialType(%d)", int(mt))
	}
}
