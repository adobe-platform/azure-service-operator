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

func Test_Image_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Image_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImage_STATUS_ARM, Image_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImage_STATUS_ARM runs a test to see if a specific instance of Image_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImage_STATUS_ARM(subject Image_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Image_STATUS_ARM
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

// Generator of Image_STATUS_ARM instances for property testing - lazily instantiated by Image_STATUS_ARMGenerator()
var image_STATUS_ARMGenerator gopter.Gen

// Image_STATUS_ARMGenerator returns a generator of Image_STATUS_ARM instances for property testing.
// We first initialize image_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Image_STATUS_ARMGenerator() gopter.Gen {
	if image_STATUS_ARMGenerator != nil {
		return image_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImage_STATUS_ARM(generators)
	image_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Image_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImage_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForImage_STATUS_ARM(generators)
	image_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Image_STATUS_ARM{}), generators)

	return image_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForImage_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImage_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImage_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImage_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUS_ARMGenerator())
	gens["Properties"] = gen.PtrOf(ImageProperties_STATUS_ARMGenerator())
}

func Test_ExtendedLocation_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExtendedLocation_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExtendedLocation_STATUS_ARM, ExtendedLocation_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExtendedLocation_STATUS_ARM runs a test to see if a specific instance of ExtendedLocation_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExtendedLocation_STATUS_ARM(subject ExtendedLocation_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExtendedLocation_STATUS_ARM
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

// Generator of ExtendedLocation_STATUS_ARM instances for property testing - lazily instantiated by
// ExtendedLocation_STATUS_ARMGenerator()
var extendedLocation_STATUS_ARMGenerator gopter.Gen

// ExtendedLocation_STATUS_ARMGenerator returns a generator of ExtendedLocation_STATUS_ARM instances for property testing.
func ExtendedLocation_STATUS_ARMGenerator() gopter.Gen {
	if extendedLocation_STATUS_ARMGenerator != nil {
		return extendedLocation_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtendedLocation_STATUS_ARM(generators)
	extendedLocation_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ExtendedLocation_STATUS_ARM{}), generators)

	return extendedLocation_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForExtendedLocation_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExtendedLocation_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(ExtendedLocationType_STATUS_EdgeZone))
}

func Test_ImageProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageProperties_STATUS_ARM, ImageProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageProperties_STATUS_ARM runs a test to see if a specific instance of ImageProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageProperties_STATUS_ARM(subject ImageProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageProperties_STATUS_ARM
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

// Generator of ImageProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ImageProperties_STATUS_ARMGenerator()
var imageProperties_STATUS_ARMGenerator gopter.Gen

// ImageProperties_STATUS_ARMGenerator returns a generator of ImageProperties_STATUS_ARM instances for property testing.
// We first initialize imageProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageProperties_STATUS_ARMGenerator() gopter.Gen {
	if imageProperties_STATUS_ARMGenerator != nil {
		return imageProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageProperties_STATUS_ARM(generators)
	imageProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ImageProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForImageProperties_STATUS_ARM(generators)
	imageProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ImageProperties_STATUS_ARM{}), generators)

	return imageProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForImageProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["HyperVGeneration"] = gen.PtrOf(gen.OneConstOf(HyperVGenerationType_STATUS_V1, HyperVGenerationType_STATUS_V2))
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImageProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["SourceVirtualMachine"] = gen.PtrOf(SubResource_STATUS_ARMGenerator())
	gens["StorageProfile"] = gen.PtrOf(ImageStorageProfile_STATUS_ARMGenerator())
}

func Test_ImageStorageProfile_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageStorageProfile_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageStorageProfile_STATUS_ARM, ImageStorageProfile_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageStorageProfile_STATUS_ARM runs a test to see if a specific instance of ImageStorageProfile_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageStorageProfile_STATUS_ARM(subject ImageStorageProfile_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageStorageProfile_STATUS_ARM
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

// Generator of ImageStorageProfile_STATUS_ARM instances for property testing - lazily instantiated by
// ImageStorageProfile_STATUS_ARMGenerator()
var imageStorageProfile_STATUS_ARMGenerator gopter.Gen

// ImageStorageProfile_STATUS_ARMGenerator returns a generator of ImageStorageProfile_STATUS_ARM instances for property testing.
// We first initialize imageStorageProfile_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageStorageProfile_STATUS_ARMGenerator() gopter.Gen {
	if imageStorageProfile_STATUS_ARMGenerator != nil {
		return imageStorageProfile_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStorageProfile_STATUS_ARM(generators)
	imageStorageProfile_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ImageStorageProfile_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStorageProfile_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForImageStorageProfile_STATUS_ARM(generators)
	imageStorageProfile_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ImageStorageProfile_STATUS_ARM{}), generators)

	return imageStorageProfile_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForImageStorageProfile_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageStorageProfile_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ZoneResilient"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForImageStorageProfile_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageStorageProfile_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DataDisks"] = gen.SliceOf(ImageDataDisk_STATUS_ARMGenerator())
	gens["OsDisk"] = gen.PtrOf(ImageOSDisk_STATUS_ARMGenerator())
}

func Test_SubResource_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource_STATUS_ARM, SubResource_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource_STATUS_ARM runs a test to see if a specific instance of SubResource_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource_STATUS_ARM(subject SubResource_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource_STATUS_ARM
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

// Generator of SubResource_STATUS_ARM instances for property testing - lazily instantiated by
// SubResource_STATUS_ARMGenerator()
var subResource_STATUS_ARMGenerator gopter.Gen

// SubResource_STATUS_ARMGenerator returns a generator of SubResource_STATUS_ARM instances for property testing.
func SubResource_STATUS_ARMGenerator() gopter.Gen {
	if subResource_STATUS_ARMGenerator != nil {
		return subResource_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResource_STATUS_ARM(generators)
	subResource_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SubResource_STATUS_ARM{}), generators)

	return subResource_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSubResource_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResource_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_ImageDataDisk_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageDataDisk_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageDataDisk_STATUS_ARM, ImageDataDisk_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageDataDisk_STATUS_ARM runs a test to see if a specific instance of ImageDataDisk_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageDataDisk_STATUS_ARM(subject ImageDataDisk_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageDataDisk_STATUS_ARM
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

// Generator of ImageDataDisk_STATUS_ARM instances for property testing - lazily instantiated by
// ImageDataDisk_STATUS_ARMGenerator()
var imageDataDisk_STATUS_ARMGenerator gopter.Gen

// ImageDataDisk_STATUS_ARMGenerator returns a generator of ImageDataDisk_STATUS_ARM instances for property testing.
// We first initialize imageDataDisk_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageDataDisk_STATUS_ARMGenerator() gopter.Gen {
	if imageDataDisk_STATUS_ARMGenerator != nil {
		return imageDataDisk_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDataDisk_STATUS_ARM(generators)
	imageDataDisk_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ImageDataDisk_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDataDisk_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForImageDataDisk_STATUS_ARM(generators)
	imageDataDisk_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ImageDataDisk_STATUS_ARM{}), generators)

	return imageDataDisk_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForImageDataDisk_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageDataDisk_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["BlobUri"] = gen.PtrOf(gen.AlphaString())
	gens["Caching"] = gen.PtrOf(gen.OneConstOf(ImageDataDisk_Caching_STATUS_None, ImageDataDisk_Caching_STATUS_ReadOnly, ImageDataDisk_Caching_STATUS_ReadWrite))
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["Lun"] = gen.PtrOf(gen.Int())
	gens["StorageAccountType"] = gen.PtrOf(gen.OneConstOf(
		StorageAccountType_STATUS_Premium_LRS,
		StorageAccountType_STATUS_Premium_ZRS,
		StorageAccountType_STATUS_StandardSSD_LRS,
		StorageAccountType_STATUS_StandardSSD_ZRS,
		StorageAccountType_STATUS_Standard_LRS,
		StorageAccountType_STATUS_UltraSSD_LRS))
}

// AddRelatedPropertyGeneratorsForImageDataDisk_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageDataDisk_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSet"] = gen.PtrOf(SubResource_STATUS_ARMGenerator())
	gens["ManagedDisk"] = gen.PtrOf(SubResource_STATUS_ARMGenerator())
	gens["Snapshot"] = gen.PtrOf(SubResource_STATUS_ARMGenerator())
}

func Test_ImageOSDisk_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageOSDisk_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageOSDisk_STATUS_ARM, ImageOSDisk_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageOSDisk_STATUS_ARM runs a test to see if a specific instance of ImageOSDisk_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageOSDisk_STATUS_ARM(subject ImageOSDisk_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageOSDisk_STATUS_ARM
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

// Generator of ImageOSDisk_STATUS_ARM instances for property testing - lazily instantiated by
// ImageOSDisk_STATUS_ARMGenerator()
var imageOSDisk_STATUS_ARMGenerator gopter.Gen

// ImageOSDisk_STATUS_ARMGenerator returns a generator of ImageOSDisk_STATUS_ARM instances for property testing.
// We first initialize imageOSDisk_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageOSDisk_STATUS_ARMGenerator() gopter.Gen {
	if imageOSDisk_STATUS_ARMGenerator != nil {
		return imageOSDisk_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageOSDisk_STATUS_ARM(generators)
	imageOSDisk_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ImageOSDisk_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageOSDisk_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForImageOSDisk_STATUS_ARM(generators)
	imageOSDisk_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ImageOSDisk_STATUS_ARM{}), generators)

	return imageOSDisk_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForImageOSDisk_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageOSDisk_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["BlobUri"] = gen.PtrOf(gen.AlphaString())
	gens["Caching"] = gen.PtrOf(gen.OneConstOf(ImageOSDisk_Caching_STATUS_None, ImageOSDisk_Caching_STATUS_ReadOnly, ImageOSDisk_Caching_STATUS_ReadWrite))
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsState"] = gen.PtrOf(gen.OneConstOf(ImageOSDisk_OsState_STATUS_Generalized, ImageOSDisk_OsState_STATUS_Specialized))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(ImageOSDisk_OsType_STATUS_Linux, ImageOSDisk_OsType_STATUS_Windows))
	gens["StorageAccountType"] = gen.PtrOf(gen.OneConstOf(
		StorageAccountType_STATUS_Premium_LRS,
		StorageAccountType_STATUS_Premium_ZRS,
		StorageAccountType_STATUS_StandardSSD_LRS,
		StorageAccountType_STATUS_StandardSSD_ZRS,
		StorageAccountType_STATUS_Standard_LRS,
		StorageAccountType_STATUS_UltraSSD_LRS))
}

// AddRelatedPropertyGeneratorsForImageOSDisk_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageOSDisk_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSet"] = gen.PtrOf(SubResource_STATUS_ARMGenerator())
	gens["ManagedDisk"] = gen.PtrOf(SubResource_STATUS_ARMGenerator())
	gens["Snapshot"] = gen.PtrOf(SubResource_STATUS_ARMGenerator())
}