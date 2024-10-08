package compile

import (
	"fmt"

	"github.com/kaloseia/morphe-go/pkg/yaml"
)

func ErrUnsupportedMorpheFieldType(unsupportedType yaml.ModelFieldType) error {
	return fmt.Errorf("unsupported morphe field type for go conversion: '%s'", unsupportedType)
}

func ErrMissingMorpheIdentifierField(modelName string, identifierName string, fieldName string) error {
	return fmt.Errorf("morphe model '%s' has no field '%s' referenced in identifiers ('%s')", modelName, identifierName, fieldName)
}
