package domain

import "fmt"

type UploadStatus int

const (
	StatusUnknown UploadStatus = iota
	StatusCreated
	StatusUploading
	StatusUploaded
	StatusProcessing
	StatusReady
	StatusFailed
)

func (us UploadStatus) String() string {
	switch us {
	case StatusUnknown:
		return "UNKNOWN"
	case StatusCreated:
		return "CREATED"
	case StatusUploading:
		return "UPLOADING"
	case StatusUploaded:
		return "UPLOADED"
	case StatusProcessing:
		return "PROCESSING"
	case StatusReady:
		return "READY"
	case StatusFailed:
		return "FAILED"
	default:
		return fmt.Sprintf("UploadStatus(%d)", int(us))
	}
}
