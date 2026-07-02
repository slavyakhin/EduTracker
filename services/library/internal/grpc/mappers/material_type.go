package mappers

import (
	"errors"
	"fmt"

	librarycommonv1 "github.com/slavyakhin/EduTracker/protos/gen/go/library/common/v1"
	"github.com/slavyakhin/EduTracker/services/library/internal/domain"
)

var (
	ErrUnsupportedMaterialType      = errors.New("unknown material type")
	ErrUnsupportedProtoMaterialType = errors.New("unknown proto material type")
)

func MaterialTypeToProto(
	mt domain.MaterialType,
) (librarycommonv1.MaterialType, error) {
	const op = "mappers.MaterialTypeToProto"

	switch mt {
	case domain.MaterialUnknown:
		return librarycommonv1.MaterialType_MATERIAL_TYPE_UNSPECIFIED, nil
	case domain.MaterialImage:
		return librarycommonv1.MaterialType_MATERIAL_TYPE_IMAGE, nil
	case domain.MaterialDocument:
		return librarycommonv1.MaterialType_MATERIAL_TYPE_DOCUMENT, nil
	case domain.MaterialText:
		return librarycommonv1.MaterialType_MATERIAL_TYPE_TEXT, nil
	case domain.MaterialVideo:
		return librarycommonv1.MaterialType_MATERIAL_TYPE_VIDEO, nil
	default:
		return librarycommonv1.MaterialType_MATERIAL_TYPE_UNSPECIFIED,
			fmt.Errorf("%s: %w: %v", op, ErrUnsupportedMaterialType, mt)
	}
}

func ProtoToMaterialType(
	pr librarycommonv1.MaterialType,
) (domain.MaterialType, error) {
	const op = "mappers.ProtoToMaterialType"

	switch pr {
	case librarycommonv1.MaterialType_MATERIAL_TYPE_UNSPECIFIED:
		return domain.MaterialUnknown, nil
	case librarycommonv1.MaterialType_MATERIAL_TYPE_IMAGE:
		return domain.MaterialImage, nil
	case librarycommonv1.MaterialType_MATERIAL_TYPE_DOCUMENT:
		return domain.MaterialDocument, nil
	case librarycommonv1.MaterialType_MATERIAL_TYPE_TEXT:
		return domain.MaterialText, nil
	case librarycommonv1.MaterialType_MATERIAL_TYPE_VIDEO:
		return domain.MaterialVideo, nil
	default:
		return domain.MaterialUnknown,
			fmt.Errorf("%s: %w: %v", op, ErrUnsupportedProtoMaterialType, pr)
	}
}
