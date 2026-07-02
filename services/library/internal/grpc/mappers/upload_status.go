package mappers

import (
	"errors"
	"fmt"

	librarycommonv1 "github.com/slavyakhin/EduTracker/protos/gen/go/library/common/v1"
	"github.com/slavyakhin/EduTracker/services/library/internal/domain"
)

var (
	ErrUnsupportedUploadStatus      = errors.New("unsupported upload status")
	ErrUnsupportedProtoUploadStatus = errors.New("unsupported proto upload status")
)

func UploadStatusToProto(
	us domain.UploadStatus,
) (librarycommonv1.UploadStatus, error) {
	const op = "mappers.UploadStatusToProto"

	switch us {
	case domain.StatusUnknown:
		return librarycommonv1.UploadStatus_UPLOAD_STATUS_UNSPECIFIED, nil
	case domain.StatusCreated:
		return librarycommonv1.UploadStatus_UPLOAD_STATUS_CREATED, nil
	case domain.StatusUploading:
		return librarycommonv1.UploadStatus_UPLOAD_STATUS_UPLOADING, nil
	case domain.StatusUploaded:
		return librarycommonv1.UploadStatus_UPLOAD_STATUS_UPLOADED, nil
	case domain.StatusProcessing:
		return librarycommonv1.UploadStatus_UPLOAD_STATUS_PROCESSING, nil
	case domain.StatusReady:
		return librarycommonv1.UploadStatus_UPLOAD_STATUS_READY, nil
	case domain.StatusFailed:
		return librarycommonv1.UploadStatus_UPLOAD_STATUS_FAILED, nil
	default:
		return librarycommonv1.UploadStatus_UPLOAD_STATUS_UNSPECIFIED,
			fmt.Errorf("%s: %w: %v", op, ErrUnsupportedUploadStatus, us)
	}
}

func ProtoToUploadStatus(
	pr librarycommonv1.UploadStatus,
) (domain.UploadStatus, error) {
	const op = "mappers.ProtoToUploadStatus"

	switch pr {
	case librarycommonv1.UploadStatus_UPLOAD_STATUS_UNSPECIFIED:
		return domain.StatusUnknown, nil
	case librarycommonv1.UploadStatus_UPLOAD_STATUS_CREATED:
		return domain.StatusCreated, nil
	case librarycommonv1.UploadStatus_UPLOAD_STATUS_UPLOADING:
		return domain.StatusUploading, nil
	case librarycommonv1.UploadStatus_UPLOAD_STATUS_UPLOADED:
		return domain.StatusUploaded, nil
	case librarycommonv1.UploadStatus_UPLOAD_STATUS_PROCESSING:
		return domain.StatusProcessing, nil
	case librarycommonv1.UploadStatus_UPLOAD_STATUS_READY:
		return domain.StatusReady, nil
	case librarycommonv1.UploadStatus_UPLOAD_STATUS_FAILED:
		return domain.StatusFailed, nil
	default:
		return domain.StatusUnknown,
			fmt.Errorf("%s: %w: %v", op, ErrUnsupportedProtoUploadStatus, pr)
	}
}
