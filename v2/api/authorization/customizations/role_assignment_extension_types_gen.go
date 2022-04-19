// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	authorizationv1alpha1api20200801preview "github.com/Azure/azure-service-operator/v2/api/authorization/v1alpha1api20200801preview"
	"github.com/Azure/azure-service-operator/v2/api/authorization/v1alpha1api20200801previewstorage"
	authorizationv1beta20200801preview "github.com/Azure/azure-service-operator/v2/api/authorization/v1beta20200801preview"
	"github.com/Azure/azure-service-operator/v2/api/authorization/v1beta20200801previewstorage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type RoleAssignmentExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *RoleAssignmentExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&authorizationv1alpha1api20200801preview.RoleAssignment{},
		&v1alpha1api20200801previewstorage.RoleAssignment{},
		&authorizationv1beta20200801preview.RoleAssignment{},
		&v1beta20200801previewstorage.RoleAssignment{}}
}