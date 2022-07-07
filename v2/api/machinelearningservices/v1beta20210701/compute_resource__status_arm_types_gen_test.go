// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210701

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_ComputeResource_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ComputeResource_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForComputeResourceStatusARM, ComputeResourceStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForComputeResourceStatusARM runs a test to see if a specific instance of ComputeResource_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForComputeResourceStatusARM(subject ComputeResource_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ComputeResource_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ComputeResource_StatusARM instances for property testing - lazily instantiated by
// ComputeResourceStatusARMGenerator()
var computeResourceStatusARMGenerator gopter.Gen

// ComputeResourceStatusARMGenerator returns a generator of ComputeResource_StatusARM instances for property testing.
// We first initialize computeResourceStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ComputeResourceStatusARMGenerator() gopter.Gen {
	if computeResourceStatusARMGenerator != nil {
		return computeResourceStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComputeResourceStatusARM(generators)
	computeResourceStatusARMGenerator = gen.Struct(reflect.TypeOf(ComputeResource_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComputeResourceStatusARM(generators)
	AddRelatedPropertyGeneratorsForComputeResourceStatusARM(generators)
	computeResourceStatusARMGenerator = gen.Struct(reflect.TypeOf(ComputeResource_StatusARM{}), generators)

	return computeResourceStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForComputeResourceStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForComputeResourceStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForComputeResourceStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForComputeResourceStatusARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(IdentityStatusARMGenerator())
	gens["Properties"] = gen.PtrOf(ComputeStatusARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuStatusARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_Compute_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Compute_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForComputeStatusARM, ComputeStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForComputeStatusARM runs a test to see if a specific instance of Compute_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForComputeStatusARM(subject Compute_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Compute_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Compute_StatusARM instances for property testing - lazily instantiated by ComputeStatusARMGenerator()
var computeStatusARMGenerator gopter.Gen

// ComputeStatusARMGenerator returns a generator of Compute_StatusARM instances for property testing.
// We first initialize computeStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ComputeStatusARMGenerator() gopter.Gen {
	if computeStatusARMGenerator != nil {
		return computeStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComputeStatusARM(generators)
	computeStatusARMGenerator = gen.Struct(reflect.TypeOf(Compute_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComputeStatusARM(generators)
	AddRelatedPropertyGeneratorsForComputeStatusARM(generators)
	computeStatusARMGenerator = gen.Struct(reflect.TypeOf(Compute_StatusARM{}), generators)

	return computeStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForComputeStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForComputeStatusARM(gens map[string]gopter.Gen) {
	gens["ComputeLocation"] = gen.PtrOf(gen.AlphaString())
	gens["ComputeType"] = gen.PtrOf(gen.OneConstOf(
		ComputeType_StatusAKS,
		ComputeType_StatusAmlCompute,
		ComputeType_StatusComputeInstance,
		ComputeType_StatusDataFactory,
		ComputeType_StatusDataLakeAnalytics,
		ComputeType_StatusDatabricks,
		ComputeType_StatusHDInsight,
		ComputeType_StatusKubernetes,
		ComputeType_StatusSynapseSpark,
		ComputeType_StatusVirtualMachine))
	gens["CreatedOn"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["IsAttachedCompute"] = gen.PtrOf(gen.Bool())
	gens["ModifiedOn"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ComputeStatusProvisioningStateCanceled,
		ComputeStatusProvisioningStateCreating,
		ComputeStatusProvisioningStateDeleting,
		ComputeStatusProvisioningStateFailed,
		ComputeStatusProvisioningStateSucceeded,
		ComputeStatusProvisioningStateUnknown,
		ComputeStatusProvisioningStateUpdating))
	gens["ResourceId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForComputeStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForComputeStatusARM(gens map[string]gopter.Gen) {
	gens["ProvisioningErrors"] = gen.SliceOf(ErrorResponseStatusARMGenerator())
}

func Test_Identity_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentityStatusARM, IdentityStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentityStatusARM runs a test to see if a specific instance of Identity_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentityStatusARM(subject Identity_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Identity_StatusARM instances for property testing - lazily instantiated by IdentityStatusARMGenerator()
var identityStatusARMGenerator gopter.Gen

// IdentityStatusARMGenerator returns a generator of Identity_StatusARM instances for property testing.
// We first initialize identityStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IdentityStatusARMGenerator() gopter.Gen {
	if identityStatusARMGenerator != nil {
		return identityStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityStatusARM(generators)
	identityStatusARMGenerator = gen.Struct(reflect.TypeOf(Identity_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityStatusARM(generators)
	AddRelatedPropertyGeneratorsForIdentityStatusARM(generators)
	identityStatusARMGenerator = gen.Struct(reflect.TypeOf(Identity_StatusARM{}), generators)

	return identityStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentityStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentityStatusARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		IdentityStatusTypeNone,
		IdentityStatusTypeSystemAssigned,
		IdentityStatusTypeSystemAssignedUserAssigned,
		IdentityStatusTypeUserAssigned))
}

// AddRelatedPropertyGeneratorsForIdentityStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIdentityStatusARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(gen.AlphaString(), UserAssignedIdentityStatusARMGenerator())
}

func Test_Sku_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuStatusARM, SkuStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuStatusARM runs a test to see if a specific instance of Sku_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuStatusARM(subject Sku_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Sku_StatusARM instances for property testing - lazily instantiated by SkuStatusARMGenerator()
var skuStatusARMGenerator gopter.Gen

// SkuStatusARMGenerator returns a generator of Sku_StatusARM instances for property testing.
func SkuStatusARMGenerator() gopter.Gen {
	if skuStatusARMGenerator != nil {
		return skuStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuStatusARM(generators)
	skuStatusARMGenerator = gen.Struct(reflect.TypeOf(Sku_StatusARM{}), generators)

	return skuStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuStatusARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_SystemData_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemDataStatusARM, SystemDataStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemDataStatusARM runs a test to see if a specific instance of SystemData_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemDataStatusARM(subject SystemData_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SystemData_StatusARM instances for property testing - lazily instantiated by
// SystemDataStatusARMGenerator()
var systemDataStatusARMGenerator gopter.Gen

// SystemDataStatusARMGenerator returns a generator of SystemData_StatusARM instances for property testing.
func SystemDataStatusARMGenerator() gopter.Gen {
	if systemDataStatusARMGenerator != nil {
		return systemDataStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemDataStatusARM(generators)
	systemDataStatusARMGenerator = gen.Struct(reflect.TypeOf(SystemData_StatusARM{}), generators)

	return systemDataStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemDataStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemDataStatusARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemDataStatusCreatedByTypeApplication,
		SystemDataStatusCreatedByTypeKey,
		SystemDataStatusCreatedByTypeManagedIdentity,
		SystemDataStatusCreatedByTypeUser))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemDataStatusLastModifiedByTypeApplication,
		SystemDataStatusLastModifiedByTypeKey,
		SystemDataStatusLastModifiedByTypeManagedIdentity,
		SystemDataStatusLastModifiedByTypeUser))
}

func Test_ErrorResponse_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ErrorResponse_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForErrorResponseStatusARM, ErrorResponseStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForErrorResponseStatusARM runs a test to see if a specific instance of ErrorResponse_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForErrorResponseStatusARM(subject ErrorResponse_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ErrorResponse_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ErrorResponse_StatusARM instances for property testing - lazily instantiated by
// ErrorResponseStatusARMGenerator()
var errorResponseStatusARMGenerator gopter.Gen

// ErrorResponseStatusARMGenerator returns a generator of ErrorResponse_StatusARM instances for property testing.
func ErrorResponseStatusARMGenerator() gopter.Gen {
	if errorResponseStatusARMGenerator != nil {
		return errorResponseStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForErrorResponseStatusARM(generators)
	errorResponseStatusARMGenerator = gen.Struct(reflect.TypeOf(ErrorResponse_StatusARM{}), generators)

	return errorResponseStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForErrorResponseStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForErrorResponseStatusARM(gens map[string]gopter.Gen) {
	gens["Error"] = gen.PtrOf(ErrorDetailStatusARMGenerator())
}

func Test_UserAssignedIdentity_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityStatusARM, UserAssignedIdentityStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityStatusARM runs a test to see if a specific instance of UserAssignedIdentity_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityStatusARM(subject UserAssignedIdentity_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of UserAssignedIdentity_StatusARM instances for property testing - lazily instantiated by
// UserAssignedIdentityStatusARMGenerator()
var userAssignedIdentityStatusARMGenerator gopter.Gen

// UserAssignedIdentityStatusARMGenerator returns a generator of UserAssignedIdentity_StatusARM instances for property testing.
func UserAssignedIdentityStatusARMGenerator() gopter.Gen {
	if userAssignedIdentityStatusARMGenerator != nil {
		return userAssignedIdentityStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentityStatusARM(generators)
	userAssignedIdentityStatusARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_StatusARM{}), generators)

	return userAssignedIdentityStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentityStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentityStatusARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

func Test_ErrorDetail_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ErrorDetail_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForErrorDetailStatusARM, ErrorDetailStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForErrorDetailStatusARM runs a test to see if a specific instance of ErrorDetail_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForErrorDetailStatusARM(subject ErrorDetail_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ErrorDetail_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ErrorDetail_StatusARM instances for property testing - lazily instantiated by
// ErrorDetailStatusARMGenerator()
var errorDetailStatusARMGenerator gopter.Gen

// ErrorDetailStatusARMGenerator returns a generator of ErrorDetail_StatusARM instances for property testing.
// We first initialize errorDetailStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ErrorDetailStatusARMGenerator() gopter.Gen {
	if errorDetailStatusARMGenerator != nil {
		return errorDetailStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorDetailStatusARM(generators)
	errorDetailStatusARMGenerator = gen.Struct(reflect.TypeOf(ErrorDetail_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorDetailStatusARM(generators)
	AddRelatedPropertyGeneratorsForErrorDetailStatusARM(generators)
	errorDetailStatusARMGenerator = gen.Struct(reflect.TypeOf(ErrorDetail_StatusARM{}), generators)

	return errorDetailStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForErrorDetailStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForErrorDetailStatusARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForErrorDetailStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForErrorDetailStatusARM(gens map[string]gopter.Gen) {
	gens["AdditionalInfo"] = gen.SliceOf(ErrorAdditionalInfoStatusARMGenerator())
	gens["Details"] = gen.SliceOf(ErrorDetailStatusUnrolledARMGenerator())
}

func Test_ErrorAdditionalInfo_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ErrorAdditionalInfo_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForErrorAdditionalInfoStatusARM, ErrorAdditionalInfoStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForErrorAdditionalInfoStatusARM runs a test to see if a specific instance of ErrorAdditionalInfo_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForErrorAdditionalInfoStatusARM(subject ErrorAdditionalInfo_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ErrorAdditionalInfo_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ErrorAdditionalInfo_StatusARM instances for property testing - lazily instantiated by
// ErrorAdditionalInfoStatusARMGenerator()
var errorAdditionalInfoStatusARMGenerator gopter.Gen

// ErrorAdditionalInfoStatusARMGenerator returns a generator of ErrorAdditionalInfo_StatusARM instances for property testing.
func ErrorAdditionalInfoStatusARMGenerator() gopter.Gen {
	if errorAdditionalInfoStatusARMGenerator != nil {
		return errorAdditionalInfoStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorAdditionalInfoStatusARM(generators)
	errorAdditionalInfoStatusARMGenerator = gen.Struct(reflect.TypeOf(ErrorAdditionalInfo_StatusARM{}), generators)

	return errorAdditionalInfoStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForErrorAdditionalInfoStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForErrorAdditionalInfoStatusARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_ErrorDetail_Status_UnrolledARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ErrorDetail_Status_UnrolledARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForErrorDetailStatusUnrolledARM, ErrorDetailStatusUnrolledARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForErrorDetailStatusUnrolledARM runs a test to see if a specific instance of ErrorDetail_Status_UnrolledARM round trips to JSON and back losslessly
func RunJSONSerializationTestForErrorDetailStatusUnrolledARM(subject ErrorDetail_Status_UnrolledARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ErrorDetail_Status_UnrolledARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ErrorDetail_Status_UnrolledARM instances for property testing - lazily instantiated by
// ErrorDetailStatusUnrolledARMGenerator()
var errorDetailStatusUnrolledARMGenerator gopter.Gen

// ErrorDetailStatusUnrolledARMGenerator returns a generator of ErrorDetail_Status_UnrolledARM instances for property testing.
// We first initialize errorDetailStatusUnrolledARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ErrorDetailStatusUnrolledARMGenerator() gopter.Gen {
	if errorDetailStatusUnrolledARMGenerator != nil {
		return errorDetailStatusUnrolledARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorDetailStatusUnrolledARM(generators)
	errorDetailStatusUnrolledARMGenerator = gen.Struct(reflect.TypeOf(ErrorDetail_Status_UnrolledARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorDetailStatusUnrolledARM(generators)
	AddRelatedPropertyGeneratorsForErrorDetailStatusUnrolledARM(generators)
	errorDetailStatusUnrolledARMGenerator = gen.Struct(reflect.TypeOf(ErrorDetail_Status_UnrolledARM{}), generators)

	return errorDetailStatusUnrolledARMGenerator
}

// AddIndependentPropertyGeneratorsForErrorDetailStatusUnrolledARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForErrorDetailStatusUnrolledARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForErrorDetailStatusUnrolledARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForErrorDetailStatusUnrolledARM(gens map[string]gopter.Gen) {
	gens["AdditionalInfo"] = gen.SliceOf(ErrorAdditionalInfoStatusARMGenerator())
}