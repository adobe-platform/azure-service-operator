// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180901

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

func Test_PrivateZone_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateZone_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateZoneStatusARM, PrivateZoneStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateZoneStatusARM runs a test to see if a specific instance of PrivateZone_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateZoneStatusARM(subject PrivateZone_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateZone_StatusARM
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

// Generator of PrivateZone_StatusARM instances for property testing - lazily instantiated by
// PrivateZoneStatusARMGenerator()
var privateZoneStatusARMGenerator gopter.Gen

// PrivateZoneStatusARMGenerator returns a generator of PrivateZone_StatusARM instances for property testing.
// We first initialize privateZoneStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateZoneStatusARMGenerator() gopter.Gen {
	if privateZoneStatusARMGenerator != nil {
		return privateZoneStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateZoneStatusARM(generators)
	privateZoneStatusARMGenerator = gen.Struct(reflect.TypeOf(PrivateZone_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateZoneStatusARM(generators)
	AddRelatedPropertyGeneratorsForPrivateZoneStatusARM(generators)
	privateZoneStatusARMGenerator = gen.Struct(reflect.TypeOf(PrivateZone_StatusARM{}), generators)

	return privateZoneStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateZoneStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateZoneStatusARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateZoneStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateZoneStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrivateZonePropertiesStatusARMGenerator())
}

func Test_PrivateZoneProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateZoneProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateZonePropertiesStatusARM, PrivateZonePropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateZonePropertiesStatusARM runs a test to see if a specific instance of PrivateZoneProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateZonePropertiesStatusARM(subject PrivateZoneProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateZoneProperties_StatusARM
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

// Generator of PrivateZoneProperties_StatusARM instances for property testing - lazily instantiated by
// PrivateZonePropertiesStatusARMGenerator()
var privateZonePropertiesStatusARMGenerator gopter.Gen

// PrivateZonePropertiesStatusARMGenerator returns a generator of PrivateZoneProperties_StatusARM instances for property testing.
func PrivateZonePropertiesStatusARMGenerator() gopter.Gen {
	if privateZonePropertiesStatusARMGenerator != nil {
		return privateZonePropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateZonePropertiesStatusARM(generators)
	privateZonePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(PrivateZoneProperties_StatusARM{}), generators)

	return privateZonePropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateZonePropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateZonePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["MaxNumberOfRecordSets"] = gen.PtrOf(gen.Int())
	gens["MaxNumberOfVirtualNetworkLinks"] = gen.PtrOf(gen.Int())
	gens["MaxNumberOfVirtualNetworkLinksWithRegistration"] = gen.PtrOf(gen.Int())
	gens["NumberOfRecordSets"] = gen.PtrOf(gen.Int())
	gens["NumberOfVirtualNetworkLinks"] = gen.PtrOf(gen.Int())
	gens["NumberOfVirtualNetworkLinksWithRegistration"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		PrivateZonePropertiesStatusProvisioningStateCanceled,
		PrivateZonePropertiesStatusProvisioningStateCreating,
		PrivateZonePropertiesStatusProvisioningStateDeleting,
		PrivateZonePropertiesStatusProvisioningStateFailed,
		PrivateZonePropertiesStatusProvisioningStateSucceeded,
		PrivateZonePropertiesStatusProvisioningStateUpdating))
}