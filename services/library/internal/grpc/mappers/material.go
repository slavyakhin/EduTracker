package mappers

import (
	"errors"

	"github.com/google/uuid"
	librarycommonv1 "github.com/slavyakhin/EduTracker/protos/gen/go/library/common/v1"
	"github.com/slavyakhin/EduTracker/services/library/internal/domain/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MaterialToProto(
	m models.Material,
) (librarycommonv1.Material, error) {
	const op = "mappers.MaterialToProto"

	protoMaterialType, err := MaterialTypeToProto(m.MaterialType)
	if err != nil {
		// TODO: handle the error
	}

	protoUploadStatus, err := UploadStatusToProto(m.UploadStatus)
	if err != nil {
		// TODO: handle the error
	}

	return librarycommonv1.Material{
		Id:      m.ID.String(),
		OwnerId: m.OwnerID,

		Title:       m.Title,
		Description: m.Description,

		MaterialType: protoMaterialType,
		MimeType:     m.MimeType,
		Extension:    m.Extension,
		SizeBytes:    m.SizeBytes,
		Checksum:     m.Checksum,

		Bucket:       m.Bucket,
		ObjectKey:    m.ObjectKey,
		ThumbnailKey: m.ThumbnailKey,
		UploadStatus: protoUploadStatus,

		// TODO: UserMetadata: ,
		// TODO: SystemMetadata: ,

		CreatedAt:  timestamppb.New(m.CreatedAt),
		UploadedAt: timestamppb.New(m.UpdatedAt),
		DeletedAt:  timestamppb.New(m.DeletedAt),
	}, nil
}

func ProtoToMaterial(
	p *librarycommonv1.Material,
) (models.Material, error) {
	const op = "mappers.ProtoToMaterial"

	id, err := uuid.Parse(p.GetId())
	if err != nil {
		return models.Material{}, errors.New("Invalid material UUID")
	}

	materialType, err := ProtoToMaterialType(p.GetMaterialType())
	if err != nil {
		// TODO: handele the error
	}

	uploadStatus, err := ProtoToUploadStatus(p.GetUploadStatus())
	if err != nil {
		// TODO: handle the error
	}

	return models.Material{
		ID:      id,
		OwnerID: p.GetOwnerId(),

		Title:       p.GetTitle(),
		Description: p.Description,

		MaterialType: materialType,
		MimeType:     p.GetMimeType(),
		Extension:    p.GetExtension(),
		SizeBytes:    p.GetSizeBytes(),
		Checksum:     p.Checksum,

		Bucket:       p.GetBucket(),
		ObjectKey:    p.GetObjectKey(),
		ThumbnailKey: p.ThumbnailKey,
		UploadStatus: uploadStatus,

		// TODO: UserMetadata: ,
		// TODO: SystemMetadata: ,

		UpdatedAt: p.GetCreatedAt().AsTime(),
		DeletedAt: p.DeletedAt.AsTime(),
	}, nil
}
