// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220301

// A web app, a mobile app backend, or an API app.
type Site_STATUS_ARM struct {
	// ExtendedLocation: Extended Location.
	ExtendedLocation *ExtendedLocation_STATUS_ARM `json:"extendedLocation,omitempty"`

	// Id: Resource Id.
	Id *string `json:"id,omitempty"`

	// Identity: Managed service identity.
	Identity *ManagedServiceIdentity_STATUS_ARM `json:"identity,omitempty"`

	// Kind: Kind of resource.
	Kind *string `json:"kind,omitempty"`

	// Location: Resource Location.
	Location *string `json:"location,omitempty"`

	// Name: Resource Name.
	Name *string `json:"name,omitempty"`

	// Properties: Site resource specific properties
	Properties *Site_Properties_STATUS_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Managed service identity.
type ManagedServiceIdentity_STATUS_ARM struct {
	// PrincipalId: Principal Id of managed service identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: Tenant of managed service identity.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: Type of managed service identity.
	Type *ManagedServiceIdentity_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: The list of user assigned identities associated with the resource. The user identity dictionary
	// key references will be ARM resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}
	UserAssignedIdentities map[string]UserAssignedIdentity_STATUS_ARM `json:"userAssignedIdentities,omitempty"`
}

type Site_Properties_STATUS_ARM struct {
	// AvailabilityState: Management information availability state for the app.
	AvailabilityState *Site_Properties_AvailabilityState_STATUS `json:"availabilityState,omitempty"`

	// ClientAffinityEnabled: <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity
	// cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
	ClientAffinityEnabled *bool `json:"clientAffinityEnabled,omitempty"`

	// ClientCertEnabled: <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise,
	// <code>false</code>. Default is <code>false</code>.
	ClientCertEnabled *bool `json:"clientCertEnabled,omitempty"`

	// ClientCertExclusionPaths: client certificate authentication comma-separated exclusion paths
	ClientCertExclusionPaths *string `json:"clientCertExclusionPaths,omitempty"`

	// ClientCertMode: This composes with ClientCertEnabled setting.
	// - ClientCertEnabled: false means ClientCert is ignored.
	// - ClientCertEnabled: true and ClientCertMode: Required means ClientCert is required.
	// - ClientCertEnabled: true and ClientCertMode: Optional means ClientCert is optional or accepted.
	ClientCertMode *Site_Properties_ClientCertMode_STATUS `json:"clientCertMode,omitempty"`

	// CloningInfo: If specified during app creation, the app is cloned from a source app.
	CloningInfo *CloningInfo_STATUS_ARM `json:"cloningInfo,omitempty"`

	// ContainerSize: Size of the function container.
	ContainerSize *int `json:"containerSize,omitempty"`

	// CustomDomainVerificationId: Unique identifier that verifies the custom domains assigned to the app. Customer will add
	// this id to a txt record for verification.
	CustomDomainVerificationId *string `json:"customDomainVerificationId,omitempty"`

	// DailyMemoryTimeQuota: Maximum allowed daily memory-time quota (applicable on dynamic apps only).
	DailyMemoryTimeQuota *int `json:"dailyMemoryTimeQuota,omitempty"`

	// DefaultHostName: Default hostname of the app. Read-only.
	DefaultHostName *string `json:"defaultHostName,omitempty"`

	// Enabled: <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables
	// the app (takes the app offline).
	Enabled *bool `json:"enabled,omitempty"`

	// EnabledHostNames: Enabled hostnames for the app.Hostnames need to be assigned (see HostNames) AND enabled. Otherwise,
	// the app is not served on those hostnames.
	EnabledHostNames []string `json:"enabledHostNames,omitempty"`

	// HostNameSslStates: Hostname SSL states are used to manage the SSL bindings for app's hostnames.
	HostNameSslStates []HostNameSslState_STATUS_ARM `json:"hostNameSslStates,omitempty"`

	// HostNames: Hostnames associated with the app.
	HostNames []string `json:"hostNames,omitempty"`

	// HostNamesDisabled: <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
	// If <code>true</code>, the app is only accessible via API management process.
	HostNamesDisabled *bool `json:"hostNamesDisabled,omitempty"`

	// HostingEnvironmentProfile: App Service Environment to use for the app.
	HostingEnvironmentProfile *HostingEnvironmentProfile_STATUS_ARM `json:"hostingEnvironmentProfile,omitempty"`

	// HttpsOnly: HttpsOnly: configures a web site to accept only https requests. Issues redirect for
	// http requests
	HttpsOnly *bool `json:"httpsOnly,omitempty"`

	// HyperV: Hyper-V sandbox.
	HyperV *bool `json:"hyperV,omitempty"`

	// InProgressOperationId: Specifies an operation id if this site has a pending operation.
	InProgressOperationId *string `json:"inProgressOperationId,omitempty"`

	// IsDefaultContainer: <code>true</code> if the app is a default container; otherwise, <code>false</code>.
	IsDefaultContainer *bool `json:"isDefaultContainer,omitempty"`

	// IsXenon: Obsolete: Hyper-V sandbox.
	IsXenon *bool `json:"isXenon,omitempty"`

	// KeyVaultReferenceIdentity: Identity to use for Key Vault Reference authentication.
	KeyVaultReferenceIdentity *string `json:"keyVaultReferenceIdentity,omitempty"`

	// LastModifiedTimeUtc: Last time the app was modified, in UTC. Read-only.
	LastModifiedTimeUtc *string `json:"lastModifiedTimeUtc,omitempty"`

	// MaxNumberOfWorkers: Maximum number of workers.
	// This only applies to Functions container.
	MaxNumberOfWorkers *int `json:"maxNumberOfWorkers,omitempty"`

	// OutboundIpAddresses: List of IP addresses that the app uses for outbound connections (e.g. database access). Includes
	// VIPs from tenants that site can be hosted with current settings. Read-only.
	OutboundIpAddresses *string `json:"outboundIpAddresses,omitempty"`

	// PossibleOutboundIpAddresses: List of IP addresses that the app uses for outbound connections (e.g. database access).
	// Includes VIPs from all tenants except dataComponent. Read-only.
	PossibleOutboundIpAddresses *string `json:"possibleOutboundIpAddresses,omitempty"`

	// PublicNetworkAccess: Property to allow or block all public traffic. Allowed Values: 'Enabled', 'Disabled' or an empty
	// string.
	PublicNetworkAccess *string `json:"publicNetworkAccess,omitempty"`

	// RedundancyMode: Site redundancy mode
	RedundancyMode *Site_Properties_RedundancyMode_STATUS `json:"redundancyMode,omitempty"`

	// RepositorySiteName: Name of the repository site.
	RepositorySiteName *string `json:"repositorySiteName,omitempty"`

	// Reserved: <code>true</code> if reserved; otherwise, <code>false</code>.
	Reserved *bool `json:"reserved,omitempty"`

	// ResourceGroup: Name of the resource group the app belongs to. Read-only.
	ResourceGroup *string `json:"resourceGroup,omitempty"`

	// ScmSiteAlsoStopped: <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>.
	// The default is <code>false</code>.
	ScmSiteAlsoStopped *bool `json:"scmSiteAlsoStopped,omitempty"`

	// ServerFarmId: Resource ID of the associated App Service plan, formatted as:
	// "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
	ServerFarmId *string `json:"serverFarmId,omitempty"`

	// SiteConfig: Configuration of the app.
	SiteConfig *SiteConfig_STATUS_ARM `json:"siteConfig,omitempty"`

	// SlotSwapStatus: Status of the last deployment slot swap operation.
	SlotSwapStatus *SlotSwapStatus_STATUS_ARM `json:"slotSwapStatus,omitempty"`

	// State: Current state of the app.
	State *string `json:"state,omitempty"`

	// StorageAccountRequired: Checks if Customer provided storage account is required
	StorageAccountRequired *bool `json:"storageAccountRequired,omitempty"`

	// SuspendedTill: App suspended till in case memory-time quota is exceeded.
	SuspendedTill *string `json:"suspendedTill,omitempty"`

	// TargetSwapSlot: Specifies which deployment slot this app will swap into. Read-only.
	TargetSwapSlot *string `json:"targetSwapSlot,omitempty"`

	// TrafficManagerHostNames: Azure Traffic Manager hostnames associated with the app. Read-only.
	TrafficManagerHostNames []string `json:"trafficManagerHostNames,omitempty"`

	// UsageState: State indicating whether the app has exceeded its quota usage. Read-only.
	UsageState *Site_Properties_UsageState_STATUS `json:"usageState,omitempty"`

	// VirtualNetworkSubnetId: Azure Resource Manager ID of the Virtual network and subnet to be joined by Regional VNET
	// Integration.
	// This must be of the form
	// /subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}
	VirtualNetworkSubnetId *string `json:"virtualNetworkSubnetId,omitempty"`

	// VnetContentShareEnabled: To enable accessing content over virtual network
	VnetContentShareEnabled *bool `json:"vnetContentShareEnabled,omitempty"`

	// VnetImagePullEnabled: To enable pulling image over Virtual Network
	VnetImagePullEnabled *bool `json:"vnetImagePullEnabled,omitempty"`

	// VnetRouteAllEnabled: Virtual Network Route All enabled. This causes all outbound traffic to have Virtual Network
	// Security Groups and User Defined Routes applied.
	VnetRouteAllEnabled *bool `json:"vnetRouteAllEnabled,omitempty"`
}

// Information needed for cloning operation.
type CloningInfo_STATUS_ARM struct {
	// AppSettingsOverrides: Application setting overrides for cloned app. If specified, these settings override the settings
	// cloned
	// from source app. Otherwise, application settings from source app are retained.
	AppSettingsOverrides map[string]string `json:"appSettingsOverrides,omitempty"`

	// CloneCustomHostNames: <code>true</code> to clone custom hostnames from source app; otherwise, <code>false</code>.
	CloneCustomHostNames *bool `json:"cloneCustomHostNames,omitempty"`

	// CloneSourceControl: <code>true</code> to clone source control from source app; otherwise, <code>false</code>.
	CloneSourceControl *bool `json:"cloneSourceControl,omitempty"`

	// ConfigureLoadBalancing: <code>true</code> to configure load balancing for source and destination app.
	ConfigureLoadBalancing *bool `json:"configureLoadBalancing,omitempty"`

	// CorrelationId: Correlation ID of cloning operation. This ID ties multiple cloning operations
	// together to use the same snapshot.
	CorrelationId *string `json:"correlationId,omitempty"`

	// HostingEnvironment: App Service Environment.
	HostingEnvironment *string `json:"hostingEnvironment,omitempty"`

	// Overwrite: <code>true</code> to overwrite destination app; otherwise, <code>false</code>.
	Overwrite *bool `json:"overwrite,omitempty"`

	// SourceWebAppId: ARM resource ID of the source app. App resource ID is of the form
	// /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName} for production slots
	// and
	// /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slotName} for
	// other slots.
	SourceWebAppId *string `json:"sourceWebAppId,omitempty"`

	// SourceWebAppLocation: Location of source app ex: West US or North Europe
	SourceWebAppLocation *string `json:"sourceWebAppLocation,omitempty"`

	// TrafficManagerProfileId: ARM resource ID of the Traffic Manager profile to use, if it exists. Traffic Manager resource
	// ID is of the form
	// /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{profileName}.
	TrafficManagerProfileId *string `json:"trafficManagerProfileId,omitempty"`

	// TrafficManagerProfileName: Name of Traffic Manager profile to create. This is only needed if Traffic Manager profile
	// does not already exist.
	TrafficManagerProfileName *string `json:"trafficManagerProfileName,omitempty"`
}

// SSL-enabled hostname.
type HostNameSslState_STATUS_ARM struct {
	// HostType: Indicates whether the hostname is a standard or repository hostname.
	HostType *HostNameSslState_HostType_STATUS `json:"hostType,omitempty"`

	// Name: Hostname.
	Name *string `json:"name,omitempty"`

	// SslState: SSL type.
	SslState *HostNameSslState_SslState_STATUS `json:"sslState,omitempty"`

	// Thumbprint: SSL certificate thumbprint.
	Thumbprint *string `json:"thumbprint,omitempty"`

	// ToUpdate: Set to <code>true</code> to update existing hostname.
	ToUpdate *bool `json:"toUpdate,omitempty"`

	// VirtualIP: Virtual IP address assigned to the hostname if IP based SSL is enabled.
	VirtualIP *string `json:"virtualIP,omitempty"`
}

type ManagedServiceIdentity_Type_STATUS string

const (
	ManagedServiceIdentity_Type_STATUS_None                       = ManagedServiceIdentity_Type_STATUS("None")
	ManagedServiceIdentity_Type_STATUS_SystemAssigned             = ManagedServiceIdentity_Type_STATUS("SystemAssigned")
	ManagedServiceIdentity_Type_STATUS_SystemAssignedUserAssigned = ManagedServiceIdentity_Type_STATUS("SystemAssigned, UserAssigned")
	ManagedServiceIdentity_Type_STATUS_UserAssigned               = ManagedServiceIdentity_Type_STATUS("UserAssigned")
)

// Configuration of an App Service app.
type SiteConfig_STATUS_ARM struct {
	// AcrUseManagedIdentityCreds: Flag to use Managed Identity Creds for ACR pull
	AcrUseManagedIdentityCreds *bool `json:"acrUseManagedIdentityCreds,omitempty"`

	// AcrUserManagedIdentityID: If using user managed identity, the user managed identity ClientId
	AcrUserManagedIdentityID *string `json:"acrUserManagedIdentityID,omitempty"`

	// AlwaysOn: <code>true</code> if Always On is enabled; otherwise, <code>false</code>.
	AlwaysOn *bool `json:"alwaysOn,omitempty"`

	// ApiDefinition: Information about the formal API definition for the app.
	ApiDefinition *ApiDefinitionInfo_STATUS_ARM `json:"apiDefinition,omitempty"`

	// ApiManagementConfig: Azure API management settings linked to the app.
	ApiManagementConfig *ApiManagementConfig_STATUS_ARM `json:"apiManagementConfig,omitempty"`

	// AppCommandLine: App command line to launch.
	AppCommandLine *string `json:"appCommandLine,omitempty"`

	// AppSettings: Application settings.
	AppSettings []NameValuePair_STATUS_ARM `json:"appSettings,omitempty"`

	// AutoHealEnabled: <code>true</code> if Auto Heal is enabled; otherwise, <code>false</code>.
	AutoHealEnabled *bool `json:"autoHealEnabled,omitempty"`

	// AutoHealRules: Auto Heal rules.
	AutoHealRules *AutoHealRules_STATUS_ARM `json:"autoHealRules,omitempty"`

	// AutoSwapSlotName: Auto-swap slot name.
	AutoSwapSlotName *string `json:"autoSwapSlotName,omitempty"`

	// AzureStorageAccounts: List of Azure Storage Accounts.
	AzureStorageAccounts map[string]AzureStorageInfoValue_STATUS_ARM `json:"azureStorageAccounts,omitempty"`

	// ConnectionStrings: Connection strings.
	ConnectionStrings []ConnStringInfo_STATUS_ARM `json:"connectionStrings,omitempty"`

	// Cors: Cross-Origin Resource Sharing (CORS) settings.
	Cors *CorsSettings_STATUS_ARM `json:"cors,omitempty"`

	// DefaultDocuments: Default documents.
	DefaultDocuments []string `json:"defaultDocuments,omitempty"`

	// DetailedErrorLoggingEnabled: <code>true</code> if detailed error logging is enabled; otherwise, <code>false</code>.
	DetailedErrorLoggingEnabled *bool `json:"detailedErrorLoggingEnabled,omitempty"`

	// DocumentRoot: Document root.
	DocumentRoot *string `json:"documentRoot,omitempty"`

	// Experiments: This is work around for polymorphic types.
	Experiments *Experiments_STATUS_ARM `json:"experiments,omitempty"`

	// FtpsState: State of FTP / FTPS service
	FtpsState *SiteConfig_FtpsState_STATUS `json:"ftpsState,omitempty"`

	// FunctionAppScaleLimit: Maximum number of workers that a site can scale out to.
	// This setting only applies to the Consumption and Elastic Premium Plans
	FunctionAppScaleLimit *int `json:"functionAppScaleLimit,omitempty"`

	// FunctionsRuntimeScaleMonitoringEnabled: Gets or sets a value indicating whether functions runtime scale monitoring is
	// enabled. When enabled,
	// the ScaleController will not monitor event sources directly, but will instead call to the
	// runtime to get scale status.
	FunctionsRuntimeScaleMonitoringEnabled *bool `json:"functionsRuntimeScaleMonitoringEnabled,omitempty"`

	// HandlerMappings: Handler mappings.
	HandlerMappings []HandlerMapping_STATUS_ARM `json:"handlerMappings,omitempty"`

	// HealthCheckPath: Health check path
	HealthCheckPath *string `json:"healthCheckPath,omitempty"`

	// Http20Enabled: Http20Enabled: configures a web site to allow clients to connect over http2.0
	Http20Enabled *bool `json:"http20Enabled,omitempty"`

	// HttpLoggingEnabled: <code>true</code> if HTTP logging is enabled; otherwise, <code>false</code>.
	HttpLoggingEnabled *bool `json:"httpLoggingEnabled,omitempty"`

	// IpSecurityRestrictions: IP security restrictions for main.
	IpSecurityRestrictions []IpSecurityRestriction_STATUS_ARM `json:"ipSecurityRestrictions,omitempty"`

	// JavaContainer: Java container.
	JavaContainer *string `json:"javaContainer,omitempty"`

	// JavaContainerVersion: Java container version.
	JavaContainerVersion *string `json:"javaContainerVersion,omitempty"`

	// JavaVersion: Java version.
	JavaVersion *string `json:"javaVersion,omitempty"`

	// KeyVaultReferenceIdentity: Identity to use for Key Vault Reference authentication.
	KeyVaultReferenceIdentity *string `json:"keyVaultReferenceIdentity,omitempty"`

	// Limits: Site limits.
	Limits *SiteLimits_STATUS_ARM `json:"limits,omitempty"`

	// LinuxFxVersion: Linux App Framework and version
	LinuxFxVersion *string `json:"linuxFxVersion,omitempty"`

	// LoadBalancing: Site load balancing.
	LoadBalancing *SiteConfig_LoadBalancing_STATUS `json:"loadBalancing,omitempty"`

	// LocalMySqlEnabled: <code>true</code> to enable local MySQL; otherwise, <code>false</code>.
	LocalMySqlEnabled *bool `json:"localMySqlEnabled,omitempty"`

	// LogsDirectorySizeLimit: HTTP logs directory size limit.
	LogsDirectorySizeLimit *int `json:"logsDirectorySizeLimit,omitempty"`

	// MachineKey: Site MachineKey.
	MachineKey *SiteMachineKey_STATUS_ARM `json:"machineKey,omitempty"`

	// ManagedPipelineMode: Managed pipeline mode.
	ManagedPipelineMode *SiteConfig_ManagedPipelineMode_STATUS `json:"managedPipelineMode,omitempty"`

	// ManagedServiceIdentityId: Managed Service Identity Id
	ManagedServiceIdentityId *int `json:"managedServiceIdentityId,omitempty"`

	// MinTlsVersion: MinTlsVersion: configures the minimum version of TLS required for SSL requests
	MinTlsVersion *SiteConfig_MinTlsVersion_STATUS `json:"minTlsVersion,omitempty"`

	// MinimumElasticInstanceCount: Number of minimum instance count for a site
	// This setting only applies to the Elastic Plans
	MinimumElasticInstanceCount *int `json:"minimumElasticInstanceCount,omitempty"`

	// NetFrameworkVersion: .NET Framework version.
	NetFrameworkVersion *string `json:"netFrameworkVersion,omitempty"`

	// NodeVersion: Version of Node.js.
	NodeVersion *string `json:"nodeVersion,omitempty"`

	// NumberOfWorkers: Number of workers.
	NumberOfWorkers *int `json:"numberOfWorkers,omitempty"`

	// PhpVersion: Version of PHP.
	PhpVersion *string `json:"phpVersion,omitempty"`

	// PowerShellVersion: Version of PowerShell.
	PowerShellVersion *string `json:"powerShellVersion,omitempty"`

	// PreWarmedInstanceCount: Number of preWarmed instances.
	// This setting only applies to the Consumption and Elastic Plans
	PreWarmedInstanceCount *int `json:"preWarmedInstanceCount,omitempty"`

	// PublicNetworkAccess: Property to allow or block all public traffic.
	PublicNetworkAccess *string `json:"publicNetworkAccess,omitempty"`

	// PublishingUsername: Publishing user name.
	PublishingUsername *string `json:"publishingUsername,omitempty"`

	// Push: Push endpoint settings.
	Push *PushSettings_STATUS_ARM `json:"push,omitempty"`

	// PythonVersion: Version of Python.
	PythonVersion *string `json:"pythonVersion,omitempty"`

	// RemoteDebuggingEnabled: <code>true</code> if remote debugging is enabled; otherwise, <code>false</code>.
	RemoteDebuggingEnabled *bool `json:"remoteDebuggingEnabled,omitempty"`

	// RemoteDebuggingVersion: Remote debugging version.
	RemoteDebuggingVersion *string `json:"remoteDebuggingVersion,omitempty"`

	// RequestTracingEnabled: <code>true</code> if request tracing is enabled; otherwise, <code>false</code>.
	RequestTracingEnabled *bool `json:"requestTracingEnabled,omitempty"`

	// RequestTracingExpirationTime: Request tracing expiration time.
	RequestTracingExpirationTime *string `json:"requestTracingExpirationTime,omitempty"`

	// ScmIpSecurityRestrictions: IP security restrictions for scm.
	ScmIpSecurityRestrictions []IpSecurityRestriction_STATUS_ARM `json:"scmIpSecurityRestrictions,omitempty"`

	// ScmIpSecurityRestrictionsUseMain: IP security restrictions for scm to use main.
	ScmIpSecurityRestrictionsUseMain *bool `json:"scmIpSecurityRestrictionsUseMain,omitempty"`

	// ScmMinTlsVersion: ScmMinTlsVersion: configures the minimum version of TLS required for SSL requests for SCM site
	ScmMinTlsVersion *SiteConfig_ScmMinTlsVersion_STATUS `json:"scmMinTlsVersion,omitempty"`

	// ScmType: SCM type.
	ScmType *SiteConfig_ScmType_STATUS `json:"scmType,omitempty"`

	// TracingOptions: Tracing options.
	TracingOptions *string `json:"tracingOptions,omitempty"`

	// Use32BitWorkerProcess: <code>true</code> to use 32-bit worker process; otherwise, <code>false</code>.
	Use32BitWorkerProcess *bool `json:"use32BitWorkerProcess,omitempty"`

	// VirtualApplications: Virtual applications.
	VirtualApplications []VirtualApplication_STATUS_ARM `json:"virtualApplications,omitempty"`

	// VnetName: Virtual Network name.
	VnetName *string `json:"vnetName,omitempty"`

	// VnetPrivatePortsCount: The number of private ports assigned to this app. These will be assigned dynamically on runtime.
	VnetPrivatePortsCount *int `json:"vnetPrivatePortsCount,omitempty"`

	// VnetRouteAllEnabled: Virtual Network Route All enabled. This causes all outbound traffic to have Virtual Network
	// Security Groups and User Defined Routes applied.
	VnetRouteAllEnabled *bool `json:"vnetRouteAllEnabled,omitempty"`

	// WebSocketsEnabled: <code>true</code> if WebSocket is enabled; otherwise, <code>false</code>.
	WebSocketsEnabled *bool `json:"webSocketsEnabled,omitempty"`

	// WebsiteTimeZone: Sets the time zone a site uses for generating timestamps. Compatible with Linux and Windows App
	// Service. Setting the WEBSITE_TIME_ZONE app setting takes precedence over this config. For Linux, expects tz database
	// values https://www.iana.org/time-zones (for a quick reference see
	// https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). For Windows, expects one of the time zones listed under
	// HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones
	WebsiteTimeZone *string `json:"websiteTimeZone,omitempty"`

	// WindowsFxVersion: Xenon App Framework and version
	WindowsFxVersion *string `json:"windowsFxVersion,omitempty"`

	// XManagedServiceIdentityId: Explicit Managed Service Identity Id
	XManagedServiceIdentityId *int `json:"xManagedServiceIdentityId,omitempty"`
}

// The status of the last successful slot swap operation.
type SlotSwapStatus_STATUS_ARM struct {
	// DestinationSlotName: The destination slot of the last swap operation.
	DestinationSlotName *string `json:"destinationSlotName,omitempty"`

	// SourceSlotName: The source slot of the last swap operation.
	SourceSlotName *string `json:"sourceSlotName,omitempty"`

	// TimestampUtc: The time the last successful slot swap completed.
	TimestampUtc *string `json:"timestampUtc,omitempty"`
}

// User Assigned identity.
type UserAssignedIdentity_STATUS_ARM struct {
	// ClientId: Client Id of user assigned identity
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: Principal Id of user assigned identity
	PrincipalId *string `json:"principalId,omitempty"`
}

// Information about the formal API definition for the app.
type ApiDefinitionInfo_STATUS_ARM struct {
	// Url: The URL of the API definition.
	Url *string `json:"url,omitempty"`
}

// Azure API management (APIM) configuration linked to the app.
type ApiManagementConfig_STATUS_ARM struct {
	// Id: APIM-Api Identifier.
	Id *string `json:"id,omitempty"`
}

// Rules that can be defined for auto-heal.
type AutoHealRules_STATUS_ARM struct {
	// Actions: Actions to be executed when a rule is triggered.
	Actions *AutoHealActions_STATUS_ARM `json:"actions,omitempty"`

	// Triggers: Conditions that describe when to execute the auto-heal actions.
	Triggers *AutoHealTriggers_STATUS_ARM `json:"triggers,omitempty"`
}

// Azure Files or Blob Storage access information value for dictionary storage.
type AzureStorageInfoValue_STATUS_ARM struct {
	// AccountName: Name of the storage account.
	AccountName *string `json:"accountName,omitempty"`

	// MountPath: Path to mount the storage within the site's runtime environment.
	MountPath *string `json:"mountPath,omitempty"`

	// ShareName: Name of the file share (container name, for Blob storage).
	ShareName *string `json:"shareName,omitempty"`

	// State: State of the storage account.
	State *AzureStorageInfoValue_State_STATUS `json:"state,omitempty"`

	// Type: Type of storage.
	Type *AzureStorageInfoValue_Type_STATUS `json:"type,omitempty"`
}

// Database connection string information.
type ConnStringInfo_STATUS_ARM struct {
	// ConnectionString: Connection string value.
	ConnectionString *string `json:"connectionString,omitempty"`

	// Name: Name of connection string.
	Name *string `json:"name,omitempty"`

	// Type: Type of database.
	Type *ConnStringInfo_Type_STATUS `json:"type,omitempty"`
}

// Cross-Origin Resource Sharing (CORS) settings for the app.
type CorsSettings_STATUS_ARM struct {
	// AllowedOrigins: Gets or sets the list of origins that should be allowed to make cross-origin
	// calls (for example: http://example.com:12345). Use "*" to allow all.
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`

	// SupportCredentials: Gets or sets whether CORS requests with credentials are allowed. See
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS#Requests_with_credentials
	// for more details.
	SupportCredentials *bool `json:"supportCredentials,omitempty"`
}

// Routing rules in production experiments.
type Experiments_STATUS_ARM struct {
	// RampUpRules: List of ramp-up rules.
	RampUpRules []RampUpRule_STATUS_ARM `json:"rampUpRules,omitempty"`
}

// The IIS handler mappings used to define which handler processes HTTP requests with certain extension.
// For example, it
// is used to configure php-cgi.exe process to handle all HTTP requests with *.php extension.
type HandlerMapping_STATUS_ARM struct {
	// Arguments: Command-line arguments to be passed to the script processor.
	Arguments *string `json:"arguments,omitempty"`

	// Extension: Requests with this extension will be handled using the specified FastCGI application.
	Extension *string `json:"extension,omitempty"`

	// ScriptProcessor: The absolute path to the FastCGI application.
	ScriptProcessor *string `json:"scriptProcessor,omitempty"`
}

// IP security restriction on an app.
type IpSecurityRestriction_STATUS_ARM struct {
	// Action: Allow or Deny access for this IP range.
	Action *string `json:"action,omitempty"`

	// Description: IP restriction rule description.
	Description *string `json:"description,omitempty"`

	// Headers: IP restriction rule headers.
	// X-Forwarded-Host (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-Host#Examples).
	// The matching logic is ..
	// - If the property is null or empty (default), all hosts(or lack of) are allowed.
	// - A value is compared using ordinal-ignore-case (excluding port number).
	// - Subdomain wildcards are permitted but don't match the root domain. For example, *.contoso.com matches the subdomain
	// foo.contoso.com
	// but not the root domain contoso.com or multi-level foo.bar.contoso.com
	// - Unicode host names are allowed but are converted to Punycode for matching.
	// X-Forwarded-For (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For#Examples).
	// The matching logic is ..
	// - If the property is null or empty (default), any forwarded-for chains (or lack of) are allowed.
	// - If any address (excluding port number) in the chain (comma separated) matches the CIDR defined by the property.
	// X-Azure-FDID and X-FD-HealthProbe.
	// The matching logic is exact match.
	Headers map[string][]string `json:"headers,omitempty"`

	// IpAddress: IP address the security restriction is valid for.
	// It can be in form of pure ipv4 address (required SubnetMask property) or
	// CIDR notation such as ipv4/mask (leading bit match). For CIDR,
	// SubnetMask property must not be specified.
	IpAddress *string `json:"ipAddress,omitempty"`

	// Name: IP restriction rule name.
	Name *string `json:"name,omitempty"`

	// Priority: Priority of IP restriction rule.
	Priority *int `json:"priority,omitempty"`

	// SubnetMask: Subnet mask for the range of IP addresses the restriction is valid for.
	SubnetMask *string `json:"subnetMask,omitempty"`

	// SubnetTrafficTag: (internal) Subnet traffic tag
	SubnetTrafficTag *int `json:"subnetTrafficTag,omitempty"`

	// Tag: Defines what this IP filter will be used for. This is to support IP filtering on proxies.
	Tag *IpSecurityRestriction_Tag_STATUS `json:"tag,omitempty"`

	// VnetSubnetResourceId: Virtual network resource id
	VnetSubnetResourceId *string `json:"vnetSubnetResourceId,omitempty"`

	// VnetTrafficTag: (internal) Vnet traffic tag
	VnetTrafficTag *int `json:"vnetTrafficTag,omitempty"`
}

// Name value pair.
type NameValuePair_STATUS_ARM struct {
	// Name: Pair name.
	Name *string `json:"name,omitempty"`

	// Value: Pair value.
	Value *string `json:"value,omitempty"`
}

// Push settings for the App.
type PushSettings_STATUS_ARM struct {
	// Id: Resource Id.
	Id *string `json:"id,omitempty"`

	// Kind: Kind of resource.
	Kind *string `json:"kind,omitempty"`

	// Name: Resource Name.
	Name *string `json:"name,omitempty"`

	// Properties: PushSettings resource specific properties
	Properties *PushSettings_Properties_STATUS_ARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Metric limits set on an app.
type SiteLimits_STATUS_ARM struct {
	// MaxDiskSizeInMb: Maximum allowed disk size usage in MB.
	MaxDiskSizeInMb *int `json:"maxDiskSizeInMb,omitempty"`

	// MaxMemoryInMb: Maximum allowed memory usage in MB.
	MaxMemoryInMb *int `json:"maxMemoryInMb,omitempty"`

	// MaxPercentageCpu: Maximum allowed CPU usage percentage.
	MaxPercentageCpu *float64 `json:"maxPercentageCpu,omitempty"`
}

// MachineKey of an app.
type SiteMachineKey_STATUS_ARM struct {
	// Decryption: Algorithm used for decryption.
	Decryption *string `json:"decryption,omitempty"`

	// DecryptionKey: Decryption key.
	DecryptionKey *string `json:"decryptionKey,omitempty"`

	// Validation: MachineKey validation.
	Validation *string `json:"validation,omitempty"`

	// ValidationKey: Validation key.
	ValidationKey *string `json:"validationKey,omitempty"`
}

// Virtual application in an app.
type VirtualApplication_STATUS_ARM struct {
	// PhysicalPath: Physical path.
	PhysicalPath *string `json:"physicalPath,omitempty"`

	// PreloadEnabled: <code>true</code> if preloading is enabled; otherwise, <code>false</code>.
	PreloadEnabled *bool `json:"preloadEnabled,omitempty"`

	// VirtualDirectories: Virtual directories for virtual application.
	VirtualDirectories []VirtualDirectory_STATUS_ARM `json:"virtualDirectories,omitempty"`

	// VirtualPath: Virtual path.
	VirtualPath *string `json:"virtualPath,omitempty"`
}

// Actions which to take by the auto-heal module when a rule is triggered.
type AutoHealActions_STATUS_ARM struct {
	// ActionType: Predefined action to be taken.
	ActionType *AutoHealActions_ActionType_STATUS `json:"actionType,omitempty"`

	// CustomAction: Custom action to be taken.
	CustomAction *AutoHealCustomAction_STATUS_ARM `json:"customAction,omitempty"`

	// MinProcessExecutionTime: Minimum time the process must execute
	// before taking the action
	MinProcessExecutionTime *string `json:"minProcessExecutionTime,omitempty"`
}

// Triggers for auto-heal.
type AutoHealTriggers_STATUS_ARM struct {
	// PrivateBytesInKB: A rule based on private bytes.
	PrivateBytesInKB *int `json:"privateBytesInKB,omitempty"`

	// Requests: A rule based on total requests.
	Requests *RequestsBasedTrigger_STATUS_ARM `json:"requests,omitempty"`

	// SlowRequests: A rule based on request execution time.
	SlowRequests *SlowRequestsBasedTrigger_STATUS_ARM `json:"slowRequests,omitempty"`

	// SlowRequestsWithPath: A rule based on multiple Slow Requests Rule with path
	SlowRequestsWithPath []SlowRequestsBasedTrigger_STATUS_ARM `json:"slowRequestsWithPath,omitempty"`

	// StatusCodes: A rule based on status codes.
	StatusCodes []StatusCodesBasedTrigger_STATUS_ARM `json:"statusCodes,omitempty"`

	// StatusCodesRange: A rule based on status codes ranges.
	StatusCodesRange []StatusCodesRangeBasedTrigger_STATUS_ARM `json:"statusCodesRange,omitempty"`
}

type PushSettings_Properties_STATUS_ARM struct {
	// DynamicTagsJson: Gets or sets a JSON string containing a list of dynamic tags that will be evaluated from user claims in
	// the push registration endpoint.
	DynamicTagsJson *string `json:"dynamicTagsJson,omitempty"`

	// IsPushEnabled: Gets or sets a flag indicating whether the Push endpoint is enabled.
	IsPushEnabled *bool `json:"isPushEnabled,omitempty"`

	// TagWhitelistJson: Gets or sets a JSON string containing a list of tags that are in the allowed list for use by the push
	// registration endpoint.
	TagWhitelistJson *string `json:"tagWhitelistJson,omitempty"`

	// TagsRequiringAuth: Gets or sets a JSON string containing a list of tags that require user authentication to be used in
	// the push registration endpoint.
	// Tags can consist of alphanumeric characters and the following:
	// '_', '@', '#', '.', ':', '-'.
	// Validation should be performed at the PushRequestHandler.
	TagsRequiringAuth *string `json:"tagsRequiringAuth,omitempty"`
}

// Routing rules for ramp up testing. This rule allows to redirect static traffic % to a slot or to gradually change
// routing % based on performance.
type RampUpRule_STATUS_ARM struct {
	// ActionHostName: Hostname of a slot to which the traffic will be redirected if decided to. E.g.
	// myapp-stage.azurewebsites.net.
	ActionHostName *string `json:"actionHostName,omitempty"`

	// ChangeDecisionCallbackUrl: Custom decision algorithm can be provided in TiPCallback site extension which URL can be
	// specified. See TiPCallback site extension for the scaffold and contracts.
	// https://www.siteextensions.net/packages/TiPCallback/
	ChangeDecisionCallbackUrl *string `json:"changeDecisionCallbackUrl,omitempty"`

	// ChangeIntervalInMinutes: Specifies interval in minutes to reevaluate ReroutePercentage.
	ChangeIntervalInMinutes *int `json:"changeIntervalInMinutes,omitempty"`

	// ChangeStep: In auto ramp up scenario this is the step to add/remove from <code>ReroutePercentage</code> until it reaches
	// \n<code>MinReroutePercentage</code> or
	// <code>MaxReroutePercentage</code>. Site metrics are checked every N minutes specified in
	// <code>ChangeIntervalInMinutes</code>.\nCustom decision algorithm
	// can be provided in TiPCallback site extension which URL can be specified in <code>ChangeDecisionCallbackUrl</code>.
	ChangeStep *float64 `json:"changeStep,omitempty"`

	// MaxReroutePercentage: Specifies upper boundary below which ReroutePercentage will stay.
	MaxReroutePercentage *float64 `json:"maxReroutePercentage,omitempty"`

	// MinReroutePercentage: Specifies lower boundary above which ReroutePercentage will stay.
	MinReroutePercentage *float64 `json:"minReroutePercentage,omitempty"`

	// Name: Name of the routing rule. The recommended name would be to point to the slot which will receive the traffic in the
	// experiment.
	Name *string `json:"name,omitempty"`

	// ReroutePercentage: Percentage of the traffic which will be redirected to <code>ActionHostName</code>.
	ReroutePercentage *float64 `json:"reroutePercentage,omitempty"`
}

// Directory for virtual application.
type VirtualDirectory_STATUS_ARM struct {
	// PhysicalPath: Physical path.
	PhysicalPath *string `json:"physicalPath,omitempty"`

	// VirtualPath: Path to virtual application.
	VirtualPath *string `json:"virtualPath,omitempty"`
}

// Custom action to be executed
// when an auto heal rule is triggered.
type AutoHealCustomAction_STATUS_ARM struct {
	// Exe: Executable to be run.
	Exe *string `json:"exe,omitempty"`

	// Parameters: Parameters for the executable.
	Parameters *string `json:"parameters,omitempty"`
}

// Trigger based on total requests.
type RequestsBasedTrigger_STATUS_ARM struct {
	// Count: Request Count.
	Count *int `json:"count,omitempty"`

	// TimeInterval: Time interval.
	TimeInterval *string `json:"timeInterval,omitempty"`
}

// Trigger based on request execution time.
type SlowRequestsBasedTrigger_STATUS_ARM struct {
	// Count: Request Count.
	Count *int `json:"count,omitempty"`

	// Path: Request Path.
	Path *string `json:"path,omitempty"`

	// TimeInterval: Time interval.
	TimeInterval *string `json:"timeInterval,omitempty"`

	// TimeTaken: Time taken.
	TimeTaken *string `json:"timeTaken,omitempty"`
}

// Trigger based on status code.
type StatusCodesBasedTrigger_STATUS_ARM struct {
	// Count: Request Count.
	Count *int `json:"count,omitempty"`

	// Path: Request Path
	Path *string `json:"path,omitempty"`

	// Status: HTTP status code.
	Status *int `json:"status,omitempty"`

	// SubStatus: Request Sub Status.
	SubStatus *int `json:"subStatus,omitempty"`

	// TimeInterval: Time interval.
	TimeInterval *string `json:"timeInterval,omitempty"`

	// Win32Status: Win32 error code.
	Win32Status *int `json:"win32Status,omitempty"`
}

// Trigger based on range of status codes.
type StatusCodesRangeBasedTrigger_STATUS_ARM struct {
	// Count: Request Count.
	Count *int    `json:"count,omitempty"`
	Path  *string `json:"path,omitempty"`

	// StatusCodes: HTTP status code.
	StatusCodes *string `json:"statusCodes,omitempty"`

	// TimeInterval: Time interval.
	TimeInterval *string `json:"timeInterval,omitempty"`
}