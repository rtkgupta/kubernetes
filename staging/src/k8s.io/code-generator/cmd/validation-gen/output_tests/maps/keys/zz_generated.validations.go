//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by validation-gen. DO NOT EDIT.

package eachkey

import (
	context "context"
	fmt "fmt"

	operation "k8s.io/apimachinery/pkg/api/operation"
	safe "k8s.io/apimachinery/pkg/api/safe"
	validate "k8s.io/apimachinery/pkg/api/validate"
	field "k8s.io/apimachinery/pkg/util/validation/field"
	testscheme "k8s.io/code-generator/cmd/validation-gen/testscheme"
)

func init() { localSchemeBuilder.Register(RegisterValidations) }

// RegisterValidations adds validation functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterValidations(scheme *testscheme.Scheme) error {
	scheme.AddValidationFunc((*Struct)(nil), func(ctx context.Context, op operation.Operation, obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_Struct(ctx, op, nil /* fldPath */, obj.(*Struct), safe.Cast[*Struct](oldObj))
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	return nil
}

func Validate_Struct(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *Struct) (errs field.ErrorList) {
	// type Struct
	errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "type Struct")...)

	// field Struct.TypeMeta has no validation

	// field Struct.MapField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj map[string]string) (errs field.ErrorList) {
			errs = append(errs, validate.EachMapKey(ctx, op, fldPath, obj, oldObj, func(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *string) field.ErrorList {
				return validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.MapField(keys)")
			})...)
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.MapField")...)
			return
		}(fldPath.Child("mapField"), obj.MapField, safe.Field(oldObj, func(oldObj *Struct) map[string]string { return oldObj.MapField }))...)

	// field Struct.MapTypedefField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj map[UnvalidatedStringType]string) (errs field.ErrorList) {
			errs = append(errs, validate.EachMapKey(ctx, op, fldPath, obj, oldObj, func(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *UnvalidatedStringType) field.ErrorList {
				return validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.MapTypedefField(keys)")
			})...)
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.MapTypedefField")...)
			return
		}(fldPath.Child("mapTypedefField"), obj.MapTypedefField, safe.Field(oldObj, func(oldObj *Struct) map[UnvalidatedStringType]string { return oldObj.MapTypedefField }))...)

	// field Struct.MapValidatedTypedefField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj map[ValidatedStringType]string) (errs field.ErrorList) {
			errs = append(errs, validate.EachMapKey(ctx, op, fldPath, obj, oldObj, func(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *ValidatedStringType) field.ErrorList {
				return validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.MapValidatedTypedefField(keys)")
			})...)
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.MapValidatedTypedefField")...)
			errs = append(errs, validate.EachMapKey(ctx, op, fldPath, obj, oldObj, Validate_ValidatedStringType)...)
			return
		}(fldPath.Child("mapValidatedTypedefField"), obj.MapValidatedTypedefField, safe.Field(oldObj, func(oldObj *Struct) map[ValidatedStringType]string { return oldObj.MapValidatedTypedefField }))...)

	// field Struct.MapTypeField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj UnvalidatedMapType) (errs field.ErrorList) {
			errs = append(errs, validate.EachMapKey(ctx, op, fldPath, obj, oldObj, func(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *string) field.ErrorList {
				return validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.MapTypeField(keys)")
			})...)
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.MapTypeField")...)
			return
		}(fldPath.Child("mapTypeField"), obj.MapTypeField, safe.Field(oldObj, func(oldObj *Struct) UnvalidatedMapType { return oldObj.MapTypeField }))...)

	// field Struct.ValidatedMapTypeField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj ValidatedMapType) (errs field.ErrorList) {
			errs = append(errs, validate.EachMapKey(ctx, op, fldPath, obj, oldObj, func(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *string) field.ErrorList {
				return validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.ValidatedMapTypeField(keys)")
			})...)
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.ValidatedMapTypeField")...)
			errs = append(errs, Validate_ValidatedMapType(ctx, op, fldPath, obj, oldObj)...)
			return
		}(fldPath.Child("validatedMapTypeField"), obj.ValidatedMapTypeField, safe.Field(oldObj, func(oldObj *Struct) ValidatedMapType { return oldObj.ValidatedMapTypeField }))...)

	return errs
}

func Validate_ValidatedMapType(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj ValidatedMapType) (errs field.ErrorList) {
	// type ValidatedMapType
	errs = append(errs, validate.EachMapKey(ctx, op, fldPath, obj, oldObj, func(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *string) field.ErrorList {
		return validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "type ValidatedMapType(keys)")
	})...)
	errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "ValidatedMapType")...)

	return errs
}

func Validate_ValidatedStringType(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *ValidatedStringType) (errs field.ErrorList) {
	// type ValidatedStringType
	errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "ValidatedStringType")...)

	return errs
}
