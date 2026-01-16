package domain

type UploadStatus string

const (
	StatusCreated    UploadStatus = "CREATED"
	StatusUploading  UploadStatus = "UPLOADING"
	StatusUploaded   UploadStatus = "UPLOADED"
	StatusProcessing UploadStatus = "PROCESSING"
	StatusReady      UploadStatus = "READY"
	StatusFailed     UploadStatus = "FAILED"
)
