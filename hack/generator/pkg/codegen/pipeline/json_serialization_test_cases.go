/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/testcases"

	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// InjectJsonSerializationTestsID is the unique identifier for this pipeline stage
const InjectJsonSerializationTestsID = "injectJSONTestCases"

func InjectJsonSerializationTests(idFactory astmodel.IdentifierFactory) Stage {

	return MakeStage(
		InjectJsonSerializationTestsID,
		"Add test cases to verify JSON serialization",
		func(ctx context.Context, state *State) (*State, error) {
			factory := makeObjectSerializationTestCaseFactory(idFactory)
			modifiedTypes := make(astmodel.Types)
			var errs []error
			for _, def := range state.Types() {
				if factory.NeedsTest(def) {
					updated, err := factory.AddTestTo(def)
					if err != nil {
						errs = append(errs, err)
					} else {
						modifiedTypes[updated.Name()] = updated
					}
				}
			}

			if len(errs) > 0 {
				return nil, kerrors.NewAggregate(errs)
			}

			return state.WithTypes(state.Types().OverlayWith(modifiedTypes)), nil
		})
}

type objectSerializationTestCaseFactory struct {
	injector     *astmodel.TestCaseInjector
	idFactory    astmodel.IdentifierFactory
	suppressions []string
}

func makeObjectSerializationTestCaseFactory(idFactory astmodel.IdentifierFactory) objectSerializationTestCaseFactory {
	result := objectSerializationTestCaseFactory{
		injector:  astmodel.NewTestCaseInjector(),
		idFactory: idFactory,
		suppressions: []string{
			"DatabaseAccounts_SpecARM",
			"DatabaseAccountCreateUpdatePropertiesARM",
			"BackupPolicyARM",
		},
	}

	return result
}

// NeedsTest returns true if we should generate a testcase for the specified definition
func (s *objectSerializationTestCaseFactory) NeedsTest(def astmodel.TypeDefinition) bool {
	_, ok := astmodel.AsPropertyContainer(def.Type())
	if !ok {
		// Can only generate tests for property containers
		return false
	}

	// Check for types that we need to suppress - these are ARM types that don't currently round trip because they're
	// OneOf implementations that are only used in one direction.
	//
	// See https://github.com/Azure/azure-service-operator/issues/1721 for more information
	//
	result := true
	for _, s := range s.suppressions {
		if def.Name().Name() == s {
			result = false
			break
		}
	}

	return result
}

func (s *objectSerializationTestCaseFactory) AddTestTo(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	container, ok := astmodel.AsPropertyContainer(def.Type())
	if !ok {
		return astmodel.TypeDefinition{}, errors.Errorf("expected %s to be a property container", def.Name())
	}

	testcase := testcases.NewJSONSerializationTestCase(def.Name(), container, s.idFactory)
	return s.injector.Inject(def, testcase)
}
