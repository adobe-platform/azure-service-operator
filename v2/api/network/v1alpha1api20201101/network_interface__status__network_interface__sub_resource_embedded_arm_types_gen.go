// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

//Deprecated version of NetworkInterface_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded instead
type NetworkInterface_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	Etag             *string                                     `json:"etag,omitempty"`
	ExtendedLocation *ExtendedLocation_StatusARM                 `json:"extendedLocation,omitempty"`
	Id               *string                                     `json:"id,omitempty"`
	Location         *string                                     `json:"location,omitempty"`
	Name             *string                                     `json:"name,omitempty"`
	Properties       *NetworkInterfacePropertiesFormat_StatusARM `json:"properties,omitempty"`
	Tags             map[string]string                           `json:"tags,omitempty"`
	Type             *string                                     `json:"type,omitempty"`
}

//Deprecated version of NetworkInterfacePropertiesFormat_Status. Use v1beta20201101.NetworkInterfacePropertiesFormat_Status instead
type NetworkInterfacePropertiesFormat_StatusARM struct {
	DnsSettings                 *NetworkInterfaceDnsSettings_StatusARM                                            `json:"dnsSettings,omitempty"`
	DscpConfiguration           *SubResource_StatusARM                                                            `json:"dscpConfiguration,omitempty"`
	EnableAcceleratedNetworking *bool                                                                             `json:"enableAcceleratedNetworking,omitempty"`
	EnableIPForwarding          *bool                                                                             `json:"enableIPForwarding,omitempty"`
	HostedWorkloads             []string                                                                          `json:"hostedWorkloads,omitempty"`
	IpConfigurations            []NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbeddedARM  `json:"ipConfigurations,omitempty"`
	MacAddress                  *string                                                                           `json:"macAddress,omitempty"`
	MigrationPhase              *NetworkInterfacePropertiesFormatStatusMigrationPhase                             `json:"migrationPhase,omitempty"`
	NetworkSecurityGroup        *NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbeddedARM              `json:"networkSecurityGroup,omitempty"`
	NicType                     *NetworkInterfacePropertiesFormatStatusNicType                                    `json:"nicType,omitempty"`
	Primary                     *bool                                                                             `json:"primary,omitempty"`
	PrivateEndpoint             *PrivateEndpoint_Status_NetworkInterface_SubResourceEmbeddedARM                   `json:"privateEndpoint,omitempty"`
	PrivateLinkService          *PrivateLinkService_Status_NetworkInterface_SubResourceEmbeddedARM                `json:"privateLinkService,omitempty"`
	ProvisioningState           *ProvisioningState_Status                                                         `json:"provisioningState,omitempty"`
	ResourceGuid                *string                                                                           `json:"resourceGuid,omitempty"`
	TapConfigurations           []NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbeddedARM `json:"tapConfigurations,omitempty"`
	VirtualMachine              *SubResource_StatusARM                                                            `json:"virtualMachine,omitempty"`
}

//Deprecated version of NetworkInterfaceDnsSettings_Status. Use v1beta20201101.NetworkInterfaceDnsSettings_Status instead
type NetworkInterfaceDnsSettings_StatusARM struct {
	AppliedDnsServers        []string `json:"appliedDnsServers,omitempty"`
	DnsServers               []string `json:"dnsServers,omitempty"`
	InternalDnsNameLabel     *string  `json:"internalDnsNameLabel,omitempty"`
	InternalDomainNameSuffix *string  `json:"internalDomainNameSuffix,omitempty"`
	InternalFqdn             *string  `json:"internalFqdn,omitempty"`
}

//Deprecated version of NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded instead
type NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	Etag       *string                                                                                         `json:"etag,omitempty"`
	Id         *string                                                                                         `json:"id,omitempty"`
	Name       *string                                                                                         `json:"name,omitempty"`
	Properties *NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM `json:"properties,omitempty"`
	Type       *string                                                                                         `json:"type,omitempty"`
}

//Deprecated version of NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbedded instead
type NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

//Deprecated version of NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbedded instead
type NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

//Deprecated version of PrivateEndpoint_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.PrivateEndpoint_Status_NetworkInterface_SubResourceEmbedded instead
type PrivateEndpoint_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	ExtendedLocation *ExtendedLocation_StatusARM `json:"extendedLocation,omitempty"`
	Id               *string                     `json:"id,omitempty"`
}

//Deprecated version of PrivateLinkService_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.PrivateLinkService_Status_NetworkInterface_SubResourceEmbedded instead
type PrivateLinkService_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	ExtendedLocation *ExtendedLocation_StatusARM `json:"extendedLocation,omitempty"`
	Id               *string                     `json:"id,omitempty"`
}

//Deprecated version of SubResource_Status. Use v1beta20201101.SubResource_Status instead
type SubResource_StatusARM struct {
	Id *string `json:"id,omitempty"`
}

//Deprecated version of NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbedded instead
type NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	ApplicationGatewayBackendAddressPools []ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbeddedARM `json:"applicationGatewayBackendAddressPools,omitempty"`
	ApplicationSecurityGroups             []ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbeddedARM             `json:"applicationSecurityGroups,omitempty"`
	LoadBalancerBackendAddressPools       []BackendAddressPool_Status_NetworkInterface_SubResourceEmbeddedARM                   `json:"loadBalancerBackendAddressPools,omitempty"`
	LoadBalancerInboundNatRules           []InboundNatRule_Status_NetworkInterface_SubResourceEmbeddedARM                       `json:"loadBalancerInboundNatRules,omitempty"`
	Primary                               *bool                                                                                 `json:"primary,omitempty"`
	PrivateIPAddress                      *string                                                                               `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion               *IPVersion_Status                                                                     `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod             *IPAllocationMethod_Status                                                            `json:"privateIPAllocationMethod,omitempty"`
	PrivateLinkConnectionProperties       *NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_StatusARM             `json:"privateLinkConnectionProperties,omitempty"`
	ProvisioningState                     *ProvisioningState_Status                                                             `json:"provisioningState,omitempty"`
	PublicIPAddress                       *PublicIPAddress_Status_NetworkInterface_SubResourceEmbeddedARM                       `json:"publicIPAddress,omitempty"`
	Subnet                                *Subnet_Status_NetworkInterface_SubResourceEmbeddedARM                                `json:"subnet,omitempty"`
	VirtualNetworkTaps                    []VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbeddedARM                    `json:"virtualNetworkTaps,omitempty"`
}

//Deprecated version of ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbedded instead
type ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	Etag       *string                                                                                              `json:"etag,omitempty"`
	Id         *string                                                                                              `json:"id,omitempty"`
	Name       *string                                                                                              `json:"name,omitempty"`
	Properties *ApplicationGatewayBackendAddressPoolPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM `json:"properties,omitempty"`
	Type       *string                                                                                              `json:"type,omitempty"`
}

//Deprecated version of ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbedded instead
type ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

//Deprecated version of BackendAddressPool_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.BackendAddressPool_Status_NetworkInterface_SubResourceEmbedded instead
type BackendAddressPool_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

//Deprecated version of InboundNatRule_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.InboundNatRule_Status_NetworkInterface_SubResourceEmbedded instead
type InboundNatRule_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

//Deprecated version of NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_Status. Use v1beta20201101.NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_Status instead
type NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_StatusARM struct {
	Fqdns              []string `json:"fqdns,omitempty"`
	GroupId            *string  `json:"groupId,omitempty"`
	RequiredMemberName *string  `json:"requiredMemberName,omitempty"`
}

//Deprecated version of PublicIPAddress_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.PublicIPAddress_Status_NetworkInterface_SubResourceEmbedded instead
type PublicIPAddress_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	ExtendedLocation *ExtendedLocation_StatusARM   `json:"extendedLocation,omitempty"`
	Id               *string                       `json:"id,omitempty"`
	Sku              *PublicIPAddressSku_StatusARM `json:"sku,omitempty"`
	Zones            []string                      `json:"zones,omitempty"`
}

//Deprecated version of Subnet_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.Subnet_Status_NetworkInterface_SubResourceEmbedded instead
type Subnet_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

//Deprecated version of VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbedded instead
type VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

//Deprecated version of ApplicationGatewayBackendAddressPoolPropertiesFormat_Status_NetworkInterface_SubResourceEmbedded. Use v1beta20201101.ApplicationGatewayBackendAddressPoolPropertiesFormat_Status_NetworkInterface_SubResourceEmbedded instead
type ApplicationGatewayBackendAddressPoolPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM struct {
	BackendAddresses  []ApplicationGatewayBackendAddress_StatusARM `json:"backendAddresses,omitempty"`
	ProvisioningState *ProvisioningState_Status                    `json:"provisioningState,omitempty"`
}

//Deprecated version of ApplicationGatewayBackendAddress_Status. Use v1beta20201101.ApplicationGatewayBackendAddress_Status instead
type ApplicationGatewayBackendAddress_StatusARM struct {
	Fqdn      *string `json:"fqdn,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
}