// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

// Deprecated version of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded. Use v1beta20201101.NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded instead
type NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM struct {
	Etag       *string                                          `json:"etag,omitempty"`
	Id         *string                                          `json:"id,omitempty"`
	Location   *string                                          `json:"location,omitempty"`
	Name       *string                                          `json:"name,omitempty"`
	Properties *NetworkSecurityGroupPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
	Tags       map[string]string                                `json:"tags,omitempty"`
	Type       *string                                          `json:"type,omitempty"`
}

// Deprecated version of NetworkSecurityGroupPropertiesFormat_STATUS. Use v1beta20201101.NetworkSecurityGroupPropertiesFormat_STATUS instead
type NetworkSecurityGroupPropertiesFormat_STATUS_ARM struct {
	DefaultSecurityRules []SecurityRule_STATUS_ARM                                              `json:"defaultSecurityRules,omitempty"`
	FlowLogs             []FlowLog_STATUS_ARM                                                   `json:"flowLogs,omitempty"`
	NetworkInterfaces    []NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM `json:"networkInterfaces,omitempty"`
	ProvisioningState    *ProvisioningState_STATUS                                              `json:"provisioningState,omitempty"`
	ResourceGuid         *string                                                                `json:"resourceGuid,omitempty"`
	SecurityRules        []SecurityRule_STATUS_ARM                                              `json:"securityRules,omitempty"`
	Subnets              []Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM           `json:"subnets,omitempty"`
}

// Deprecated version of FlowLog_STATUS. Use v1beta20201101.FlowLog_STATUS instead
type FlowLog_STATUS_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded. Use v1beta20201101.NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded instead
type NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of SecurityRule_STATUS. Use v1beta20201101.SecurityRule_STATUS instead
type SecurityRule_STATUS_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded. Use v1beta20201101.Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded instead
type Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}