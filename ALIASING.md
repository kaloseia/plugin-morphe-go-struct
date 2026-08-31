# Aliasing Support Implementation Plan for plugin-morphe-go-struct

## Overview

This document outlines the plan to add relationship aliasing support to the Go structs compilation plugin. Aliasing allows models and entities to have multiple relationships to the same target type with different semantic names.

## Background

### What is Aliasing?

Aliasing enables declaring multiple relationships to the same target model with different field names:

```yaml
# Model with aliased relationships to the same Contact model
name: Person
related:
  WorkContact:
    type: ForOne
    aliased: Contact
  PersonalContact:
    type: ForOne
    aliased: Contact
```

### Relationship Type Coverage

1. **ForOne Aliasing**: One-to-one forward relationships with multiple targets
2. **ForMany Aliasing**: One-to-many forward relationships with multiple targets
3. **HasOne/HasMany Aliasing**: Inverse relationships (IDs only in model)
4. **Polymorphic Aliasing**: ForOnePoly, ForManyPoly, HasOnePoly, HasManyPoly with `through` attribute

## Implementation Plan

### Phase 1: Core Infrastructure

#### 1.1 Import yamlops utilities
- Use `yamlops.GetRelationTargetName(relationshipName, aliasedField)` to resolve actual target names
- Already available in morphe-go, needs integration into plugin

#### 1.2 Add validation for aliased relationships
- Validate that aliased target models exist in registry
- Check for circular/invalid alias patterns
- Integrate into existing validation flows

### Phase 2: Model Compilation (compile_models.go)

#### 2.1 Update getRelatedGoFieldsForMorpheModel
**Location**: compile_models.go ~line 148

**Current behavior**: 
- Uses relationshipName directly as both field name and lookup key
- Doesn't handle aliasing

**Required changes**:
- Use `yamlops.GetRelationTargetName()` to get actual target model name for lookup
- Keep using relationshipName (from Related key) for field naming (maintains semantic names)
- This creates proper separation: field names use relationship names, type references use aliased target

**Example**:
```go
// Before: WorkContact lookup uses "WorkContact" for both name and type
relatedModelDef, _ := r.GetModel("WorkContact")  // ❌ Wrong

// After: Lookup uses actual target, name stays semantic
targetName := yamlops.GetRelationTargetName("WorkContact", relation.Aliased)  // "Contact"
relatedModelDef, _ := r.GetModel(targetName)  // ✅ Correct
// But field is still named "WorkContact" or "WorkContacts"
```

#### 2.2 Update getRelatedGoFieldForMorpheModelPrimaryID
**Location**: compile_models.go ~line 171

**Changes**:
- Accept aliased field to determine target model name
- Pass target model name (not relationship name) for type resolution
- Keep field naming based on relationship name for semantics

#### 2.3 Update getRelatedGoFieldForMorpheModel  
**Location**: compile_models.go ~line 210

**Changes**:
- Receive actual target model name as parameter
- Use it to construct correct GoTypeStruct name
- Keep field naming based on relationship name

#### 2.4 Add validation: validateAliasedRelations
**Location**: compile_models.go (similar to psql plugin pattern)

**Purpose**: Ensure aliased target models exist in registry
**Called from**: morpheModelToGoStructs function

### Phase 3: Entity Compilation (compile_entities.go)

#### 3.1 Update getRelatedGoFieldsForMorpheEntity
**Location**: compile_entities.go ~line 130

**Required changes**:
- Use `yamlops.GetRelationTargetName()` to resolve actual target entity
- Keep semantic field naming using relationship name
- Lookup target entity using resolved name

#### 3.2 Update getRelatedGoFieldForEntityPrimaryID
**Location**: compile_entities.go ~line 159

**Changes**:
- Accept aliased field parameter
- Pass resolved target name for entity lookup
- Keep field naming semantic (relationship name based)

#### 3.3 Update getRelatedGoFieldForEntity
**Location**: compile_entities.go ~line 200

**Changes**:
- Receive actual target entity name for GoTypeStruct
- Keep field naming based on relationship name
- Properly handle ForOne/ForMany/HasOne/HasMany with aliasing

#### 3.4 Extend field type resolution
**Location**: compile_entities.go - getModelFieldType function

**Context**: Entities reference model fields via paths like "Model.Field.Related"
**Changes needed**: When traversing relationships, account for aliasing in the path

### Phase 4: Polymorphic Aliasing Support

#### 4.1 ForOnePoly/ForManyPoly with aliasing
**Pattern**: `type: ForOnePoly` with `through:` and `aliased:`

**Implementation**:
- Similar structure to regular aliasing
- Resolve target name using aliased field
- Generate proper struct fields with semantic names

#### 4.2 HasOnePoly/HasManyPoly with aliasing  
**Pattern**: `type: HasOnePoly` with `through:` and `aliased:`

**Implementation**:
- Resolve target using aliased field
- Generate ID fields only (inverse pattern)
- Maintain semantic naming

### Phase 5: Test Data Setup (testdata/)

#### 5.1 Create aliasing test registry
**Directory**: testdata/registry/aliasing/

**Models needed**:
- Person.mod - with ForOne aliased relationships to Contact
- Contact.mod - with HasMany aliased relationships to Person
- Project.mod - with HasMany aliased relationships to Person

**Entities needed**:
- Person.ent - references aliased relationships
- Contact.ent - inverse aliased relationships

**Structures needed** (if any polymorphic examples):
- Structures for polymorphic aliasing tests

#### 5.2 Create ground-truth files
**Directory**: testdata/ground-truth/compile-aliasing/

**Expected outputs**:
- models/person.go - with WorkContact/PersonalContact fields
- models/contact.go - with semantic inverted relationship names
- models/project.go - with semantic aliased relationships
- entities/person.go - with aliased entity relationships
- entities/contact.go
- models/*_id_primary.go (identifier structs)
- entities/*_id_primary.go (identifier structs)

### Phase 6: Integration Testing (compile_test.go)

#### 6.1 Extend existing TestMorpheToGo
- Add aliasing model/entity registry paths to SetupTest
- Extend test with new assertions for aliasing files
- Keep as single integrated test (don't create separate test)
- Compare generated outputs against ground-truth aliasing files

**Additions to TestMorpheToGo**:
1. Add aliasing registry paths to config
2. Add working dir paths for aliasing output
3. Assert aliasing model files exist and match ground-truth
4. Assert aliasing entity files exist and match ground-truth
5. Assert all identifier struct variations exist

#### 6.2 Specific assertion coverage
- Person.go with WorkContact, PersonalContact fields
- Contact.go with inverse relationships
- Project.go with aliased HasMany relationships
- All identifier structs compile correctly
- Type references are correct (aliased target names)
- Field names are semantic (relationship names)

## Implementation Order

1. **Phase 1**: Infrastructure (yamlops integration)
2. **Phase 2**: Model compilation (core logic)
3. **Phase 3**: Entity compilation 
4. **Phase 4**: Polymorphic support
5. **Phase 5**: Test data setup
6. **Phase 6**: Integration tests

## Key Design Principles

### Semantic Field Naming
- Field names use relationship names (WorkContact, PersonalContact)
- This makes the Go code readable and semantically clear

### Type References Use Aliases
- Struct types reference the actual target model (Contact, Project)
- This ensures correct relationships and proper Go compilation

### Relationship Name Resolution
```
getRelatedGoFieldsForMorpheModel(model.Related)
  ↓
for relationshipName, relation := range model.Related
  ↓
targetName := GetRelationTargetName(relationshipName, relation.Aliased)
  ↓
lookup model using targetName
  ↓
but create fields using relationshipName
```

## Files to Modify

### Core Compilation Files
- [ ] compile_models.go - getRelatedGoFieldsForMorpheModel, getRelatedGoFieldForMorpheModelPrimaryID, getRelatedGoFieldForMorpheModel, morpheModelToGoStructs
- [ ] compile_entities.go - getRelatedGoFieldsForMorpheEntity, getRelatedGoFieldForEntityPrimaryID, getRelatedGoFieldForEntity, morpheEntityToGoStructs
- [ ] compile.go - May need validation integration

### Test Data
- [ ] testdata/registry/aliasing/ (new directory structure)
- [ ] testdata/ground-truth/compile-aliasing/ (new directory structure)

### Test Files
- [ ] compile_test.go - Extend TestMorpheToGo with aliasing coverage

## Success Criteria

1. ✅ All models with ForOne aliased relationships compile correctly
2. ✅ All models with ForMany aliased relationships compile correctly  
3. ✅ Inverse (HasOne/HasMany) aliased relationships work
4. ✅ Polymorphic aliased relationships work
5. ✅ Field names are semantic (relationship names)
6. ✅ Type references point to correct aliased targets
7. ✅ All identifier structs generate properly
8. ✅ Integration test covers all aliasing scenarios
9. ✅ No existing tests break
10. ✅ Validation catches missing aliased targets

## Notes

- Psql plugin already has working aliasing implementation as reference
- Uses same yamlops.GetRelationTargetName utility
- Focus on maintaining backward compatibility (non-aliased relationships unchanged)
- Semantic field naming is critical for Go readability
