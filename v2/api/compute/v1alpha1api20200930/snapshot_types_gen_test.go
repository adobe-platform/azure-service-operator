// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200930

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/compute/v1alpha1api20200930storage"
	"github.com/Azure/azure-service-operator/v2/api/compute/v1beta20200930storage"
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

func Test_Snapshot_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Snapshot to hub returns original",
		prop.ForAll(RunResourceConversionTestForSnapshot, SnapshotGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSnapshot tests if a specific instance of Snapshot round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSnapshot(subject Snapshot) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1beta20200930storage.Snapshot
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Snapshot
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Snapshot_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Snapshot to Snapshot via AssignPropertiesToSnapshot & AssignPropertiesFromSnapshot returns original",
		prop.ForAll(RunPropertyAssignmentTestForSnapshot, SnapshotGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSnapshot tests if a specific instance of Snapshot can be assigned to v1alpha1api20200930storage and back losslessly
func RunPropertyAssignmentTestForSnapshot(subject Snapshot) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200930storage.Snapshot
	err := copied.AssignPropertiesToSnapshot(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Snapshot
	err = actual.AssignPropertiesFromSnapshot(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Snapshot_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshot via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshot, SnapshotGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshot runs a test to see if a specific instance of Snapshot round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshot(subject Snapshot) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshot
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

// Generator of Snapshot instances for property testing - lazily instantiated by SnapshotGenerator()
var snapshotGenerator gopter.Gen

// SnapshotGenerator returns a generator of Snapshot instances for property testing.
func SnapshotGenerator() gopter.Gen {
	if snapshotGenerator != nil {
		return snapshotGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSnapshot(generators)
	snapshotGenerator = gen.Struct(reflect.TypeOf(Snapshot{}), generators)

	return snapshotGenerator
}

// AddRelatedPropertyGeneratorsForSnapshot is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshot(gens map[string]gopter.Gen) {
	gens["Spec"] = SnapshotsSpecGenerator()
	gens["Status"] = SnapshotStatusGenerator()
}

func Test_Snapshot_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Snapshot_Status to Snapshot_Status via AssignPropertiesToSnapshotStatus & AssignPropertiesFromSnapshotStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForSnapshotStatus, SnapshotStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSnapshotStatus tests if a specific instance of Snapshot_Status can be assigned to v1alpha1api20200930storage and back losslessly
func RunPropertyAssignmentTestForSnapshotStatus(subject Snapshot_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200930storage.Snapshot_Status
	err := copied.AssignPropertiesToSnapshotStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Snapshot_Status
	err = actual.AssignPropertiesFromSnapshotStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Snapshot_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshot_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotStatus, SnapshotStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotStatus runs a test to see if a specific instance of Snapshot_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotStatus(subject Snapshot_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshot_Status
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

// Generator of Snapshot_Status instances for property testing - lazily instantiated by SnapshotStatusGenerator()
var snapshotStatusGenerator gopter.Gen

// SnapshotStatusGenerator returns a generator of Snapshot_Status instances for property testing.
// We first initialize snapshotStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SnapshotStatusGenerator() gopter.Gen {
	if snapshotStatusGenerator != nil {
		return snapshotStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotStatus(generators)
	snapshotStatusGenerator = gen.Struct(reflect.TypeOf(Snapshot_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotStatus(generators)
	AddRelatedPropertyGeneratorsForSnapshotStatus(generators)
	snapshotStatusGenerator = gen.Struct(reflect.TypeOf(Snapshot_Status{}), generators)

	return snapshotStatusGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotStatus(gens map[string]gopter.Gen) {
	gens["DiskAccessId"] = gen.PtrOf(gen.AlphaString())
	gens["DiskSizeBytes"] = gen.PtrOf(gen.Int())
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["DiskState"] = gen.PtrOf(gen.OneConstOf(
		DiskState_StatusActiveSAS,
		DiskState_StatusActiveUpload,
		DiskState_StatusAttached,
		DiskState_StatusReadyToUpload,
		DiskState_StatusReserved,
		DiskState_StatusUnattached))
	gens["HyperVGeneration"] = gen.PtrOf(gen.OneConstOf(SnapshotPropertiesStatusHyperVGenerationV1, SnapshotPropertiesStatusHyperVGenerationV2))
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Incremental"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["ManagedBy"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["NetworkAccessPolicy"] = gen.PtrOf(gen.OneConstOf(NetworkAccessPolicy_StatusAllowAll, NetworkAccessPolicy_StatusAllowPrivate, NetworkAccessPolicy_StatusDenyAll))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(SnapshotPropertiesStatusOsTypeLinux, SnapshotPropertiesStatusOsTypeWindows))
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TimeCreated"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UniqueId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSnapshotStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshotStatus(gens map[string]gopter.Gen) {
	gens["CreationData"] = gen.PtrOf(CreationDataStatusGenerator())
	gens["Encryption"] = gen.PtrOf(EncryptionStatusGenerator())
	gens["EncryptionSettingsCollection"] = gen.PtrOf(EncryptionSettingsCollectionStatusGenerator())
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationStatusGenerator())
	gens["PurchasePlan"] = gen.PtrOf(PurchasePlanStatusGenerator())
	gens["Sku"] = gen.PtrOf(SnapshotSkuStatusGenerator())
}

func Test_Snapshots_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Snapshots_Spec to Snapshots_Spec via AssignPropertiesToSnapshotsSpec & AssignPropertiesFromSnapshotsSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForSnapshotsSpec, SnapshotsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSnapshotsSpec tests if a specific instance of Snapshots_Spec can be assigned to v1alpha1api20200930storage and back losslessly
func RunPropertyAssignmentTestForSnapshotsSpec(subject Snapshots_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200930storage.Snapshots_Spec
	err := copied.AssignPropertiesToSnapshotsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Snapshots_Spec
	err = actual.AssignPropertiesFromSnapshotsSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Snapshots_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshots_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotsSpec, SnapshotsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotsSpec runs a test to see if a specific instance of Snapshots_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotsSpec(subject Snapshots_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshots_Spec
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

// Generator of Snapshots_Spec instances for property testing - lazily instantiated by SnapshotsSpecGenerator()
var snapshotsSpecGenerator gopter.Gen

// SnapshotsSpecGenerator returns a generator of Snapshots_Spec instances for property testing.
// We first initialize snapshotsSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SnapshotsSpecGenerator() gopter.Gen {
	if snapshotsSpecGenerator != nil {
		return snapshotsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotsSpec(generators)
	snapshotsSpecGenerator = gen.Struct(reflect.TypeOf(Snapshots_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotsSpec(generators)
	AddRelatedPropertyGeneratorsForSnapshotsSpec(generators)
	snapshotsSpecGenerator = gen.Struct(reflect.TypeOf(Snapshots_Spec{}), generators)

	return snapshotsSpecGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotsSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["DiskState"] = gen.PtrOf(gen.OneConstOf(
		SnapshotPropertiesDiskStateActiveSAS,
		SnapshotPropertiesDiskStateActiveUpload,
		SnapshotPropertiesDiskStateAttached,
		SnapshotPropertiesDiskStateReadyToUpload,
		SnapshotPropertiesDiskStateReserved,
		SnapshotPropertiesDiskStateUnattached))
	gens["HyperVGeneration"] = gen.PtrOf(gen.OneConstOf(SnapshotPropertiesHyperVGenerationV1, SnapshotPropertiesHyperVGenerationV2))
	gens["Incremental"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["NetworkAccessPolicy"] = gen.PtrOf(gen.OneConstOf(SnapshotPropertiesNetworkAccessPolicyAllowAll, SnapshotPropertiesNetworkAccessPolicyAllowPrivate, SnapshotPropertiesNetworkAccessPolicyDenyAll))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(SnapshotPropertiesOsTypeLinux, SnapshotPropertiesOsTypeWindows))
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSnapshotsSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshotsSpec(gens map[string]gopter.Gen) {
	gens["CreationData"] = gen.PtrOf(CreationDataGenerator())
	gens["Encryption"] = gen.PtrOf(EncryptionGenerator())
	gens["EncryptionSettingsCollection"] = gen.PtrOf(EncryptionSettingsCollectionGenerator())
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationGenerator())
	gens["PurchasePlan"] = gen.PtrOf(PurchasePlanGenerator())
	gens["Sku"] = gen.PtrOf(SnapshotSkuGenerator())
}

func Test_SnapshotSku_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SnapshotSku to SnapshotSku via AssignPropertiesToSnapshotSku & AssignPropertiesFromSnapshotSku returns original",
		prop.ForAll(RunPropertyAssignmentTestForSnapshotSku, SnapshotSkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSnapshotSku tests if a specific instance of SnapshotSku can be assigned to v1alpha1api20200930storage and back losslessly
func RunPropertyAssignmentTestForSnapshotSku(subject SnapshotSku) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200930storage.SnapshotSku
	err := copied.AssignPropertiesToSnapshotSku(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SnapshotSku
	err = actual.AssignPropertiesFromSnapshotSku(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SnapshotSku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotSku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotSku, SnapshotSkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotSku runs a test to see if a specific instance of SnapshotSku round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotSku(subject SnapshotSku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotSku
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

// Generator of SnapshotSku instances for property testing - lazily instantiated by SnapshotSkuGenerator()
var snapshotSkuGenerator gopter.Gen

// SnapshotSkuGenerator returns a generator of SnapshotSku instances for property testing.
func SnapshotSkuGenerator() gopter.Gen {
	if snapshotSkuGenerator != nil {
		return snapshotSkuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotSku(generators)
	snapshotSkuGenerator = gen.Struct(reflect.TypeOf(SnapshotSku{}), generators)

	return snapshotSkuGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotSku(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(SnapshotSkuNamePremiumLRS, SnapshotSkuNameStandardLRS, SnapshotSkuNameStandardZRS))
}

func Test_SnapshotSku_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SnapshotSku_Status to SnapshotSku_Status via AssignPropertiesToSnapshotSkuStatus & AssignPropertiesFromSnapshotSkuStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForSnapshotSkuStatus, SnapshotSkuStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSnapshotSkuStatus tests if a specific instance of SnapshotSku_Status can be assigned to v1alpha1api20200930storage and back losslessly
func RunPropertyAssignmentTestForSnapshotSkuStatus(subject SnapshotSku_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200930storage.SnapshotSku_Status
	err := copied.AssignPropertiesToSnapshotSkuStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SnapshotSku_Status
	err = actual.AssignPropertiesFromSnapshotSkuStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SnapshotSku_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotSku_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotSkuStatus, SnapshotSkuStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotSkuStatus runs a test to see if a specific instance of SnapshotSku_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotSkuStatus(subject SnapshotSku_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotSku_Status
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

// Generator of SnapshotSku_Status instances for property testing - lazily instantiated by SnapshotSkuStatusGenerator()
var snapshotSkuStatusGenerator gopter.Gen

// SnapshotSkuStatusGenerator returns a generator of SnapshotSku_Status instances for property testing.
func SnapshotSkuStatusGenerator() gopter.Gen {
	if snapshotSkuStatusGenerator != nil {
		return snapshotSkuStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotSkuStatus(generators)
	snapshotSkuStatusGenerator = gen.Struct(reflect.TypeOf(SnapshotSku_Status{}), generators)

	return snapshotSkuStatusGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotSkuStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotSkuStatus(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(SnapshotSkuStatusNamePremiumLRS, SnapshotSkuStatusNameStandardLRS, SnapshotSkuStatusNameStandardZRS))
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}