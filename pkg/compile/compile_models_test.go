package compile_test

import (
	"fmt"
	"testing"

	"github.com/kaloseia/go/pkg/godef"
	"github.com/kaloseia/morphe-go/pkg/yaml"
	"github.com/kaloseia/plugin-morphe-go-struct/pkg/compile"
	"github.com/kaloseia/plugin-morphe-go-struct/pkg/compile/cfg"
	"github.com/kaloseia/plugin-morphe-go-struct/pkg/compile/hook"
	"github.com/stretchr/testify/suite"
)

type CompileModelsTestSuite struct {
	suite.Suite
}

func TestCompileModelsTestSuite(t *testing.T) {
	suite.Run(t, new(CompileModelsTestSuite))
}

func (suite *CompileModelsTestSuite) SetupTest() {
}

func (suite *CompileModelsTestSuite) TearDownTest() {
}

func (suite *CompileModelsTestSuite) TestMorpheModelToGoStructs() {
	modelHooks := hook.CompileMorpheModel{}

	modelsConfig := cfg.MorpheModelsConfig{
		Package: godef.Package{
			Path: "github.com/kaloseia/project/domain/models",
			Name: "models",
		},
		ReceiverName: "m",
	}

	model0 := yaml.Model{
		Name: "Basic",
		Fields: map[string]yaml.ModelField{
			"AutoIncrement": {
				Type: yaml.ModelFieldTypeAutoIncrement,
			},
			"Boolean": {
				Type: yaml.ModelFieldTypeBoolean,
			},
			"Date": {
				Type: yaml.ModelFieldTypeDate,
			},
			"Float": {
				Type: yaml.ModelFieldTypeFloat,
			},
			"Integer": {
				Type: yaml.ModelFieldTypeInteger,
			},
			"Protected": {
				Type: yaml.ModelFieldTypeProtected,
			},
			"Sealed": {
				Type: yaml.ModelFieldTypeSealed,
			},
			"String": {
				Type: yaml.ModelFieldTypeString,
			},
			"Time": {
				Type: yaml.ModelFieldTypeTime,
			},
			"UUID": {
				Type: yaml.ModelFieldTypeUUID,
				Attributes: []string{
					"immutable",
				},
			},
		},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {
				Fields: []string{
					"UUID",
				},
			},
		},
		Related: map[string]yaml.ModelRelation{},
	}

	allGoStructs, allStructsErr := compile.MorpheModelToGoStructs(modelHooks, modelsConfig, model0)

	suite.Nil(allStructsErr)
	suite.Len(allGoStructs, 2)

	goStruct0 := allGoStructs[0]

	suite.Equal(goStruct0.Package.Path, modelsConfig.Package.Path)
	suite.Equal(goStruct0.Package.Name, modelsConfig.Package.Name)

	structImports0 := goStruct0.Imports
	suite.Len(structImports0, 1)

	structImports00 := structImports0[0]
	suite.Equal(structImports00, "time")

	suite.Equal(goStruct0.Name, "Basic")

	structFields0 := goStruct0.Fields
	suite.Len(structFields0, 10)

	structFields00 := structFields0[0]
	suite.Equal(structFields00.Name, "AutoIncrement")
	suite.Equal(structFields00.Type, godef.GoTypeUint)
	suite.Len(structFields00.Tags, 0)

	structFields01 := structFields0[1]
	suite.Equal(structFields01.Name, "Boolean")
	suite.Equal(structFields01.Type, godef.GoTypeBool)
	suite.Len(structFields01.Tags, 0)

	structFields02 := structFields0[2]
	suite.Equal(structFields02.Name, "Date")
	suite.Equal(structFields02.Type, godef.GoTypeTime)
	suite.Len(structFields02.Tags, 0)

	structFields03 := structFields0[3]
	suite.Equal(structFields03.Name, "Float")
	suite.Equal(structFields03.Type, godef.GoTypeFloat)
	suite.Len(structFields03.Tags, 0)

	structFields04 := structFields0[4]
	suite.Equal(structFields04.Name, "Integer")
	suite.Equal(structFields04.Type, godef.GoTypeInt)
	suite.Len(structFields04.Tags, 0)

	structFields05 := structFields0[5]
	suite.Equal(structFields05.Name, "Protected")
	suite.Equal(structFields05.Type, godef.GoTypeString)
	suite.Len(structFields05.Tags, 0)

	structFields06 := structFields0[6]
	suite.Equal(structFields06.Name, "Sealed")
	suite.Equal(structFields06.Type, godef.GoTypeString)
	suite.Len(structFields06.Tags, 0)

	structFields07 := structFields0[7]
	suite.Equal(structFields07.Name, "String")
	suite.Equal(structFields07.Type, godef.GoTypeString)
	suite.Len(structFields07.Tags, 0)

	structFields08 := structFields0[8]
	suite.Equal(structFields08.Name, "Time")
	suite.Equal(structFields08.Type, godef.GoTypeTime)
	suite.Len(structFields08.Tags, 0)

	structFields09 := structFields0[9]
	suite.Equal(structFields09.Name, "UUID")
	suite.Equal(structFields09.Type, godef.GoTypeString)
	suite.Len(structFields09.Tags, 1)
	suite.Equal(structFields09.Tags[0], "immutable")

	structMethods0 := goStruct0.Methods
	suite.Len(structMethods0, 1)

	basicType := godef.GoTypeStruct{PackagePath: modelsConfig.Package.Path, Name: "Basic"}
	basicIDPrimaryType := godef.GoTypeStruct{PackagePath: modelsConfig.Package.Path, Name: "BasicIDPrimary"}
	structMethods00 := structMethods0[0]
	suite.Equal(structMethods00.ReceiverName, modelsConfig.ReceiverName)
	suite.Equal(structMethods00.ReceiverType, basicType)
	suite.Equal(structMethods00.Name, "GetIDPrimary")
	suite.Nil(structMethods00.Parameters)
	suite.Equal(structMethods00.ReturnTypes, []godef.GoType{
		basicIDPrimaryType,
	})
	suite.Equal(structMethods00.BodyLines, []string{
		"	return BasicIDPrimary{",
		"		UUID: m.UUID,",
		"	}",
	})

	goStruct1 := allGoStructs[1]

	suite.Equal(goStruct1.Package.Path, modelsConfig.Package.Path)
	suite.Equal(goStruct1.Package.Name, modelsConfig.Package.Name)

	structImports1 := goStruct1.Imports
	suite.Len(structImports1, 0)

	suite.Equal(goStruct1.Name, "BasicIDPrimary")

	structFields1 := goStruct1.Fields
	suite.Len(structFields1, 1)

	structFields10 := structFields1[0]
	suite.Equal(structFields10.Name, "UUID")
	suite.Equal(structFields10.Type, godef.GoTypeString)
	suite.Len(structFields10.Tags, 0)
}

func (suite *CompileModelsTestSuite) TestMorpheModelToGoStructs_NoPackagePath() {
	modelHooks := hook.CompileMorpheModel{}

	modelsConfig := cfg.MorpheModelsConfig{
		Package: godef.Package{
			Path: "",
			Name: "models",
		},
		ReceiverName: "m",
	}

	model0 := yaml.Model{
		Name: "Basic",
		Fields: map[string]yaml.ModelField{
			"AutoIncrement": {
				Type: yaml.ModelFieldTypeAutoIncrement,
			},
		},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {
				Fields: []string{
					"AutoIncrement",
				},
			},
		},
		Related: map[string]yaml.ModelRelation{},
	}

	allGoStructs, allStructsErr := compile.MorpheModelToGoStructs(modelHooks, modelsConfig, model0)

	suite.NotNil(allStructsErr)
	suite.ErrorContains(allStructsErr, "models package path cannot be empty")

	suite.Nil(allGoStructs)
}

func (suite *CompileModelsTestSuite) TestMorpheModelToGoStructs_NoPackageName() {
	modelHooks := hook.CompileMorpheModel{}

	modelsConfig := cfg.MorpheModelsConfig{
		Package: godef.Package{
			Path: "github.com/kaloseia/project/domain/models",
			Name: "",
		},
		ReceiverName: "m",
	}

	model0 := yaml.Model{
		Name: "Basic",
		Fields: map[string]yaml.ModelField{
			"AutoIncrement": {
				Type: yaml.ModelFieldTypeAutoIncrement,
			},
		},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {
				Fields: []string{
					"AutoIncrement",
				},
			},
		},
		Related: map[string]yaml.ModelRelation{},
	}

	allGoStructs, allStructsErr := compile.MorpheModelToGoStructs(modelHooks, modelsConfig, model0)

	suite.NotNil(allStructsErr)
	suite.ErrorContains(allStructsErr, "models package name cannot be empty")

	suite.Nil(allGoStructs)
}

func (suite *CompileModelsTestSuite) TestMorpheModelToGoStructs_NoReceiverName() {
	modelHooks := hook.CompileMorpheModel{}

	modelsConfig := cfg.MorpheModelsConfig{
		Package: godef.Package{
			Path: "github.com/kaloseia/project/domain/models",
			Name: "models",
		},
		ReceiverName: "",
	}

	model0 := yaml.Model{
		Name: "Basic",
		Fields: map[string]yaml.ModelField{
			"AutoIncrement": {
				Type: yaml.ModelFieldTypeAutoIncrement,
			},
		},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {
				Fields: []string{
					"AutoIncrement",
				},
			},
		},
		Related: map[string]yaml.ModelRelation{},
	}

	allGoStructs, allStructsErr := compile.MorpheModelToGoStructs(modelHooks, modelsConfig, model0)

	suite.NotNil(allStructsErr)
	suite.ErrorContains(allStructsErr, "models method receiver name cannot be empty")

	suite.Nil(allGoStructs)
}

func (suite *CompileModelsTestSuite) TestMorpheModelToGoStructs_NoModelName() {
	modelHooks := hook.CompileMorpheModel{}

	modelsConfig := cfg.MorpheModelsConfig{
		Package: godef.Package{
			Path: "github.com/kaloseia/project/domain/models",
			Name: "models",
		},
		ReceiverName: "m",
	}

	model0 := yaml.Model{
		Name: "",
		Fields: map[string]yaml.ModelField{
			"AutoIncrement": {
				Type: yaml.ModelFieldTypeAutoIncrement,
			},
		},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {
				Fields: []string{
					"AutoIncrement",
				},
			},
		},
		Related: map[string]yaml.ModelRelation{},
	}

	allGoStructs, allStructsErr := compile.MorpheModelToGoStructs(modelHooks, modelsConfig, model0)

	suite.NotNil(allStructsErr)
	suite.ErrorContains(allStructsErr, "morphe model has no name")

	suite.Nil(allGoStructs)
}

func (suite *CompileModelsTestSuite) TestMorpheModelToGoStructs_NoFields() {
	modelHooks := hook.CompileMorpheModel{}

	modelsConfig := cfg.MorpheModelsConfig{
		Package: godef.Package{
			Path: "github.com/kaloseia/project/domain/models",
			Name: "models",
		},
		ReceiverName: "m",
	}

	model0 := yaml.Model{
		Name:   "Basic",
		Fields: map[string]yaml.ModelField{},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {
				Fields: []string{
					"AutoIncrement",
				},
			},
		},
		Related: map[string]yaml.ModelRelation{},
	}

	allGoStructs, allStructsErr := compile.MorpheModelToGoStructs(modelHooks, modelsConfig, model0)

	suite.NotNil(allStructsErr)
	suite.ErrorContains(allStructsErr, "morphe model has no fields")

	suite.Nil(allGoStructs)
}

func (suite *CompileModelsTestSuite) TestMorpheModelToGoStructs_NoIdentifiers() {
	modelHooks := hook.CompileMorpheModel{}

	modelsConfig := cfg.MorpheModelsConfig{
		Package: godef.Package{
			Path: "github.com/kaloseia/project/domain/models",
			Name: "models",
		},
		ReceiverName: "m",
	}

	model0 := yaml.Model{
		Name: "Basic",
		Fields: map[string]yaml.ModelField{
			"AutoIncrement": {
				Type: yaml.ModelFieldTypeAutoIncrement,
			},
		},
		Identifiers: map[string]yaml.ModelIdentifier{},
		Related:     map[string]yaml.ModelRelation{},
	}

	allGoStructs, allStructsErr := compile.MorpheModelToGoStructs(modelHooks, modelsConfig, model0)

	suite.NotNil(allStructsErr)
	suite.ErrorContains(allStructsErr, "morphe model has no identifiers")

	suite.Nil(allGoStructs)
}

func (suite *CompileModelsTestSuite) TestMorpheModelToGoStructs_StartHook_Successful() {
	var featureFlag = "otherName"
	modelHooks := hook.CompileMorpheModel{
		OnCompileMorpheModelStart: func(config cfg.MorpheModelsConfig, model yaml.Model) (cfg.MorpheModelsConfig, yaml.Model, error) {
			if featureFlag != "otherName" {
				return config, model, nil
			}
			config.ReceiverName = "CHANGED"
			model.Name = model.Name + "CHANGED"
			delete(model.Fields, "Float")
			return config, model, nil
		},
	}

	modelsConfig := cfg.MorpheModelsConfig{
		Package: godef.Package{
			Path: "github.com/kaloseia/project/domain/models",
			Name: "models",
		},
		ReceiverName: "m",
	}

	model0 := yaml.Model{
		Name: "Basic",
		Fields: map[string]yaml.ModelField{
			"AutoIncrement": {
				Type: yaml.ModelFieldTypeAutoIncrement,
			},
			"Boolean": {
				Type: yaml.ModelFieldTypeBoolean,
			},
			"Date": {
				Type: yaml.ModelFieldTypeDate,
			},
			"Float": {
				Type: yaml.ModelFieldTypeFloat,
			},
			"Integer": {
				Type: yaml.ModelFieldTypeInteger,
			},
			"Protected": {
				Type: yaml.ModelFieldTypeProtected,
			},
			"Sealed": {
				Type: yaml.ModelFieldTypeSealed,
			},
			"String": {
				Type: yaml.ModelFieldTypeString,
			},
			"Time": {
				Type: yaml.ModelFieldTypeTime,
			},
			"UUID": {
				Type: yaml.ModelFieldTypeUUID,
				Attributes: []string{
					"immutable",
				},
			},
		},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {
				Fields: []string{
					"UUID",
				},
			},
		},
		Related: map[string]yaml.ModelRelation{},
	}

	allGoStructs, allStructsErr := compile.MorpheModelToGoStructs(modelHooks, modelsConfig, model0)

	suite.Nil(allStructsErr)
	suite.Len(allGoStructs, 2)

	goStruct0 := allGoStructs[0]

	suite.Equal(goStruct0.Package.Path, modelsConfig.Package.Path)
	suite.Equal(goStruct0.Package.Name, modelsConfig.Package.Name)

	structImports0 := goStruct0.Imports
	suite.Len(structImports0, 1)

	structImports00 := structImports0[0]
	suite.Equal(structImports00, "time")

	suite.Equal(goStruct0.Name, "BasicCHANGED")

	structFields0 := goStruct0.Fields
	suite.Len(structFields0, 9)

	structFields00 := structFields0[0]
	suite.Equal(structFields00.Name, "AutoIncrement")
	suite.Equal(structFields00.Type, godef.GoTypeUint)
	suite.Len(structFields00.Tags, 0)

	structFields01 := structFields0[1]
	suite.Equal(structFields01.Name, "Boolean")
	suite.Equal(structFields01.Type, godef.GoTypeBool)
	suite.Len(structFields01.Tags, 0)

	structFields02 := structFields0[2]
	suite.Equal(structFields02.Name, "Date")
	suite.Equal(structFields02.Type, godef.GoTypeTime)
	suite.Len(structFields02.Tags, 0)

	structFields04 := structFields0[3]
	suite.Equal(structFields04.Name, "Integer")
	suite.Equal(structFields04.Type, godef.GoTypeInt)
	suite.Len(structFields04.Tags, 0)

	structFields05 := structFields0[4]
	suite.Equal(structFields05.Name, "Protected")
	suite.Equal(structFields05.Type, godef.GoTypeString)
	suite.Len(structFields05.Tags, 0)

	structFields06 := structFields0[5]
	suite.Equal(structFields06.Name, "Sealed")
	suite.Equal(structFields06.Type, godef.GoTypeString)
	suite.Len(structFields06.Tags, 0)

	structFields07 := structFields0[6]
	suite.Equal(structFields07.Name, "String")
	suite.Equal(structFields07.Type, godef.GoTypeString)
	suite.Len(structFields07.Tags, 0)

	structFields08 := structFields0[7]
	suite.Equal(structFields08.Name, "Time")
	suite.Equal(structFields08.Type, godef.GoTypeTime)
	suite.Len(structFields08.Tags, 0)

	structFields09 := structFields0[8]
	suite.Equal(structFields09.Name, "UUID")
	suite.Equal(structFields09.Type, godef.GoTypeString)
	suite.Len(structFields09.Tags, 1)
	suite.Equal(structFields09.Tags[0], "immutable")

	structMethods0 := goStruct0.Methods
	suite.Len(structMethods0, 1)

	basicType := godef.GoTypeStruct{PackagePath: modelsConfig.Package.Path, Name: "BasicCHANGED"}
	basicIDPrimaryType := godef.GoTypeStruct{PackagePath: modelsConfig.Package.Path, Name: "BasicCHANGEDIDPrimary"}
	structMethods00 := structMethods0[0]
	suite.Equal(structMethods00.ReceiverName, "CHANGED")
	suite.Equal(structMethods00.ReceiverType, basicType)
	suite.Equal(structMethods00.Name, "GetIDPrimary")
	suite.Nil(structMethods00.Parameters)
	suite.Equal(structMethods00.ReturnTypes, []godef.GoType{
		basicIDPrimaryType,
	})
	suite.Equal(structMethods00.BodyLines, []string{
		"	return BasicCHANGEDIDPrimary{",
		"		UUID: CHANGED.UUID,",
		"	}",
	})

	goStruct1 := allGoStructs[1]

	suite.Equal(goStruct1.Package.Path, modelsConfig.Package.Path)
	suite.Equal(goStruct1.Package.Name, modelsConfig.Package.Name)

	structImports1 := goStruct1.Imports
	suite.Len(structImports1, 0)

	suite.Equal(goStruct1.Name, "BasicCHANGEDIDPrimary")

	structFields1 := goStruct1.Fields
	suite.Len(structFields1, 1)

	structFields10 := structFields1[0]
	suite.Equal(structFields10.Name, "UUID")
	suite.Equal(structFields10.Type, godef.GoTypeString)
	suite.Len(structFields10.Tags, 0)
}

func (suite *CompileModelsTestSuite) TestMorpheModelToGoStructs_StartHook_Failure() {
	var featureFlag = "otherName"
	modelHooks := hook.CompileMorpheModel{
		OnCompileMorpheModelStart: func(config cfg.MorpheModelsConfig, model yaml.Model) (cfg.MorpheModelsConfig, yaml.Model, error) {
			if featureFlag != "otherName" {
				return config, model, nil
			}
			return config, model, fmt.Errorf("compile model start hook error")
		},
	}

	modelsConfig := cfg.MorpheModelsConfig{
		Package: godef.Package{
			Path: "github.com/kaloseia/project/domain/models",
			Name: "models",
		},
		ReceiverName: "m",
	}

	model0 := yaml.Model{
		Name: "Basic",
		Fields: map[string]yaml.ModelField{
			"AutoIncrement": {
				Type: yaml.ModelFieldTypeAutoIncrement,
			},
			"Boolean": {
				Type: yaml.ModelFieldTypeBoolean,
			},
			"Date": {
				Type: yaml.ModelFieldTypeDate,
			},
			"Float": {
				Type: yaml.ModelFieldTypeFloat,
			},
			"Integer": {
				Type: yaml.ModelFieldTypeInteger,
			},
			"Protected": {
				Type: yaml.ModelFieldTypeProtected,
			},
			"Sealed": {
				Type: yaml.ModelFieldTypeSealed,
			},
			"String": {
				Type: yaml.ModelFieldTypeString,
			},
			"Time": {
				Type: yaml.ModelFieldTypeTime,
			},
			"UUID": {
				Type: yaml.ModelFieldTypeUUID,
				Attributes: []string{
					"immutable",
				},
			},
		},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {
				Fields: []string{
					"UUID",
				},
			},
		},
		Related: map[string]yaml.ModelRelation{},
	}

	allGoStructs, allStructsErr := compile.MorpheModelToGoStructs(modelHooks, modelsConfig, model0)

	suite.NotNil(allStructsErr)
	suite.ErrorContains(allStructsErr, "compile model start hook error")
	suite.Nil(allGoStructs)
}

func (suite *CompileModelsTestSuite) TestMorpheModelToGoStructs_SuccessHook_Successful() {
	var featureFlag = "otherName"
	modelHooks := hook.CompileMorpheModel{
		OnCompileMorpheModelSuccess: func(allModelStructs []*godef.Struct) ([]*godef.Struct, error) {
			if featureFlag != "otherName" {
				return allModelStructs, nil
			}
			for _, modelStructPtr := range allModelStructs {
				modelStructPtr.Name = modelStructPtr.Name + "CHANGED"
				newFields := []godef.StructField{}
				for _, modelStructField := range modelStructPtr.Fields {
					if modelStructField.Name == "Float" {
						continue
					}
					newFields = append(newFields, modelStructField)
				}
				modelStructPtr.Fields = newFields
			}
			return allModelStructs, nil
		},
	}

	modelsConfig := cfg.MorpheModelsConfig{
		Package: godef.Package{
			Path: "github.com/kaloseia/project/domain/models",
			Name: "models",
		},
		ReceiverName: "m",
	}

	model0 := yaml.Model{
		Name: "Basic",
		Fields: map[string]yaml.ModelField{
			"AutoIncrement": {
				Type: yaml.ModelFieldTypeAutoIncrement,
			},
			"Boolean": {
				Type: yaml.ModelFieldTypeBoolean,
			},
			"Date": {
				Type: yaml.ModelFieldTypeDate,
			},
			"Float": {
				Type: yaml.ModelFieldTypeFloat,
			},
			"Integer": {
				Type: yaml.ModelFieldTypeInteger,
			},
			"Protected": {
				Type: yaml.ModelFieldTypeProtected,
			},
			"Sealed": {
				Type: yaml.ModelFieldTypeSealed,
			},
			"String": {
				Type: yaml.ModelFieldTypeString,
			},
			"Time": {
				Type: yaml.ModelFieldTypeTime,
			},
			"UUID": {
				Type: yaml.ModelFieldTypeUUID,
				Attributes: []string{
					"immutable",
				},
			},
		},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {
				Fields: []string{
					"UUID",
				},
			},
		},
		Related: map[string]yaml.ModelRelation{},
	}

	allGoStructs, allStructsErr := compile.MorpheModelToGoStructs(modelHooks, modelsConfig, model0)

	suite.Nil(allStructsErr)
	suite.Len(allGoStructs, 2)

	goStruct0 := allGoStructs[0]

	suite.Equal(goStruct0.Package.Path, modelsConfig.Package.Path)
	suite.Equal(goStruct0.Package.Name, modelsConfig.Package.Name)

	structImports0 := goStruct0.Imports
	suite.Len(structImports0, 1)

	structImports00 := structImports0[0]
	suite.Equal(structImports00, "time")

	suite.Equal(goStruct0.Name, "BasicCHANGED")

	structFields0 := goStruct0.Fields
	suite.Len(structFields0, 9)

	structFields00 := structFields0[0]
	suite.Equal(structFields00.Name, "AutoIncrement")
	suite.Equal(structFields00.Type, godef.GoTypeUint)
	suite.Len(structFields00.Tags, 0)

	structFields01 := structFields0[1]
	suite.Equal(structFields01.Name, "Boolean")
	suite.Equal(structFields01.Type, godef.GoTypeBool)
	suite.Len(structFields01.Tags, 0)

	structFields02 := structFields0[2]
	suite.Equal(structFields02.Name, "Date")
	suite.Equal(structFields02.Type, godef.GoTypeTime)
	suite.Len(structFields02.Tags, 0)

	structFields04 := structFields0[3]
	suite.Equal(structFields04.Name, "Integer")
	suite.Equal(structFields04.Type, godef.GoTypeInt)
	suite.Len(structFields04.Tags, 0)

	structFields05 := structFields0[4]
	suite.Equal(structFields05.Name, "Protected")
	suite.Equal(structFields05.Type, godef.GoTypeString)
	suite.Len(structFields05.Tags, 0)

	structFields06 := structFields0[5]
	suite.Equal(structFields06.Name, "Sealed")
	suite.Equal(structFields06.Type, godef.GoTypeString)
	suite.Len(structFields06.Tags, 0)

	structFields07 := structFields0[6]
	suite.Equal(structFields07.Name, "String")
	suite.Equal(structFields07.Type, godef.GoTypeString)
	suite.Len(structFields07.Tags, 0)

	structFields08 := structFields0[7]
	suite.Equal(structFields08.Name, "Time")
	suite.Equal(structFields08.Type, godef.GoTypeTime)
	suite.Len(structFields08.Tags, 0)

	structFields09 := structFields0[8]
	suite.Equal(structFields09.Name, "UUID")
	suite.Equal(structFields09.Type, godef.GoTypeString)
	suite.Len(structFields09.Tags, 1)
	suite.Equal(structFields09.Tags[0], "immutable")

	structMethods0 := goStruct0.Methods
	suite.Len(structMethods0, 1)

	basicType := godef.GoTypeStruct{PackagePath: modelsConfig.Package.Path, Name: "Basic"}
	basicIDPrimaryType := godef.GoTypeStruct{PackagePath: modelsConfig.Package.Path, Name: "BasicIDPrimary"}
	structMethods00 := structMethods0[0]
	suite.Equal(structMethods00.ReceiverName, "GetIDPrimary")
	suite.Equal(structMethods00.ReceiverType, basicType)
	suite.Equal(structMethods00.Name, "GetIDPrimary")
	suite.Nil(structMethods00.Parameters)
	suite.Equal(structMethods00.ReturnTypes, []godef.GoType{
		basicIDPrimaryType,
	})
	suite.Equal(structMethods00.BodyLines, []string{
		"	return BasicIDPrimary{",
		"		UUID: m.UUID,",
		"	}",
	})

	goStruct1 := allGoStructs[1]

	suite.Equal(goStruct1.Package.Path, modelsConfig.Package.Path)
	suite.Equal(goStruct1.Package.Name, modelsConfig.Package.Name)

	structImports1 := goStruct1.Imports
	suite.Len(structImports1, 0)

	suite.Equal(goStruct1.Name, "BasicIDPrimaryCHANGED")

	structFields1 := goStruct1.Fields
	suite.Len(structFields1, 1)

	structFields10 := structFields1[0]
	suite.Equal(structFields10.Name, "UUID")
	suite.Equal(structFields10.Type, godef.GoTypeString)
	suite.Len(structFields10.Tags, 0)
}

func (suite *CompileModelsTestSuite) TestMorpheModelToGoStructs_SuccessHook_Failure() {
	var featureFlag = "otherName"
	modelHooks := hook.CompileMorpheModel{
		OnCompileMorpheModelSuccess: func(allModelStructs []*godef.Struct) ([]*godef.Struct, error) {
			if featureFlag != "otherName" {
				return allModelStructs, nil
			}
			return nil, fmt.Errorf("compile model success hook error")
		},
	}

	modelsConfig := cfg.MorpheModelsConfig{
		Package: godef.Package{
			Path: "github.com/kaloseia/project/domain/models",
			Name: "models",
		},
		ReceiverName: "m",
	}

	model0 := yaml.Model{
		Name: "Basic",
		Fields: map[string]yaml.ModelField{
			"AutoIncrement": {
				Type: yaml.ModelFieldTypeAutoIncrement,
			},
			"Boolean": {
				Type: yaml.ModelFieldTypeBoolean,
			},
			"Date": {
				Type: yaml.ModelFieldTypeDate,
			},
			"Float": {
				Type: yaml.ModelFieldTypeFloat,
			},
			"Integer": {
				Type: yaml.ModelFieldTypeInteger,
			},
			"Protected": {
				Type: yaml.ModelFieldTypeProtected,
			},
			"Sealed": {
				Type: yaml.ModelFieldTypeSealed,
			},
			"String": {
				Type: yaml.ModelFieldTypeString,
			},
			"Time": {
				Type: yaml.ModelFieldTypeTime,
			},
			"UUID": {
				Type: yaml.ModelFieldTypeUUID,
				Attributes: []string{
					"immutable",
				},
			},
		},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {
				Fields: []string{
					"UUID",
				},
			},
		},
		Related: map[string]yaml.ModelRelation{},
	}

	allGoStructs, allStructsErr := compile.MorpheModelToGoStructs(modelHooks, modelsConfig, model0)

	suite.NotNil(allStructsErr)
	suite.ErrorContains(allStructsErr, "compile model success hook error")
	suite.Nil(allGoStructs)
}

func (suite *CompileModelsTestSuite) TestMorpheModelToGoStructs_FailureHook_NoPackagePath() {
	modelHooks := hook.CompileMorpheModel{
		OnCompileMorpheModelFailure: func(config cfg.MorpheModelsConfig, model yaml.Model, compileFailure error) error {
			return fmt.Errorf("Model %s: %w", model.Name, compileFailure)
		},
	}

	modelsConfig := cfg.MorpheModelsConfig{
		Package: godef.Package{
			Path: "",
			Name: "models",
		},
		ReceiverName: "m",
	}

	model0 := yaml.Model{
		Name: "Basic",
		Fields: map[string]yaml.ModelField{
			"AutoIncrement": {
				Type: yaml.ModelFieldTypeAutoIncrement,
			},
			"Boolean": {
				Type: yaml.ModelFieldTypeBoolean,
			},
			"Date": {
				Type: yaml.ModelFieldTypeDate,
			},
			"Float": {
				Type: yaml.ModelFieldTypeFloat,
			},
			"Integer": {
				Type: yaml.ModelFieldTypeInteger,
			},
			"Protected": {
				Type: yaml.ModelFieldTypeProtected,
			},
			"Sealed": {
				Type: yaml.ModelFieldTypeSealed,
			},
			"String": {
				Type: yaml.ModelFieldTypeString,
			},
			"Time": {
				Type: yaml.ModelFieldTypeTime,
			},
			"UUID": {
				Type: yaml.ModelFieldTypeUUID,
				Attributes: []string{
					"immutable",
				},
			},
		},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {
				Fields: []string{
					"UUID",
				},
			},
		},
		Related: map[string]yaml.ModelRelation{},
	}

	allGoStructs, allStructsErr := compile.MorpheModelToGoStructs(modelHooks, modelsConfig, model0)

	suite.NotNil(allStructsErr)
	suite.ErrorContains(allStructsErr, "Model Basic: models package path cannot be empty")
	suite.Nil(allGoStructs)
}
